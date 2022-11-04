# Hertz zap (This is a community driven project)

## Introduction

This is a logger library that uses [zap](https://github.com/uber-go/zap) to implement the [Hertz logger interface](https://www.cloudwego.io/docs/hertz/tutorials/framework-exten/log/)

## Usage

Download and install it:

```go
go get github.com/hertz-contrib/logger/zap
```

Import it in your code:

```go
import hertzzap github.com/hertz-contrib/logger/
```

Simple Example:

```go
package main

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"
)

func main() {
	logger := hertzzap.NewLogger()
	hlog.SetLogger(logger)

	// ...

	hlog.Infof("hello %s", "hertz")
}
```

> For some reason, zap don't support log context info yet.
