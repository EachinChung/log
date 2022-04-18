package log

import (
	"testing"
	"time"

	"go.uber.org/zap/zapcore"
)

type enc struct{}

func (enc) AppendBool(bool) {}

func (enc) AppendByteString([]byte) {}

func (enc) AppendComplex128(complex128) {}

func (enc) AppendComplex64(complex64) {}

func (enc) AppendFloat64(float64) {}

func (enc) AppendFloat32(float32) {}

func (enc) AppendInt(int) {}

func (enc) AppendInt64(int64) {}

func (enc) AppendInt32(int32) {}

func (enc) AppendInt16(int16) {}

func (enc) AppendInt8(int8) {}

func (enc) AppendString(string) {}

func (enc) AppendUint(uint) {}

func (enc) AppendUint64(uint64) {}

func (enc) AppendUint32(uint32) {}

func (enc) AppendUint16(uint16) {}

func (enc) AppendUint8(uint8) {}

func (enc) AppendUintptr(uintptr) {}

var _ zapcore.PrimitiveArrayEncoder = enc{}

func Test_milliSecondsDurationEncoder(t *testing.T) {
	type args struct {
		d   time.Duration
		enc zapcore.PrimitiveArrayEncoder
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "milliSecondsDurationEncoder",
			args: args{
				d:   time.Millisecond,
				enc: enc{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			milliSecondsDurationEncoder(tt.args.d, tt.args.enc)
		})
	}
}

func Test_timeEncoder(t *testing.T) {
	type args struct {
		t   time.Time
		enc zapcore.PrimitiveArrayEncoder
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "timeEncoder",
			args: args{
				t:   time.Now(),
				enc: enc{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeEncoder(tt.args.t, tt.args.enc)
		})
	}
}
