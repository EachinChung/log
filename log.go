package log

import (
	"context"
	"fmt"
	"log"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InfoLogger 表示能够以特定的详细程度记录非错误消息。
type InfoLogger interface {
	// Info 使用给定的键/值对作为上下文记录非错误消息。
	//
	// msg 参数应该用于向日志行添加一些常量描述。
	// 然后可以使用键/值对添加额外的变量信息。
	// 键/值对应该交替使用字符串键和任意值。
	Info(msg string, fields ...Field)
	Infof(format string, v ...interface{})
	Infow(msg string, keysAndValues ...interface{})

	// Enabled 测试是否启用了这个 InfoLogger。
	// 例如，命令行标志可用于设置日志记录详细程度并禁用某些信息日志。
	Enabled() bool
}

// Logger 记录日志消息。
type Logger interface {
	// InfoLogger 所有 Logger 都实现了 InfoLogger。
	// 直接在 Logger 值上调用 InfoLogger 方法等效于在 V(0) InfoLogger 上调用它们。
	// 例如，logger.Info() 产生与 logger.V(0).Info 相同的结果。
	InfoLogger

	Debug(msg string, fields ...Field)
	Debugf(format string, v ...interface{})
	Debugw(msg string, keysAndValues ...interface{})

	Warn(msg string, fields ...Field)
	Warnf(format string, v ...interface{})
	Warnw(msg string, keysAndValues ...interface{})

	Error(msg string, fields ...Field)
	Errorf(format string, v ...interface{})
	Errorw(msg string, keysAndValues ...interface{})

	Panic(msg string, fields ...Field)
	Panicf(format string, v ...interface{})
	Panicw(msg string, keysAndValues ...interface{})

	Fatal(msg string, fields ...Field)
	Fatalf(format string, v ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})

	// V 返回特定详细级别的 InfoLogger 值。
	// 一个更高的详细级别意味着日志消息不太重要。
	// 传递小于零的日志级别是违法的。
	V(level int) InfoLogger

	Write(p []byte) (n int, err error)

	// WithValues 向 logger 添加一些上下文的键值对。
	// 有关键/值对如何工作的文档，请参阅 Info。
	WithValues(keysAndValues ...interface{}) Logger

	// WithName 向记录器的名称添加一个新元素。
	// 使用 WithName 的连续调用继续向记录器的名称附加后缀。
	// 强烈建议名称段仅包含字母、数字和连字符。
	WithName(name string) Logger

	// WithContext 返回设置日志值的上下文副本。
	WithContext(ctx context.Context) context.Context

	// Flush 调用底层 Core 的 Sync 方法，刷新所有缓冲的日志条目。
	// 应用程序应注意在退出前调用 Sync。
	Flush()
}

var _ Logger = &zapLogger{}

// noopInfoLogger 是一个 log.InfoLogger，它总是被禁用，什么也不做。
type noopInfoLogger struct{}

func (l *noopInfoLogger) Enabled() bool             { return false }
func (l *noopInfoLogger) Info(_ string, _ ...Field) {}

func (l *noopInfoLogger) Infof(_ string, _ ...interface{}) {}

func (l *noopInfoLogger) Infow(_ string, _ ...interface{}) {}

var disabledInfoLogger = &noopInfoLogger{}

// infoLogger 是一个 log.InfoLogger，它使用 Zap 在特定级别进行记录。
// 该级别已经转换为Zap级别，也就是说`logLevel = -1*zapLevel`。
//
// 注意：现在，我们总是使用相当于 sugared logging 的 logger。
type infoLogger struct {
	level zapcore.Level
	log   *zap.Logger
}

func (l *infoLogger) Enabled() bool { return true }
func (l *infoLogger) Info(msg string, fields ...Field) {
	if checkedEntry := l.log.Check(l.level, msg); checkedEntry != nil {
		checkedEntry.Write(fields...)
	}
}

func (l *infoLogger) Infof(format string, args ...interface{}) {
	if checkedEntry := l.log.Check(l.level, fmt.Sprintf(format, args...)); checkedEntry != nil {
		checkedEntry.Write()
	}
}

func (l *infoLogger) Infow(msg string, keysAndValues ...interface{}) {
	if checkedEntry := l.log.Check(l.level, msg); checkedEntry != nil {
		checkedEntry.Write(handleFields(l.log, keysAndValues)...)
	}
}

// zapLogger 是一个使用 Zap 进行日志记录的 log.Logger。
//
// 注意：这看起来与 zap.SugaredLogger 非常相似，但我们希望拥有多个详细级别。
type zapLogger struct {
	zapLogger *zap.Logger
	infoLogger
}

// handleFields 将一堆任意键值对转换为 Zap 字段。 它需要额外的预先转换的 Zap 字段，用于自动附加的字段，如 `error`。
func handleFields(l *zap.Logger, args []interface{}, additional ...zap.Field) []zap.Field {
	// zap.SugaredLogger.sweetenFields 的一个稍微修改的版本
	if len(args) == 0 {
		// 如果我们没有建议的字段，则快速返回。
		return additional
	}

	// 与 Zap 不同，我们可以非常确定用户不会传递结构化字段（因为 log 没有这个概念），所以猜测我们需要更少的空间。
	fields := make([]zap.Field, 0, len(args)/2+len(additional))
	for i := 0; i < len(args); {
		// 检查强类型的 Zap 字段，这是非法的
		if _, ok := args[i].(zap.Field); ok {
			l.DPanic("传递了强类型的 Zap 字段", zap.Any("Zap 字段", args[i]))

			break
		}

		// 确保这不是一个不匹配的键
		if i == len(args)-1 {
			l.DPanic("传递给日志记录的键值对参数个数为奇数", zap.Any("忽略的键", args[i]))

			break
		}

		// 处理一个键值对，
		// 确保键是字符串
		key, val := args[i], args[i+1]
		keyStr, isString := key.(string)
		if !isString {
			// 如果键不是字符串，DPanic 并停止记录
			l.DPanic(
				"传递给 log 的参数键为非字符串，忽略所有后面的参数",
				zap.Any("无效的 key", key),
			)

			break
		}

		fields = append(fields, zap.Any(keyStr, val))
		i += 2
	}

	return append(fields, additional...)
}

var (
	std = New(NewOptions())
	mu  sync.Mutex
)

// Init 使用指定的选项初始 logger。
func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()
	std = New(opts)
}

// New 通过 opts 创建 logger。
func New(opts *Options) *zapLogger {
	if opts == nil {
		opts = NewOptions()
	}

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}
	encodeLevel := zapcore.CapitalLevelEncoder
	// 当输出到本地路径时，禁止使用颜色
	if opts.Format == consoleFormat && opts.EnableColor {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var encodeCaller zapcore.CallerEncoder
	if opts.EncodeFullCaller {
		encodeCaller = zapcore.FullCallerEncoder
	} else {
		encodeCaller = zapcore.ShortCallerEncoder
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encodeLevel,
		EncodeTime:     timeEncoder,
		EncodeDuration: milliSecondsDurationEncoder,
		EncodeCaller:   encodeCaller,
	}

	loggerConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       opts.Development,
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         opts.Format,
		EncoderConfig:    encoderConfig,
		OutputPaths:      opts.OutputPaths,
		ErrorOutputPaths: opts.ErrorOutputPaths,
	}

	l, err := loggerConfig.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logger := &zapLogger{
		zapLogger: l.Named(opts.Name),
		infoLogger: infoLogger{
			log:   l,
			level: zap.InfoLevel,
		},
	}
	zap.RedirectStdLog(l)

	return logger
}

// SugaredLogger 返回全局 sugared logger.
func SugaredLogger() *zap.SugaredLogger {
	return std.zapLogger.Sugar()
}

// StdErrLogger 返回标准库的 logger，该 logger 在 error level 写入提供的 zap logger。
func StdErrLogger() *log.Logger {
	if std == nil {
		return nil
	}
	if l, err := zap.NewStdLogAt(std.zapLogger, zapcore.ErrorLevel); err == nil {
		return l
	}

	return nil
}

// StdInfoLogger 返回标准库的 logger，该 logger 在 info level 写入提供的 zap logger。
func StdInfoLogger() *log.Logger {
	if std == nil {
		return nil
	}
	if l, err := zap.NewStdLogAt(std.zapLogger, zapcore.InfoLevel); err == nil {
		return l
	}

	return nil
}

// V 返回一个分级的 InfoLogger。
func V(level int) InfoLogger { return std.V(level) }

func (l *zapLogger) V(level int) InfoLogger {
	lvl := zapcore.Level(5 - 1*level)
	if l.zapLogger.Core().Enabled(lvl) {
		return &infoLogger{
			level: lvl,
			log:   l.zapLogger,
		}
	}

	return disabledInfoLogger
}

func (l *zapLogger) Write(p []byte) (n int, err error) {
	l.zapLogger.Info(string(p))

	return len(p), nil
}

// WithValues 创建一个 child logger 并向其添加 Zap 字段。
func WithValues(keysAndValues ...interface{}) Logger { return std.WithValues(keysAndValues...) }

func (l *zapLogger) WithValues(keysAndValues ...interface{}) Logger {
	newLogger := l.zapLogger.With(handleFields(l.zapLogger, keysAndValues)...)

	return NewLogger(newLogger)
}

// WithName 为 logger 的名称添加一个新的路径段。默认情况下，记录器是未命名的。
func WithName(s string) Logger { return std.WithName(s) }

func (l *zapLogger) WithName(name string) Logger {
	newLogger := l.zapLogger.Named(name)

	return NewLogger(newLogger)
}

// Flush 调用底层 Core 的 Sync 方法，刷新所有缓冲的日志条目。
// 应用程序应注意在退出前调用 Sync。
func Flush() { std.Flush() }

func (l *zapLogger) Flush() {
	_ = l.zapLogger.Sync()
}

// NewLogger 使用给定的 Zap Logger 创建一个新的 log.Logger 来记录日志。
func NewLogger(l *zap.Logger) Logger {
	return &zapLogger{
		zapLogger: l,
		infoLogger: infoLogger{
			log:   l,
			level: zap.InfoLevel,
		},
	}
}

// Debug method output debug level log.
func Debug(msg string, fields ...Field) {
	std.zapLogger.Debug(msg, fields...)
}

func (l *zapLogger) Debug(msg string, fields ...Field) {
	l.zapLogger.Debug(msg, fields...)
}

// Debugf method output debug level log.
func Debugf(format string, v ...interface{}) {
	std.zapLogger.Sugar().Debugf(format, v...)
}

func (l *zapLogger) Debugf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Debugf(format, v...)
}

// Debugw method output debug level log.
func Debugw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Debugw(msg, keysAndValues...)
}

func (l *zapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Debugw(msg, keysAndValues...)
}

// Info method output info level log.
func Info(msg string, fields ...Field) {
	std.zapLogger.Info(msg, fields...)
}

func (l *zapLogger) Info(msg string, fields ...Field) {
	l.zapLogger.Info(msg, fields...)
}

// Infof method output info level log.
func Infof(format string, v ...interface{}) {
	std.zapLogger.Sugar().Infof(format, v...)
}

func (l *zapLogger) Infof(format string, v ...interface{}) {
	l.zapLogger.Sugar().Infof(format, v...)
}

// Infow method output info level log.
func Infow(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Infow(msg, keysAndValues...)
}

func (l *zapLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Infow(msg, keysAndValues...)
}

// Warn method output warning level log.
func Warn(msg string, fields ...Field) {
	std.zapLogger.Warn(msg, fields...)
}

func (l *zapLogger) Warn(msg string, fields ...Field) {
	l.zapLogger.Warn(msg, fields...)
}

// Warnf method output warning level log.
func Warnf(format string, v ...interface{}) {
	std.zapLogger.Sugar().Warnf(format, v...)
}

func (l *zapLogger) Warnf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Warnf(format, v...)
}

// Warnw method output warning level log.
func Warnw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Warnw(msg, keysAndValues...)
}

func (l *zapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Warnw(msg, keysAndValues...)
}

// Error method output error level log.
func Error(msg string, fields ...Field) {
	std.zapLogger.Error(msg, fields...)
}

func (l *zapLogger) Error(msg string, fields ...Field) {
	l.zapLogger.Error(msg, fields...)
}

// Errorf method output error level log.
func Errorf(format string, v ...interface{}) {
	std.zapLogger.Sugar().Errorf(format, v...)
}

func (l *zapLogger) Errorf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Errorf(format, v...)
}

// Errorw method output error level log.
func Errorw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Errorw(msg, keysAndValues...)
}

func (l *zapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Errorw(msg, keysAndValues...)
}

// Panic method output panic level log and shutdown application.
func Panic(msg string, fields ...Field) {
	std.zapLogger.Panic(msg, fields...)
}

func (l *zapLogger) Panic(msg string, fields ...Field) {
	l.zapLogger.Panic(msg, fields...)
}

// Panicf method output panic level log and shutdown application.
func Panicf(format string, v ...interface{}) {
	std.zapLogger.Sugar().Panicf(format, v...)
}

func (l *zapLogger) Panicf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Panicf(format, v...)
}

// Panicw method output panic level log.
func Panicw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Panicw(msg, keysAndValues...)
}

func (l *zapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Panicw(msg, keysAndValues...)
}

// Fatal method output fatal level log.
func Fatal(msg string, fields ...Field) {
	std.zapLogger.Fatal(msg, fields...)
}

func (l *zapLogger) Fatal(msg string, fields ...Field) {
	l.zapLogger.Fatal(msg, fields...)
}

// Fatalf method output fatal level log.
func Fatalf(format string, v ...interface{}) {
	std.zapLogger.Sugar().Fatalf(format, v...)
}

func (l *zapLogger) Fatalf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Fatalf(format, v...)
}

// Fatalw method output Fatalw level log.
func Fatalw(msg string, keysAndValues ...interface{}) {
	std.zapLogger.Sugar().Fatalw(msg, keysAndValues...)
}

func (l *zapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Fatalw(msg, keysAndValues...)
}

// L method output with specified context value.
func L(ctx context.Context) *zapLogger {
	return std.L(ctx)
}

func (l *zapLogger) L(ctx context.Context) *zapLogger {
	lg := l.clone()

	if requestID := ctx.Value(KeyRequestID); requestID != nil {
		lg.zapLogger = lg.zapLogger.With(zap.Any("request-id", requestID))
	}

	if eID := ctx.Value(KeyEID); eID != nil {
		lg.zapLogger = lg.zapLogger.With(zap.Any("e-id", eID))
	}

	return lg
}

func (l *zapLogger) clone() *zapLogger {
	copyLogger := *l

	return &copyLogger
}
