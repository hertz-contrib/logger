## logger (This is a community driven project)

## Introduction

This middleware is used to [hertz](https://github.com/cloudwego/hertz) that logs HTTP request/response details.

## Usage

Download and install it:

```go
go get github.com/hertz-contrib/logger/logger
```

Import it in your code:

```go
import github.com/hertz-contrib/logger/logger
```

Simple Example:

```go
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
	h := server.Default(
		server.WithHostPorts("127.0.0.1:8080"),
		server.WithExitWaitTime(100*time.Millisecond),
	)
	h.Use(logger.NewLoggerMiddleware())
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(200, utils.H{"msg": "pong"})
	})
	h.Spin()
}
```

**make request**
```shell
curl --location --request GET 'http://127.0.0.1:8080/ping'
```
```shell
2022/10/17 00:00:14.302650 engine.go:572: [Debug] HERTZ: Method=GET    absolutePath=/ping                     --> handlerName=main.main.func1 (num=3 handlers)   
2022/10/17 00:00:14.302650 engine.go:366: [Info] HERTZ: Using network library=standard
2022/10/17 00:00:15.779222 logger.go:162: [Info]  00:00:15 | 200 |    11ms |  localhost:8080 | GET     | /ping
```
## Signatures

### NewLoggerMiddleware(opts ...Option) app.HandlerFunc

This function used to create a new middleware handler for logging HTTP Request/Response detail.

When no configuration changes are made, hertz prints the [default log messages](#Default Log Format) in the console in the default format

## 

## Log Format

### Default Log Format
```
${time} | ${status} | ${latency} | ${ip} | ${method} | ${path}
```
example
```go
2022/10/17 00:00:15.779222 logger.go:162: [Info]  00:00:15 | 200 |    11ms |  localhost:8080 | GET     | /ping
```
