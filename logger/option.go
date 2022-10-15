package logger

import (
	"context"
	"time"
)

const defaultFormat = "${time} ${status} ${latency} ${method} ${path}\n"

// options defines the config for middleware.
type (
	options struct {
		Format           string
		TimeFormat       string
		TimeZone         string
		TimeInterval     time.Duration
		enableLatency    bool
		outFunc          func(ctx context.Context, format string, v ...interface{})
		timeZoneLocation *time.Location
	}

	Option func(o *options)
)

func newOption(opts ...Option) *options {
	cfg := &options{
		Format:       defaultFormat,
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		//Output:       os.Stdout,
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
		o.Format = s
	}
}

func WithTimeFormat(s string) Option {
	return func(o *options) {
		o.TimeFormat = s
	}
}

func WithTimeInterval(t time.Duration) Option {
	return func(o *options) {
		o.TimeInterval = t
	}
}

func WithOutputFunc(f func(ctx context.Context, format string, v ...interface{})) Option {
	return func(o *options) {
		o.outFunc = f
	}
}
