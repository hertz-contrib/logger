package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/bytebufferpool"
	logger "github.com/hertz-contrib/logger/accesslog"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func main() {
	h := server.Default(
		server.WithHostPorts(":8080"),
		server.WithExitWaitTime(100*time.Millisecond),
	)
	h.Use(logger.NewLoggerMiddleware(
		logger.WithDelims("{{", "}}"),
		logger.WithLogTagFunc("t", func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
			return buf.WriteString(time.Now().String())
		}),
		logger.WithFormat("{{t}} | {{ip}}"),
	))
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(200, utils.H{"msg": "pong"})
	})
	h.Spin()
}
