package logger

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type (
	// options defines the config for middleware.
	options struct {
		format           string
		timeFormat       string
		timeZone         string
		timeInterval     time.Duration
		timeZoneLocation *time.Location
		enableLatency    bool
		accessLogFunc    func(ctx context.Context, format string, v ...interface{})
	}

	Option func(o *options)
)

func newOption(opts ...Option) *options {
	cfg := &options{
		format:        defaultFormat,
		timeFormat:    "15:04:05",
		timeZone:      "Local",
		timeInterval:  500 * time.Millisecond,
		accessLogFunc: hlog.CtxInfof,
	}

	// Return default config if nothing provided
	if len(opts) < 1 {
		return cfg
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

func WithFormat(s string) Option {
	return func(o *options) {
		o.format = s
	}
}

func WithTimeFormat(s string) Option {
	return func(o *options) {
		o.timeFormat = s
	}
}

func WithTimeInterval(t time.Duration) Option {
	return func(o *options) {
		o.timeInterval = t
	}
}

func WithAccessLogFunc(f func(ctx context.Context, format string, v ...interface{})) Option {
	return func(o *options) {
		o.accessLogFunc = f
	}
}
