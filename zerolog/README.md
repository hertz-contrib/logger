# Hertz zerolog
This is a logger library that uses zerolog to implement the [Hertz logger interface](https://www.cloudwego.io/docs/hertz/tutorials/framework-exten/log/)

## Usage

Download and install it:

```
go get github.com/hertz-contrib/logger/zerolog
```

Import it in your code:

```
import hertzZerolog "github.com/hertz-contrib/logger/zerolog"
```

Simple example:
```go
import (
    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/app/server"
    "github.com/cloudwego/hertz/pkg/common/hlog"
    "github.com/cloudwego/hertz/pkg/common/utils"
    "github.com/cloudwego/hertz/pkg/protocol/consts"

    hertzZerolog "github.com/hertz-contrib/logger/zerolog"
)

func main () {
    h := server.Default()
	
    hlog.SetLogger(hertzZerolog.New())

    h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
        hlog.Info("test log")
        c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
    })
	
    h.Spin()
}
```

### Options

Example:
```go
import (
    "os"
	
    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/app/server"
    "github.com/cloudwego/hertz/pkg/common/hlog"
    "github.com/cloudwego/hertz/pkg/common/utils"
    "github.com/cloudwego/hertz/pkg/protocol/consts"

    hertzZerolog "github.com/hertz-contrib/logger/zerolog"
)

func main () {
    h := server.Default()
	
    hlog.SetLogger(hertzZerolog.New(
        hertzZerolog.WithOutput(os.Stdout), // allows to specify output
        hertzZerolog.WithLevel(hlog.LevelWarn), // option with log level
	hertzZerolog.WithTimestamp(), // option with timestamp
	hertzZerolog.WithCaller())) // option with caller

    h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
        hlog.Info("test log")
        c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
    })
	
    h.Spin()
}
```

## Advanced usage

#### Implementing a request logging middleware:
```go
import (
    "context"
    "time"
    
    "github.com/cloudwego/hertz/pkg/app"

    hertzZerolog "github.com/hertz-contrib/logger/zerolog"
)

// RequestIDHeaderValue value for the request id header
const RequestIDHeaderValue = "X-Request-ID"

// LoggerMiddleware middleware for logging incoming requests
func LoggerMiddleware() app.HandlerFunc {
    return func(c context.Context, ctx *app.RequestContext) {
        start := time.Now()

        logger, err := hertzZerolog.GetLogger()
        if err != nil {
            hlog.Error(err)
            ctx.Next(c)
            return
        }

        reqId := c.Value(RequestIDHeaderValue).(string)
        if reqId != "" {
            logger = logger.WithField("request_id", reqId)
        }
        
        c = logger.WithContext(c)
        
        defer func() {
            stop := time.Now()
            
            logger.Unwrap().Info().
                Str("remote_ip", ctx.ClientIP()).
                Str("method", string(ctx.Method())).
                Str("path", string(ctx.Path())).
                Str("user_agent", string(ctx.UserAgent())).
                Int("status", ctx.Response.StatusCode()).
                Dur("latency", stop.Sub(start)).
                Str("latency_human", stop.Sub(start).String()).
                Msg("request processed")
        }()
        
        ctx.Next(c)
    }
}
```

