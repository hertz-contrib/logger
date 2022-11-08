# Hertz zerolog
This is a logger library that uses zerolog to implement the [Hertz logger interface](https://www.cloudwego.io/docs/hertz/tutorials/framework-exten/log/)

## Usage

Download and install it:

    go get github.com/hertz-contrib/logger/zerolog

Import it in your code:

    import hertzZerolog "github.com/hertz-contrib/logger/zerolog"

Simple example:
```go
import (
    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/app/server"
    "github.com/cloudwego/hertz/pkg/common/hlog"

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

Options:
```go
import (
    "os"
	
    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/app/server"
    "github.com/cloudwego/hertz/pkg/common/hlog"

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
