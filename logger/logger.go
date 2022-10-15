package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/bytebufferpool"
	"github.com/hertz-contrib/logger/logger/internal/template"
)

// Logger variables
const (
	TagPid               = "pid"
	TagTime              = "time"
	TagReferer           = "referer"
	TagProtocol          = "protocol"
	TagPort              = "port"
	TagIP                = "ip"
	TagIPs               = "ips"
	TagHost              = "host"
	TagMethod            = "method"
	TagPath              = "path"
	TagURL               = "url"
	TagUA                = "ua"
	TagLatency           = "latency"
	TagStatus            = "status"
	TagReqHeaders        = "reqHeaders"
	TagQueryStringParams = "queryParams"
	TagBytesSent         = "bytesSent"
	TagBytesReceived     = "bytesReceived"
	TagRoute             = "route"

	// DEPRECATED: Use TagReqHeader instead
	TagHeader     = "header:"
	TagReqHeader  = "reqHeader:"
	TagRespHeader = "respHeader:"
	TagQuery      = "query:"
	TagForm       = "form:"
	TagCookie     = "cookie:"
)

func NewLogger(opts ...Option) app.HandlerFunc {
	cfg := newOption(opts...)

	// Get timezone location
	tz, err := time.LoadLocation(cfg.TimeZone)
	if err != nil || tz == nil {
		cfg.timeZoneLocation = time.Local
	} else {
		cfg.timeZoneLocation = tz
	}
	// Check if format contains latency
	cfg.enableLatency = strings.Contains(cfg.Format, "${latency}")

	tmpl := template.New(cfg.Format, "${", "}")

	// Create correct timeformat
	var timestamp atomic.Value
	timestamp.Store(time.Now().In(cfg.timeZoneLocation).Format(cfg.TimeFormat))

	// Update date/time every 750 milliseconds in a separate go routine
	if strings.Contains(cfg.Format, "${time}") {
		go func() {
			for {
				time.Sleep(cfg.TimeInterval)
				timestamp.Store(time.Now().In(cfg.timeZoneLocation).Format(cfg.TimeFormat))
			}
		}()
	}

	// Set PID once
	pid := strconv.Itoa(os.Getpid())

	// Set variables
	//var mu sync.Mutex

	errPadding := 15
	errPaddingStr := strconv.Itoa(errPadding)

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

		if cfg.Format == defaultFormat {
			formatErr := ""
			// Format log to buffer
			_, _ = buf.WriteString(fmt.Sprintf(" %s | %3d | %7v | %15s | %-7s | %-"+errPaddingStr+"s %s\n",
				timestamp.Load().(string),
				c.Response.StatusCode(),
				stop.Sub(start).Round(time.Millisecond),
				c.Request.URI().Host(),
				c.Method(),
				c.Path(),
				formatErr,
			))

			// Write buffer to output
			//_, _ = cfg.Output.Write(buf.Bytes())
			//hlog.CtxInfof(ctx, buf.String())
			cfg.outFunc(ctx, buf.String())
			// Put buffer back to pool
			bytebufferpool.Put(buf)
			return
		}

		_, err := tmpl.ExecuteFunc(buf, func(w io.Writer, tag string) (int, error) {
			switch tag {
			case TagPid:
				return buf.WriteString(pid)
			case TagTime:
				return buf.WriteString(timestamp.Load().(string))
			case TagReferer:
				return buf.WriteString(c.Request.Header.Get("Referer"))
			case TagProtocol:
				return buf.WriteString(string(c.Request.URI().Scheme()))
			case TagPort:
				host := string(c.Request.URI().Host())
				split := strings.Split(host, ":")
				return buf.WriteString(split[1])
			case TagIP:
				host := string(c.Request.URI().Host())
				split := strings.Split(host, ":")
				return buf.WriteString(split[0])
			case TagIPs:
				return buf.WriteString(c.Request.Header.Get("X-Forwarded-For"))
			case TagHost:
				return buf.WriteString(string(c.Request.URI().Host()))
			case TagPath:
				return buf.WriteString(string(c.Request.Path()))
			case TagURL:
				return buf.WriteString(string(c.Request.Header.RequestURI()))
			case TagUA:
				return buf.WriteString(c.Request.Header.Get("User-Agent"))
			case TagLatency:
				return buf.WriteString(fmt.Sprintf("%7v", stop.Sub(start).Round(time.Millisecond)))
			case TagBytesReceived:
				return appendInt(buf, len(c.Request.Body()))
			case TagBytesSent:
				return appendInt(buf, len(c.Response.Body()))
			case TagRoute:
				return buf.WriteString(string(c.Path()))
			case TagStatus:
				return appendInt(buf, c.Response.StatusCode())
			case TagReqHeaders:
				reqHeaders := make([]string, 0)
				c.Request.Header.VisitAll(func(k, v []byte) {
					reqHeaders = append(reqHeaders, string(k)+"="+string(v))
				})
				return buf.Write([]byte(strings.Join(reqHeaders, "&")))
			case TagQueryStringParams:
				return buf.WriteString(c.Request.URI().QueryArgs().String())
			case TagMethod:
				return buf.WriteString(string(c.Method()))
			default:
				// Check if we have a value tag i.e.: "reqHeader:x-key"
				switch {
				case strings.HasPrefix(tag, TagReqHeader):
					return buf.WriteString(c.Request.Header.Get(tag[10:]))
				case strings.HasPrefix(tag, TagRespHeader):
					return buf.WriteString(c.Response.Header.Get(tag[11:]))
				case strings.HasPrefix(tag, TagQuery):
					return buf.WriteString(c.Query(tag[6:]))
				case strings.HasPrefix(tag, TagForm):
					return buf.WriteString(string(c.FormValue(tag[5:])))
				case strings.HasPrefix(tag, TagCookie):
					return buf.WriteString(string(c.Cookie(tag[7:])))
				}
			}
			return 0, nil
		})
		// Also write errors to the buffer
		if err != nil {
			_, _ = buf.WriteString(err.Error())
		}
		// Write buffer to output

		//if _, err := cfg.Output.Write(buf.Bytes()); err != nil {
		//	// Write error to output
		//	if _, err := cfg.Output.Write([]byte(err.Error())); err != nil {
		//		// There is something wrong with the given io.Writer
		//		fmt.Fprintf(os.Stderr, "Failed to write to log, %v\n", err)
		//	}
		//}
		//hlog.CtxInfof(ctx, buf.String())
		cfg.outFunc(ctx, buf.String())
		// Put buffer back to pool
		bytebufferpool.Put(buf)
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
