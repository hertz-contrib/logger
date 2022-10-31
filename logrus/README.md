# Hertz Logrus (This is a community driven project)

## Introduction

This is a logger library that uses logrus to implement the [Hertz logger interface](https://www.cloudwego.io/docs/hertz/tutorials/framework-exten/log/)

## Usage

Download and install it:

```go
go get github.com/hertz-contrib/logger/logrus
```

Import it in your code:

```go
import hertzlogrus github.com/hertz-contrib/logger/logrus
```


Simple Example:

```go
package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
)

func main() {
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)

	...

	hlog.CtxInfof(context.Background(), "hello %s", "hertz")
}
```

