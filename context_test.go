package log

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromContext(t *testing.T) {
	l := New(nil)
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want Logger
	}{
		{
			name: "from context",
			args: args{
				ctx: l.WithContext(context.Background()),
			},
			want: l,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FromContext(tt.args.ctx), "FromContext(%v)", tt.args.ctx)
		})
	}
}

func TestFromContext_nil(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "from nil context",
			args: args{
				ctx: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, FromContext(tt.args.ctx))
		})
	}
}

func TestWithContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "with context",
			args: args{
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, WithContext(tt.args.ctx))
		})
	}
}

func Test_zapLogger_WithContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "with context",
			args: args{
				ctx: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(nil)
			assert.NotNil(t, l.WithContext(tt.args.ctx))
		})
	}
}
