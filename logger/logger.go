package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/bytebufferpool"
	"github.com/hertz-contrib/logger/accesslog/internal/fasttemplate"
)

func NewLoggerMiddleware(opts ...Option) app.HandlerFunc {
	cfg := newOptions(opts...)

	// Get timezone location
	tz, err := time.LoadLocation(cfg.timeZone)
	if err != nil || tz == nil {
		cfg.timeZoneLocation = time.Local
	} else {
		cfg.timeZoneLocation = tz
	}
	// Check if format contains latency
	cfg.enableLatency = strings.Contains(cfg.format, cfg.withDelims("latency"))

	tmpl := fasttemplate.New(cfg.format, cfg.leftDelim, cfg.rightDelim)

	// Create correct timeformat
	var timestamp string

	// Update date/time every 750 milliseconds in a separate go routine
	if strings.Contains(cfg.format, cfg.withDelims("time")) {
		go func() {
			for {
				time.Sleep(cfg.timeInterval)
				timestamp = time.Now().In(cfg.timeZoneLocation).Format(cfg.timeFormat)
			}
		}()
	}

	// Set PID once
	pid := strconv.Itoa(os.Getpid())

	return func(ctx context.Context, c *app.RequestContext) {
		var start, stop time.Time

		// Set latency start time
		if cfg.enableLatency {
			start = time.Now()
		}
		c.Next(ctx)

		if cfg.enableLatency {
			stop = time.Now()
		}

		// Get new buffer
		buf := bytebufferpool.Get()
		defer bytebufferpool.Put(buf)

		Tags[TagPid] = func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
			return buf.WriteString(pid)
		}
		Tags[TagTime] = func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
			return buf.WriteString(timestamp)
		}
		Tags[TagLatency] = func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
			return buf.WriteString(fmt.Sprintf("%7v", stop.Sub(start).Round(time.Millisecond)))
		}

		if cfg.format == defaultFormat {
			// format log to buffer
			_, _ = buf.WriteString(fmt.Sprintf(" %s | %3d | %7v | %15s | %-7s | %-s ",
				timestamp,
				c.Response.StatusCode(),
				stop.Sub(start).Round(time.Millisecond),
				c.Request.URI().Host(),
				c.Method(),
				c.Path(),
			))

			cfg.accessLogFunc(ctx, buf.String())
			return
		}

		_, err := tmpl.ExecuteFunc(buf, func(w io.Writer, tag string) (int, error) {
			if function, ok := Tags[tag]; ok {
				return function(ctx, c, buf)
			}
			return 0, nil
		})
		// Also write errors to the buffer
		if err != nil {
			_, _ = buf.WriteString(err.Error())
		}

		cfg.accessLogFunc(ctx, buf.String())
	}
}

func appendInt(buf *bytebufferpool.ByteBuffer, v int) (int, error) {
	old := len(buf.B)
	buf.B = appendUint(buf.B, v)
	return len(buf.B) - old, nil
}

func appendUint(dst []byte, n int) []byte {
	if n < 0 {
		panic("BUG: int must be positive")
	}

	var b [20]byte
	buf := b[:]
	i := len(buf)
	var q int
	for n >= 10 {
		i--
		q = n / 10
		buf[i] = '0' + byte(n-q*10)
		n = q
	}
	i--
	buf[i] = '0' + byte(n)

	dst = append(dst, buf[i:]...)
	return dst
}
