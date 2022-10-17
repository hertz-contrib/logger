package logger

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/bytebufferpool"
	"strings"
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
	TagResBody           = "resBody"
	TagReqHeaders        = "reqHeaders"
	TagResHeaders        = "resHeaders"
	TagQueryStringParams = "queryParams"
	TagBody              = "body"
	TagBytesSent         = "bytesSent"
	TagBytesReceived     = "bytesReceived"
	TagRoute             = "route"
)

var Tags = map[string]logTagFunc{
	TagReferer: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return buf.WriteString(c.Request.Header.Get("Referer"))
	},
	TagProtocol: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return buf.WriteString(string(c.Request.URI().Scheme()))
	},
	TagPort: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		host := string(c.Request.URI().Host())
		split := strings.Split(host, ":")
		return buf.WriteString(split[1])
	},
	TagIP: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		host := string(c.Request.URI().Host())
		split := strings.Split(host, ":")
		return buf.WriteString(split[0])
	},
	TagIPs: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return buf.WriteString(c.Request.Header.Get("X-Forwarded-For"))
	},
	TagResBody: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return buf.WriteString(string(c.Response.Body()))
	},
	TagHost: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return buf.WriteString(string(c.Request.URI().Host()))
	},
	TagPath: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return buf.WriteString(string(c.Request.Path()))
	},
	TagURL: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return buf.WriteString(string(c.Request.Header.RequestURI()))
	},
	TagUA: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return buf.WriteString(c.Request.Header.Get("User-Agent"))
	},
	TagBody: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return buf.Write(c.Request.Body())
	},
	TagBytesSent: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return appendInt(buf, len(c.Response.Body()))
	},
	TagBytesReceived: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return appendInt(buf, len(c.Request.Body()))
	},
	TagRoute: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return buf.WriteString(string(c.Path()))
	},
	TagStatus: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return appendInt(buf, c.Response.StatusCode())
	},
	TagReqHeaders: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		reqHeaders := make([]string, 0)
		c.Request.Header.VisitAll(func(k, v []byte) {
			reqHeaders = append(reqHeaders, string(k)+"="+string(v))
		})
		return buf.Write([]byte(strings.Join(reqHeaders, "&")))
	},
	TagResHeaders: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		resHeaders := make([]string, 0)
		c.Response.Header.VisitAll(func(k, v []byte) {
			resHeaders = append(resHeaders, string(k)+"="+string(v))
		})
		return buf.Write([]byte(strings.Join(resHeaders, "&")))
	},
	TagQueryStringParams: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return buf.WriteString(c.Request.URI().QueryArgs().String())
	},
	TagMethod: func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return buf.WriteString(string(c.Method()))
	},
}
