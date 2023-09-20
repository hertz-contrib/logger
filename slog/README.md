# Hertz Slog (This is a community driven project)

## Introduction

This is a logger library that uses slog to implement the [Hertz logger interface](https://www.cloudwego.io/docs/hertz/tutorials/framework-exten/log/)

## Usage

Download and install it:

```go
go get github.com/hertz-contrib/logger/slog
```

Import it in your code:

```go
import hertzslog "github.com/hertz-contrib/logger/slog"
```


Simple Example:

```go
package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzslog "github.com/hertz-contrib/logger/slog"
)

func main() {
	logger := hertzslog.NewLogger()
	hlog.SetLogger(logger)

	...

	hlog.CtxInfof(context.Background(), "hello %s", "hertz")
}
```

