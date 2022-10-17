package logger

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/bytebufferpool"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/cloudwego/hertz/pkg/route"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	hlog.SetOutput(buf)
	engine := route.NewEngine(config.NewOptions([]config.Option{}))
	engine.Use(NewLoggerMiddleware(WithFormat("${route}")))
	engine.GET("/", func(ctx context.Context, c *app.RequestContext) {})
	request := ut.PerformRequest(engine, "GET", "/", nil)
	w := request.Result()
	assert.DeepEqual(t, w.StatusCode(), 200)
	assert.DeepEqual(t, "/\n", buf.String()[len(buf.String())-2:])
}

func TestLoggerAll(t *testing.T) {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	hlog.SetOutput(buf)
	engine := route.NewEngine(config.NewOptions([]config.Option{}))
	engine.Use(NewLoggerMiddleware(WithFormat("${pid}${reqHeaders}${resHeaders}${referer}${protocol}${ip}${ips}" +
		"${host}${url}${ua}${body}${route}")))
	request := ut.PerformRequest(engine, "GET", "/?foo=bar", nil)
	w := request.Result()
	assert.DeepEqual(t, 404, w.StatusCode())
	assert.DeepEqual(t, fmt.Sprintf("%vContent-Type=text/plain; charset=utf-8http/?foo=bar/\n",
		os.Getpid()), buf.String()[49:])
}

func TestQueryParams(t *testing.T) {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	hlog.SetOutput(buf)
	engine := route.NewEngine(config.NewOptions([]config.Option{}))
	engine.Use(NewLoggerMiddleware(WithFormat("${queryParams}")))
	request := ut.PerformRequest(engine, "GET", "/?foo=bar&baz=moz", nil)
	result := request.Result()
	assert.DeepEqual(t, 404, result.StatusCode())
	assert.DeepEqual(t, "foo=bar&baz=moz\n", buf.String()[49:])
}

func TestRespBody(t *testing.T) {
	const getBody = "Sample response body"
	const postBody = "Post in test"
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	hlog.SetOutput(buf)
	engine := route.NewEngine(config.NewOptions([]config.Option{}))
	engine.Use(NewLoggerMiddleware(WithFormat("${resBody}")))
	engine.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.String(200, getBody)
	})
	engine.POST("/test", func(ctx context.Context, c *app.RequestContext) {
		c.String(200, postBody)
	})
	request := ut.PerformRequest(engine, "GET", "/", nil)
	w := request.Result()
	assert.DeepEqual(t, 200, w.StatusCode())
	assert.DeepEqual(t, getBody+"\n", buf.String()[len(buf.String())-len(getBody)-1:])
	buf.Reset()
	request = ut.PerformRequest(engine, "POST", "/test", nil)
	w = request.Result()
	assert.DeepEqual(t, 200, w.StatusCode())
	assert.DeepEqual(t, postBody+"\n", buf.String()[len(buf.String())-len(postBody)-1:])
}

func TestLoggerAppendUint(t *testing.T) {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	hlog.SetOutput(buf)
	engine := route.NewEngine(config.NewOptions([]config.Option{}))
	engine.Use(NewLoggerMiddleware(WithFormat("${bytesReceived} ${bytesSent} ${status}")))
	engine.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.String(200, "hello")
	})
	request := ut.PerformRequest(engine, "GET", "/", nil)
	w := request.Result()
	expected := "0 5 200"
	assert.DeepEqual(t, 200, w.StatusCode())
	assert.DeepEqual(t, expected+"\n", buf.String()[len(buf.String())-len(expected)-1:])
}
