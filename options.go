package log

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/eachinchung/errors"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	flagLevel             = "log.level"
	flagDisableCaller     = "log.disable-caller"
	flagDisableStacktrace = "log.disable-stacktrace"
	flagFormat            = "log.format"
	flagEnableColor       = "log.enable-color"
	flagEncodeFullCaller  = "log.enable-full-caller"
	flagOutputPaths       = "log.output-paths"
	flagErrorOutputPaths  = "log.error-output-paths"
	flagDevelopment       = "log.development"
	flagName              = "log.name"

	consoleFormat = "console"
	jsonFormat    = "json"
)

// Options 日志相关的配置项。
type Options struct {
	OutputPaths       []string `json:"output-paths"       mapstructure:"output-paths"`
	ErrorOutputPaths  []string `json:"error-output-paths" mapstructure:"error-output-paths"`
	Level             string   `json:"level"              mapstructure:"level"`
	Format            string   `json:"format"             mapstructure:"format"`
	DisableCaller     bool     `json:"disable-caller"     mapstructure:"disable-caller"`
	DisableStacktrace bool     `json:"disable-stacktrace" mapstructure:"disable-stacktrace"`
	EnableColor       bool     `json:"enable-color"       mapstructure:"enable-color"`
	EncodeFullCaller  bool     `json:"enable-full-caller" mapstructure:"enable-full-caller"`
	Development       bool     `json:"development"        mapstructure:"development"`
	Name              string   `json:"name"               mapstructure:"name"`
}

// NewOptions 创建一个带有默认参数的 Options 对象。
func NewOptions() *Options {
	return &Options{
		Level:             zapcore.InfoLevel.String(),
		DisableCaller:     false,
		DisableStacktrace: false,
		Format:            consoleFormat,
		EnableColor:       true,
		EncodeFullCaller:  false,
		Development:       false,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	}
}

// Validate 验证选项字段。
func (o *Options) Validate() error {
	var errs []error

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(o.Level)); err != nil {
		errs = append(errs, err)
	}

	format := strings.ToLower(o.Format)
	if format != consoleFormat && format != jsonFormat {
		errs = append(errs, fmt.Errorf("not a valid log format: %q", o.Format))
	}

	return errors.NewAggregate(errs...)
}

// AddFlags 将 Options 的各个字段追加到传入的 pflag.FlagSet 变量中。
func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Level, flagLevel, o.Level,
		"日志级别，优先级从低到高依次为: Debug, Info, Warn, Error, Dpanic, Panic, Fatal。")
	fs.BoolVar(&o.DisableCaller, flagDisableCaller, o.DisableCaller,
		"是否开启 caller，如果开启会在日志中显示调用日志所在的文件、函数和行号。")
	fs.BoolVar(&o.DisableStacktrace, flagDisableStacktrace,
		o.DisableStacktrace, "是否在 Panic 及以上级别禁止打印堆栈信息。")
	fs.StringVar(&o.Format, flagFormat, o.Format,
		"支持的日志输出格式，目前支持 Console 和 JSON 两种。Console 其实就是 Text 格式。")
	fs.BoolVar(&o.EnableColor, flagEnableColor, o.EnableColor, "是否开启颜色输出，true，是；false，否。")
	fs.BoolVar(&o.EncodeFullCaller, flagEncodeFullCaller, o.EncodeFullCaller,
		"是否开启 full caller /full/path/to/package/file:line，true:是，false:否")
	fs.StringSliceVar(&o.OutputPaths, flagOutputPaths, o.OutputPaths,
		"支持输出到多个输出，用逗号分开。支持输出到标准输出(stdout)和文件。")
	fs.StringSliceVar(&o.ErrorOutputPaths, flagErrorOutputPaths, o.ErrorOutputPaths,
		"zap 内部 (非业务) 错误日志输出路径，多个输出，用逗号分开")
	fs.BoolVar(&o.Development, flagDevelopment, o.Development,
		"是否是开发模式。如果是开发模式，会对 DPanicLevel 进行堆栈跟踪。")
	fs.StringVar(&o.Name, flagName, o.Name, "Logger 的名字。")
}

// String 将 Options 的值以 JSON 格式字符串返回。
func (o *Options) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}

// Build 根据 Options 构建一个全局的 Logger。
func (o *Options) Build() error {
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(o.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}
	encodeLevel := zapcore.CapitalLevelEncoder
	if o.Format == consoleFormat && o.EnableColor {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var encodeCaller zapcore.CallerEncoder
	if o.EncodeFullCaller {
		encodeCaller = zapcore.FullCallerEncoder
	} else {
		encodeCaller = zapcore.ShortCallerEncoder
	}

	zc := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       o.Development,
		DisableCaller:     o.DisableCaller,
		DisableStacktrace: o.DisableStacktrace,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: o.Format,
		EncoderConfig: zapcore.EncoderConfig{
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
			EncodeName:     zapcore.FullNameEncoder,
		},
		OutputPaths:      o.OutputPaths,
		ErrorOutputPaths: o.ErrorOutputPaths,
	}
	logger, err := zc.Build(zap.AddStacktrace(zapcore.PanicLevel))
	if err != nil {
		return err
	}
	zap.RedirectStdLog(logger.Named(o.Name))
	zap.ReplaceGlobals(logger)

	return nil
}
