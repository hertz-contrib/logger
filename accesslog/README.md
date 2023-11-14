# accesslog (This is a community driven project)

## Introduction

This middleware is used to [hertz](https://github.com/cloudwego/hertz) that logs HTTP request/response details and inspired by [logger](https://github.com/gofiber/fiber/tree/master/middleware/logger).

## Usage

Download and install it:

```go
go get github.com/hertz-contrib/logger/accesslog
```

Import it in your code:

```go
import github.com/hertz-contrib/logger/accesslog
```

Simple Example:

```go
package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/logger/accesslog"
)

func main() {
	h := server.Default(
		server.WithHostPorts(":8080"),
	)
	h.Use(accesslog.New())
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(200, utils.H{"msg": "pong"})
	})
	h.Spin()
}


```

## Config

### WithFormat

The `accesslog` provides `WithFormat` to help users set the format of the log.

Simple Code:

```go
package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/logger/accesslog"
)

func main() {
	h := server.Default(
		server.WithHostPorts(":8080"),
	)
	h.Use(accesslog.New(accesslog.WithFormat("[${time}] ${status} - ${latency} ${method} ${path} ${queryParams}")))
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(200, utils.H{"msg": "pong"})
	})
	h.Spin()
}

```

### WithTimeFormat

The `accesslog` provides `WithTimeFormat` to help users set the format of the `time`.

Sample Code:

```go
package main

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/logger/accesslog"
)

func main() {
	h := server.Default(
		server.WithHostPorts(":8080"),
	)
	h.Use(accesslog.New(
		accesslog.WithTimeFormat(time.RFC822),
	))
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(200, utils.H{"msg": "pong"})
	})
	h.Spin()
}
```

### WithTimeInterval

The `accesslog` provides `WithTimeInterval` to help the user set the update interval of the timestamp.

Sample Code:

```go
package main

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/logger/accesslog"
)

func main() {
	h := server.Default(
		server.WithHostPorts(":8080"),
	)
	h.Use(accesslog.New(
		accesslog.WithTimeInterval(time.Second),
	))
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(200, utils.H{"msg": "pong"})
	})
	h.Spin()
}
```

### WithAccessLogFunc

The `accesslog` provides `WithAccessLogFunc` to help users set the log printing functions.

Sample Code:

```go
package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/logger/accesslog"
)

func main() {
	h := server.Default(
		server.WithHostPorts(":8080"),
	)
	h.Use(accesslog.New(
		accesslog.WithAccessLogFunc(hlog.CtxInfof),
	))
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(200, utils.H{"msg": "pong"})
	})
	h.Spin()
}
```

## Log Format

### Default Log Format
```
[${time}] ${status} - ${latency} ${method} ${path}
```
example
```go
[21:54:36] 200 - 2.906859ms GET /ping
```

### Supported tags

```go
const (
	TagPid               = "pid"
	TagTime              = "time"
	TagReferer           = "referer"
	TagProtocol          = "protocol"
	TagPort              = "port"
	TagIP                = "ip"
	TagIPs               = "ips"
	TagClientIP          = "clientIP"
	TagHost              = "host"
	TagMethod            = "method"
	TagPath              = "path"
	TagURL               = "url"
	TagUA                = "ua"
	TagLatency           = "latency"
	TagStatus            = "status"       // response status
	TagResBody           = "resBody"      // response body
	TagReqHeaders        = "reqHeaders"
	TagResHeaders        = "resHeaders"
	TagQueryStringParams = "queryParams"  // request query parameters
	TagBody              = "body"         // request body
	TagBytesSent         = "bytesSent"
	TagBytesReceived     = "bytesReceived"
	TagRoute             = "route"        // request path
)
```


### Custom Tag

We can add custom tags to the [accesslog.Tags](tag.go), but please note that it is not thread-safe.

Sample Code:
```go
package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/bytebufferpool"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/logger/accesslog"
)

func main() {
	accesslog.Tags["test_tag"] = func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return buf.WriteString("test")
	}
	h := server.Default(
		server.WithHostPorts(":8080"),
	)
	h.Use(accesslog.New(accesslog.WithFormat("${test_tag}")))
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(200, utils.H{"msg": "pong"})
	})
	h.Spin()
}

```

### Log By Condition

We can add a method `logConditionFunc` to determine whether to log based on the conditions.

Sample Code:

```go
package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/logger/accesslog"
)

func main() {
	h := server.Default(server.WithHostPorts(":8081"))
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, map[string]string{"ping": "pong"})
	})
	h.Use(accesslog.New(accesslog.WithLogConditionFunc(func(ctx context.Context, c *app.RequestContext) bool {
		if c.FullPath() == "/ping" {
			return false
		} else {
			return true
		}
	})))
	h.Spin()
}
```
