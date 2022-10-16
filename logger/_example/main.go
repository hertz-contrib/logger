package main

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/logger/logger"
)

func main() {
	h := server.Default(server.WithExitWaitTime(100 * time.Millisecond))
	h.Use(logger.NewLoggerMiddleware())
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.Request.Header.Set("test", "Hello fiber!")
		c.JSON(200, utils.H{"msg": "pong"})
	})
	h.Spin()
}
