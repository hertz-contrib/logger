# httplog  (This is a community driven project)

## Introduction

If you switch to Hertz from gin and prefer gin's HTTP logging style, this middleware will help you adapt.

## Effect display photo

![httplog_1](https://user-images.githubusercontent.com/78396698/201526334-c97bc174-d817-41d7-ab8f-1719690fdb8b.png)

## Usage

Download and install it:

```go
go get github.com/hertz-contrib/logger/httplog
go get github.com/mattn/go-colorable
```

Import it in your code:

```go
import github.com/hertz-contrib/logger/httplog
import github.com/mattn/go-colorable
```

Simple Example:

```go
package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/httplog"
	"github.com/mattn/go-colorable"
)

func main() {
    h := server.Default()
	httplog.ForceConsoleColor()
	httplog.DefaultWriter = colorable.NewColorableStdout()
	h.Use(httplog.Logger())
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})
	h.Spin()
}

```

## Config

### ForceConsoleColor

`httplog` provides `ForceConsoleColor` to set the mandatory use of color logging.

Simple Code:

```go
package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/httplog"
	"github.com/mattn/go-colorable"
)

func main() {
    h := server.Default()
	httplog.ForceConsoleColor()
	httplog.DefaultWriter = colorable.NewColorableStdout()
	h.Use(httplog.Logger())
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})
	h.Spin()
}
```

### DisableConsoleColor

`httplog` provides `DisableConsoleColor` to turn off color logging.

Sample Code:

```go
package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/httplog"
	"github.com/mattn/go-colorable"
)

func main() {
    h := server.Default()
	httplog.DisableConsoleColor()
	httplog.DefaultWriter = colorable.NewColorableStdout()
	h.Use(httplog.Logger())
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})
	h.Spin()
}
```

### DefaultWriter

`httplog` provides `DefaultWriter` to receive a new instance of `io.Writer`for handling standard output escape sequences. The third-party library `go-colorable` is used here to get the escape sequence of the standard output under the current operating system.

Sample Code:

```go
package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/httplog"
	"github.com/mattn/go-colorable"
)

func main() {
	h := server.Default()
    httplog.ForceConsoleColor()
	httplog.DefaultWriter = colorable.NewColorableStdout()
	h.Use(httplog.Logger())
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})
	h.Spin()
}
```

### LoggerWithWriter

`httplog` provides `LoggerWithWriter` for the user to specify the writer.

```go
package main

import (
	"context"
	"io"
	"os"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/httplog"
)

func main() {
	h := server.Default()
    f, _ := os.OpenFile("./hertz.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	httplog.ForceConsoleColor()
	h.Use(httplog.LoggerWithWriter(io.MultiWriter(f, os.Stdout)))
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})
	h.Spin()
}
```

### LoggerWithConfig

`httplog` provides `LoggerWithConfig` for creating a Logger middleware with user-defined `LoggerConfig` configuration.

```
package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/httplog"
	"os"
)

func main() {
	h := server.Default()
	httplog.ForceConsoleColor()
	f, _ := os.OpenFile("./hertz.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	h.Use(httplog.LoggerWithConfig(httplog.LoggerConfig{
		Output: f,
	}))
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})
	h.Spin()
}
```

### LoggerWithFormatter

`httplog ` provides `LoggerWithConfig` for user-defined log formats.

```go
package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/httplog"
	"time"
)

func main() {
	h := server.Default()
	httplog.ForceConsoleColor()
	h.Use(httplog.LoggerWithFormatter(myLogFormatter))
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})
	h.Spin()
}

// myLogFormatter
var myLogFormatter = func(param httplog.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}
	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}
	return fmt.Sprintf("[EduFriendChen] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}
```



## Log Format

### Default Log Format

```go
"[HERTZ] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage"
```

### Customize the log format Func

```go
// myLogFormatter
var myLogFormatter = func(param httplog.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}
	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}
	return fmt.Sprintf("[EduFriendChen] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}

//Usage
h.Use(httplog.LoggerWithFormatter(myLogFormatter))
```

