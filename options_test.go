package log

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestNewOptions(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "default"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, NewOptions())
		})
	}
}

func TestOptions_AddFlags(t *testing.T) {
	type fields struct {
		OutputPaths       []string
		ErrorOutputPaths  []string
		Level             string
		Format            string
		DisableCaller     bool
		DisableStacktrace bool
		EnableColor       bool
		EncodeFullCaller  bool
		Development       bool
		Name              string
	}
	type args struct {
		fs *pflag.FlagSet
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "default",
			fields: fields{
				OutputPaths:       []string{"stdout"},
				ErrorOutputPaths:  []string{"stderr"},
				Level:             "info",
				Format:            "text",
				DisableCaller:     false,
				DisableStacktrace: false,
				EnableColor:       true,
				EncodeFullCaller:  false,
				Development:       false,
				Name:              "",
			},
			args: args{
				fs: pflag.NewFlagSet("test", pflag.ContinueOnError),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				OutputPaths:       tt.fields.OutputPaths,
				ErrorOutputPaths:  tt.fields.ErrorOutputPaths,
				Level:             tt.fields.Level,
				Format:            tt.fields.Format,
				DisableCaller:     tt.fields.DisableCaller,
				DisableStacktrace: tt.fields.DisableStacktrace,
				EnableColor:       tt.fields.EnableColor,
				EncodeFullCaller:  tt.fields.EncodeFullCaller,
				Development:       tt.fields.Development,
				Name:              tt.fields.Name,
			}
			o.AddFlags(tt.args.fs)
		})
	}
}

func TestOptions_Build(t *testing.T) {
	type fields struct {
		OutputPaths       []string
		ErrorOutputPaths  []string
		Level             string
		Format            string
		DisableCaller     bool
		DisableStacktrace bool
		EnableColor       bool
		EncodeFullCaller  bool
		Development       bool
		Name              string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "default",
			fields: fields{
				OutputPaths:       []string{"stdout"},
				ErrorOutputPaths:  []string{"stderr"},
				Level:             "info",
				Format:            "console",
				DisableCaller:     false,
				DisableStacktrace: false,
				EnableColor:       true,
				EncodeFullCaller:  false,
				Development:       false,
			},
			wantErr: assert.NoError,
		},
		{
			name: "error",
			fields: fields{
				OutputPaths:       []string{"stdout"},
				ErrorOutputPaths:  []string{"stderr"},
				Level:             "info",
				Format:            "text",
				DisableCaller:     false,
				DisableStacktrace: false,
				EnableColor:       true,
				EncodeFullCaller:  false,
				Development:       false,
			},
			wantErr: assert.Error,
		},
		{
			name: "level error",
			fields: fields{
				OutputPaths:       []string{"stdout"},
				ErrorOutputPaths:  []string{"stderr"},
				Level:             "illegal",
				Format:            "console",
				DisableCaller:     false,
				DisableStacktrace: false,
				EnableColor:       true,
				EncodeFullCaller:  true,
				Development:       false,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				OutputPaths:       tt.fields.OutputPaths,
				ErrorOutputPaths:  tt.fields.ErrorOutputPaths,
				Level:             tt.fields.Level,
				Format:            tt.fields.Format,
				DisableCaller:     tt.fields.DisableCaller,
				DisableStacktrace: tt.fields.DisableStacktrace,
				EnableColor:       tt.fields.EnableColor,
				EncodeFullCaller:  tt.fields.EncodeFullCaller,
				Development:       tt.fields.Development,
				Name:              tt.fields.Name,
			}
			tt.wantErr(t, o.Build())
		})
	}
}

func TestOptions_String(t *testing.T) {
	type fields struct {
		OutputPaths       []string
		ErrorOutputPaths  []string
		Level             string
		Format            string
		DisableCaller     bool
		DisableStacktrace bool
		EnableColor       bool
		EncodeFullCaller  bool
		Development       bool
		Name              string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "default",
			fields: fields{
				OutputPaths:       []string{"stdout"},
				ErrorOutputPaths:  []string{"stderr"},
				Level:             "info",
				Format:            "console",
				DisableCaller:     false,
				DisableStacktrace: false,
				EnableColor:       true,
				EncodeFullCaller:  false,
				Development:       false,
			},
			want: "{\"output-paths\":[\"stdout\"],\"error-output-paths\":[\"stderr\"],\"level\":\"info\",\"format\":\"console\",\"disable-caller\":false,\"disable-stacktrace\":false,\"enable-color\":true,\"enable-full-caller\":false,\"development\":false,\"name\":\"\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				OutputPaths:       tt.fields.OutputPaths,
				ErrorOutputPaths:  tt.fields.ErrorOutputPaths,
				Level:             tt.fields.Level,
				Format:            tt.fields.Format,
				DisableCaller:     tt.fields.DisableCaller,
				DisableStacktrace: tt.fields.DisableStacktrace,
				EnableColor:       tt.fields.EnableColor,
				EncodeFullCaller:  tt.fields.EncodeFullCaller,
				Development:       tt.fields.Development,
				Name:              tt.fields.Name,
			}
			assert.Equalf(t, tt.want, o.String(), "String()")
		})
	}
}

func TestOptions_Validate(t *testing.T) {
	type fields struct {
		OutputPaths       []string
		ErrorOutputPaths  []string
		Level             string
		Format            string
		DisableCaller     bool
		DisableStacktrace bool
		EnableColor       bool
		EncodeFullCaller  bool
		Development       bool
		Name              string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "invalid level",
			fields: fields{
				OutputPaths:       []string{"stdout"},
				ErrorOutputPaths:  []string{"stderr"},
				Level:             "invalid",
				Format:            "console",
				DisableCaller:     false,
				DisableStacktrace: false,
				EnableColor:       true,
				EncodeFullCaller:  false,
				Development:       false,
			},
			want: "unrecognized level: \"invalid\"",
		},
		{
			name: "invalid format",
			fields: fields{
				OutputPaths:       []string{"stdout"},
				ErrorOutputPaths:  []string{"stderr"},
				Level:             "info",
				Format:            "invalid",
				DisableCaller:     false,
				DisableStacktrace: false,
				EnableColor:       true,
				EncodeFullCaller:  false,
				Development:       false,
			},
			want: "not a valid log format: \"invalid\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				OutputPaths:       tt.fields.OutputPaths,
				ErrorOutputPaths:  tt.fields.ErrorOutputPaths,
				Level:             tt.fields.Level,
				Format:            tt.fields.Format,
				DisableCaller:     tt.fields.DisableCaller,
				DisableStacktrace: tt.fields.DisableStacktrace,
				EnableColor:       tt.fields.EnableColor,
				EncodeFullCaller:  tt.fields.EncodeFullCaller,
				Development:       tt.fields.Development,
				Name:              tt.fields.Name,
			}
			assert.Equalf(t, tt.want, o.Validate().Error(), "Validate()")
		})
	}
}

func TestOptions_Validate_nil(t *testing.T) {
	type fields struct {
		OutputPaths       []string
		ErrorOutputPaths  []string
		Level             string
		Format            string
		DisableCaller     bool
		DisableStacktrace bool
		EnableColor       bool
		EncodeFullCaller  bool
		Development       bool
		Name              string
	}
	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{
			name: "default",
			fields: fields{
				OutputPaths:       []string{"stdout"},
				ErrorOutputPaths:  []string{"stderr"},
				Level:             "info",
				Format:            "console",
				DisableCaller:     false,
				DisableStacktrace: false,
				EnableColor:       true,
				EncodeFullCaller:  false,
				Development:       false,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				OutputPaths:       tt.fields.OutputPaths,
				ErrorOutputPaths:  tt.fields.ErrorOutputPaths,
				Level:             tt.fields.Level,
				Format:            tt.fields.Format,
				DisableCaller:     tt.fields.DisableCaller,
				DisableStacktrace: tt.fields.DisableStacktrace,
				EnableColor:       tt.fields.EnableColor,
				EncodeFullCaller:  tt.fields.EncodeFullCaller,
				Development:       tt.fields.Development,
				Name:              tt.fields.Name,
			}
			assert.Equalf(t, tt.want, o.Validate(), "Validate()")
		})
	}
}
