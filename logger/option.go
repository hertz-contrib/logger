package logger

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/bytebufferpool"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type (
	// options defines the config for middleware.
	options struct {
		format           string
		leftDelim        string
		rightDelim       string
		timeFormat       string
		timeZone         string
		timeInterval     time.Duration
		timeZoneLocation *time.Location
		enableLatency    bool
		accessLogFunc    func(ctx context.Context, format string, v ...interface{})
	}

	Option     func(o *options)
	logTagFunc func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error)
)

var defaultFormat = "[${time}] ${status} - ${latency} ${method} ${path}"

func newOptions(opts ...Option) *options {
	cfg := &options{
		format:        defaultFormat,
		leftDelim:     "${",
		rightDelim:    "}",
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

func WithLogTagFunc(tag string, f logTagFunc) Option {
	return func(o *options) {
		Tags[tag] = f
	}
}

func WithDelims(left, right string) Option {
	return func(o *options) {
		o.leftDelim = left
		o.rightDelim = right
		defaultFormat = fmt.Sprintf("%s %s %s %s %s",
			o.withDelims(TagTime), o.withDelims(TagStatus), o.withDelims(TagLatency),
			o.withDelims(TagMethod), o.withDelims(TagPath),
		)
	}
}

func (o *options) withDelims(content string) string {
	return o.leftDelim + content + o.rightDelim
}
