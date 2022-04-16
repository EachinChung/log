package log

import (
	"testing"
	"time"

	"go.uber.org/zap/zapcore"
)

type enc struct{}

func (e enc) AppendBool(b bool) {}

func (e enc) AppendByteString(bytes []byte) {}

func (e enc) AppendComplex128(c complex128) {}

func (e enc) AppendComplex64(c complex64) {}

func (e enc) AppendFloat64(f float64) {}

func (e enc) AppendFloat32(f float32) {}

func (e enc) AppendInt(i int) {}

func (e enc) AppendInt64(i int64) {}

func (e enc) AppendInt32(i int32) {}

func (e enc) AppendInt16(i int16) {}

func (e enc) AppendInt8(i int8) {}

func (e enc) AppendString(s string) {}

func (e enc) AppendUint(u uint) {}

func (e enc) AppendUint64(u uint64) {}

func (e enc) AppendUint32(u uint32) {}

func (e enc) AppendUint16(u uint16) {}

func (e enc) AppendUint8(u uint8) {}

func (e enc) AppendUintptr(u uintptr) {}

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
