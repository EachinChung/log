package log

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestDebug(t *testing.T) {
	type args struct {
		msg    string
		fields []Field
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "debug",
			args: args{
				msg: "debug",
			},
		},
		{
			name: "debug with fields",
			args: args{
				msg: "debug with fields",
				fields: []Field{
					String("key", "value"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.msg, tt.args.fields...)
		})
	}
}

func TestDebugf(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "debugf",
			args: args{
				format: "debugf",
			},
		},
		{
			name: "debugf with fields",
			args: args{
				format: "debugf with fields %s",
				v:      []interface{}{"value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debugf(tt.args.format, tt.args.v...)
		})
	}
}

func TestDebugw(t *testing.T) {
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "debugw",
			args: args{
				msg: "debugw",
			},
		},
		{
			name: "debugw with fields",
			args: args{
				msg: "debugw with fields",
				keysAndValues: []interface{}{
					"key", "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debugw(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		msg    string
		fields []Field
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "error",
			args: args{
				msg: "error",
			},
		},
		{
			name: "error with fields",
			args: args{
				msg: "error with fields",
				fields: []Field{
					String("key", "value"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.msg, tt.args.fields...)
		})
	}
}

func TestErrorf(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "errorf",
			args: args{
				format: "errorf",
			},
		},
		{
			name: "errorf with fields",
			args: args{
				format: "errorf with fields %s",
				v:      []interface{}{"value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Errorf(tt.args.format, tt.args.v...)
		})
	}
}

func TestErrorw(t *testing.T) {
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "errorw",
			args: args{
				msg: "errorw",
			},
		},
		{
			name: "errorw with fields",
			args: args{
				msg: "errorw with fields",
				keysAndValues: []interface{}{
					"key", "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Errorw(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func TestFlush(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "flush",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Flush()
		})
	}
}

func TestInfo(t *testing.T) {
	type args struct {
		msg    string
		fields []Field
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "info",
			args: args{
				msg: "info",
			},
		},
		{
			name: "info with fields",
			args: args{
				msg: "info with fields",
				fields: []Field{
					String("key", "value"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.msg, tt.args.fields...)
		})
	}
}

func TestInfof(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "infof",
			args: args{
				format: "infof",
			},
		},
		{
			name: "infof with fields",
			args: args{
				format: "infof with fields %s",
				v:      []interface{}{"value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Infof(tt.args.format, tt.args.v...)
		})
	}
}

func TestInfow(t *testing.T) {
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "infow",
			args: args{
				msg: "infow",
			},
		},
		{
			name: "infow with fields",
			args: args{
				msg: "infow with fields",
				keysAndValues: []interface{}{
					"key", "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Infow(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func TestInit(t *testing.T) {
	type args struct {
		opts *Options
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "init",
			args: args{
				opts: NewOptions(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init(tt.args.opts)
		})
	}
}

func TestL(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want *zapLogger
	}{
		{
			name: "l",
			args: args{
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, L(tt.args.ctx))
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		opts *Options
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "opts nil",
			args: args{
				opts: nil,
			},
		},
		{
			name: "new",
			args: args{
				opts: NewOptions(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, New(tt.args.opts), "New(%v)")
		})
	}
}

func TestNewLogger(t *testing.T) {
	type args struct {
		l *zap.Logger
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "new logger",
			args: args{
				l: zap.NewNop(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, NewLogger(tt.args.l), "NewLogger(%v)", tt.args.l)
		})
	}
}

func TestPanic(t *testing.T) {
	type args struct {
		msg    string
		fields []Field
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "panic",
			args: args{
				msg: "panic",
			},
		},
		{
			name: "panic with fields",
			args: args{
				msg: "panic with fields",
				fields: []Field{
					String("key", "value"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panics(t, func() {
				Panic(tt.args.msg, tt.args.fields...)
			})
		})
	}
}

func TestPanicf(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "panicf",
			args: args{
				format: "panicf",
			},
		},
		{
			name: "panicf with fields",
			args: args{
				format: "panicf with fields %s",
				v:      []interface{}{"value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panics(t, func() {
				Panicf(tt.args.format, tt.args.v...)
			})
		})
	}
}

func TestPanicw(t *testing.T) {
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "panicw",
			args: args{
				msg: "panicw",
			},
		},
		{
			name: "panicw with fields",
			args: args{
				msg: "panicw with fields",
				keysAndValues: []interface{}{
					"key", "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panics(t, func() {
				Panicw(tt.args.msg, tt.args.keysAndValues...)
			})
		})
	}
}

func TestStdErrLogger(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "std err logger"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, StdErrLogger())
		})
	}
}

func TestStdInfoLogger(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "std info logger"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, StdInfoLogger())
		})
	}
}

func TestSugaredLogger(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "sugared logger"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, SugaredLogger())
		})
	}
}

func TestV(t *testing.T) {
	type args struct {
		level int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "v",
			args: args{
				level: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, V(tt.args.level))
		})
	}
}

func TestWarn(t *testing.T) {
	type args struct {
		msg    string
		fields []Field
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "warn",
			args: args{
				msg: "warn",
			},
		},
		{
			name: "warn with fields",
			args: args{
				msg: "warn with fields",
				fields: []Field{
					String("key", "value"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warn(tt.args.msg, tt.args.fields...)
		})
	}
}

func TestWarnf(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "warnf",
			args: args{
				format: "warnf",
			},
		},
		{
			name: "warnf with fields",
			args: args{
				format: "warnf with fields %s",
				v:      []interface{}{"value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warnf(tt.args.format, tt.args.v...)
		})
	}
}

func TestWarnw(t *testing.T) {
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "warnw",
			args: args{
				msg: "warnw",
			},
		},
		{
			name: "warnw with fields",
			args: args{
				msg: "warnw with fields",
				keysAndValues: []interface{}{
					"key", "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warnw(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func TestWithName(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "with name",
			args: args{
				s: "with name",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, WithName(tt.args.s))
		})
	}
}

func TestWithValues(t *testing.T) {
	type args struct {
		keysAndValues []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "with values",
			args: args{
				keysAndValues: []interface{}{
					"key", "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, WithValues(tt.args.keysAndValues...))
		})
	}
}

func Test_infoLogger_Enabled(t *testing.T) {
	type fields struct {
		level zapcore.Level
		log   *zap.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "enabled",
			fields: fields{
				level: zap.InfoLevel,
				log:   zap.NewNop(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &infoLogger{
				level: tt.fields.level,
				log:   tt.fields.log,
			}
			assert.Equalf(t, tt.want, l.Enabled(), "Enabled()")
		})
	}
}

func Test_infoLogger_Info(t *testing.T) {
	type fields struct {
		level zapcore.Level
		log   *zap.Logger
	}
	type args struct {
		msg    string
		fields []Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "info",
			fields: fields{
				level: zap.InfoLevel,
				log:   zap.NewNop(),
			},
			args: args{
				msg: "info",
			},
		},
		{
			name: "info with fields",
			fields: fields{
				level: zap.InfoLevel,
				log:   zap.NewNop(),
			},
			args: args{
				msg: "info with fields",
				fields: []Field{
					String("key", "value"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &infoLogger{
				level: tt.fields.level,
				log:   tt.fields.log,
			}
			l.Info(tt.args.msg, tt.args.fields...)
		})
	}
}

func Test_infoLogger_Infof(t *testing.T) {
	type fields struct {
		level zapcore.Level
		log   *zap.Logger
	}
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "infof",
			fields: fields{
				level: zap.InfoLevel,
				log:   zap.NewNop(),
			},
			args: args{
				format: "infof",
			},
		},
		{
			name: "infof with fields",
			fields: fields{
				level: zap.InfoLevel,
				log:   zap.NewNop(),
			},
			args: args{
				format: "infof with fields %s",
				args:   []interface{}{"value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &infoLogger{
				level: tt.fields.level,
				log:   tt.fields.log,
			}
			l.Infof(tt.args.format, tt.args.args...)
		})
	}
}

func Test_infoLogger_Infow(t *testing.T) {
	type fields struct {
		level zapcore.Level
		log   *zap.Logger
	}
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "infow",
			fields: fields{
				level: zap.InfoLevel,
				log:   zap.NewNop(),
			},
			args: args{
				msg: "infow",
			},
		},
		{
			name: "infow with fields",
			fields: fields{
				level: zap.InfoLevel,
				log:   zap.NewNop(),
			},
			args: args{
				msg: "infow with fields",
				keysAndValues: []interface{}{
					"key", "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &infoLogger{
				level: tt.fields.level,
				log:   tt.fields.log,
			}
			l.Infow(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func Test_noopInfoLogger_Enabled(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "enabled",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &noopInfoLogger{}
			assert.Equalf(t, tt.want, l.Enabled(), "Enabled()")
		})
	}
}

func Test_noopInfoLogger_Info(t *testing.T) {
	type args struct {
		in0 string
		in1 []Field
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "info",
			args: args{
				in0: "info",
			},
		},
		{
			name: "info with fields",
			args: args{
				in0: "info with fields",
				in1: []Field{
					String("key", "value"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &noopInfoLogger{}
			l.Info(tt.args.in0, tt.args.in1...)
		})
	}
}

func Test_noopInfoLogger_Infof(t *testing.T) {
	type args struct {
		in0 string
		in1 []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "infof",
			args: args{
				in0: "infof",
			},
		},
		{
			name: "infof with fields",
			args: args{
				in0: "infof with fields %s",
				in1: []interface{}{"value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &noopInfoLogger{}
			l.Infof(tt.args.in0, tt.args.in1...)
		})
	}
}

func Test_noopInfoLogger_Infow(t *testing.T) {
	type args struct {
		in0 string
		in1 []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "infow",
			args: args{
				in0: "infow",
			},
		},
		{
			name: "infow with fields",
			args: args{
				in0: "infow with fields",
				in1: []interface{}{
					"key", "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &noopInfoLogger{}
			l.Infow(tt.args.in0, tt.args.in1...)
		})
	}
}

func Test_zapLogger_Debug(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		msg    string
		fields []Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "debug",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "debug",
			},
		},
		{
			name: "debug with fields",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "debug with fields",
				fields: []Field{
					String("key", "value"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			l.Debug(tt.args.msg, tt.args.fields...)
		})
	}
}

func Test_zapLogger_Debugf(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "debugf",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				format: "debugf",
			},
		},
		{
			name: "debugf with fields",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				format: "debugf with fields %s",
				v:      []interface{}{"value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			l.Debugf(tt.args.format, tt.args.v...)
		})
	}
}

func Test_zapLogger_Debugw(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "debugw",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "debugw",
			},
		},
		{
			name: "debugw with fields",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "debugw with fields",
				keysAndValues: []interface{}{
					"key", "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			l.Debugw(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func Test_zapLogger_Error(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		msg    string
		fields []Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "error",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "error",
			},
		},
		{
			name: "error with fields",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "error with fields",
				fields: []Field{
					String("key", "value"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			l.Error(tt.args.msg, tt.args.fields...)
		})
	}
}

func Test_zapLogger_Errorf(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "errorf",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				format: "errorf",
			},
		},
		{
			name: "errorf with fields",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				format: "errorf with fields %s",
				v:      []interface{}{"value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			l.Errorf(tt.args.format, tt.args.v...)
		})
	}
}

func Test_zapLogger_Errorw(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "errorw",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "errorw",
			},
		},
		{
			name: "errorw with fields",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "errorw with fields",
				keysAndValues: []interface{}{
					"key", "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			l.Errorw(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func Test_zapLogger_Flush(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "flush",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			l.Flush()
		})
	}
}

func Test_zapLogger_Info(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		msg    string
		fields []Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "info",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "info",
			},
		},
		{
			name: "info with fields",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "info with fields",
				fields: []Field{
					String("key", "value"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			l.Info(tt.args.msg, tt.args.fields...)
		})
	}
}

func Test_zapLogger_Infof(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "infof",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				format: "infof",
			},
		},
		{
			name: "infof with fields",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				format: "infof with fields %s",
				v:      []interface{}{"value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			l.Infof(tt.args.format, tt.args.v...)
		})
	}
}

func Test_zapLogger_Infow(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "infow",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "infow",
			},
		},
		{
			name: "infow with fields",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "infow with fields",
				keysAndValues: []interface{}{
					"key", "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			l.Infow(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func Test_zapLogger_L(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "l",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			assert.NotNil(t, l.L(tt.args.ctx), "L(%v)", tt.args.ctx)
		})
	}
}

func Test_zapLogger_Panic(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		msg    string
		fields []Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "panic",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "panic",
			},
		},
		{
			name: "panic with fields",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "panic with fields",
				fields: []Field{
					String("key", "value"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panics(t, func() {
				l := &zapLogger{
					zapLogger:  tt.fields.zapLogger,
					infoLogger: tt.fields.infoLogger,
				}
				l.Panic(tt.args.msg, tt.args.fields...)
			})
		})
	}
}

func Test_zapLogger_Panicf(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "panicf",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				format: "panicf",
			},
		},
		{
			name: "panicf with fields",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				format: "panicf with fields",
				v:      []interface{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panics(t, func() {
				l := &zapLogger{
					zapLogger:  tt.fields.zapLogger,
					infoLogger: tt.fields.infoLogger,
				}
				l.Panicf(tt.args.format, tt.args.v...)
			})
		})
	}
}

func Test_zapLogger_Panicw(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "panicw",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "panicw",
			},
		},
		{
			name: "panicw with fields",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "panicw with fields",
				keysAndValues: []interface{}{
					"key", "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panics(t, func() {
				l := &zapLogger{
					zapLogger:  tt.fields.zapLogger,
					infoLogger: tt.fields.infoLogger,
				}
				l.Panicw(tt.args.msg, tt.args.keysAndValues...)
			})
		})
	}
}

func Test_zapLogger_V(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		level int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "v",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				level: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			assert.NotNil(t, l.V(tt.args.level), "V(%v)", tt.args.level)
		})
	}
}

func Test_zapLogger_Warn(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		msg    string
		fields []Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "warn",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "warn",
			},
		},
		{
			name: "warn with fields",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "warn with fields",
				fields: []Field{
					String("key", "value"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			l.Warn(tt.args.msg, tt.args.fields...)
		})
	}
}

func Test_zapLogger_Warnf(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "warnf",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				format: "warnf",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			l.Warnf(tt.args.format, tt.args.v...)
		})
	}
}

func Test_zapLogger_Warnw(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		msg           string
		keysAndValues []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "warnw",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "warnw",
			},
		},
		{
			name: "warnw with fields",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				msg: "warnw with fields",
				keysAndValues: []interface{}{
					"key", "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			l.Warnw(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func Test_zapLogger_WithName(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "withName",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				name: "withName",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			assert.NotNil(t, l.WithName(tt.args.name), "WithName(%v)", tt.args.name)
		})
	}
}

func Test_zapLogger_WithValues(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		keysAndValues []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Logger
	}{
		{
			name: "withValues",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				keysAndValues: []interface{}{
					"key", "value",
				},
			},
			want: &zapLogger{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			assert.NotNil(t, tt.want, l.WithValues(tt.args.keysAndValues...))
		})
	}
}

func Test_zapLogger_Write(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "write",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
			args: args{
				p: []byte("write"),
			},
			wantN:   len([]byte("write")),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			gotN, err := l.Write(tt.args.p)
			if !tt.wantErr(t, err, fmt.Sprintf("Write(%v)", tt.args.p)) {
				return
			}
			assert.Equalf(t, tt.wantN, gotN, "Write(%v)", tt.args.p)
		})
	}
}

func Test_zapLogger_clone(t *testing.T) {
	type fields struct {
		zapLogger  *zap.Logger
		infoLogger infoLogger
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "clone",
			fields: fields{
				zapLogger:  zap.NewNop(),
				infoLogger: infoLogger{level: zap.InfoLevel, log: zap.NewNop()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &zapLogger{
				zapLogger:  tt.fields.zapLogger,
				infoLogger: tt.fields.infoLogger,
			}
			assert.NotNil(t, l.clone(), "clone()")
		})
	}
}
