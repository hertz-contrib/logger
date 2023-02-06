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

Multiple `zapcore` Example:

```go
package main

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	dynamicLevel := zap.NewAtomicLevel()

	dynamicLevel.SetLevel(zap.DebugLevel)

	logger := hertzzap.NewLogger(
		hertzzap.WithCores([]hertzzap.CoreConfig{
			{
				Enc: zapcore.NewConsoleEncoder(humanEncoderConfig()),
				Ws:  os.Stdout,
				Lvl: dynamicLevel,
			},
			{
				Enc: zapcore.NewJSONEncoder(humanEncoderConfig()),
				Ws:  getWriteSyncer("./all/log.log"),
				Lvl: zap.NewAtomicLevelAt(zapcore.DebugLevel),
			},
			{
				Enc: zapcore.NewJSONEncoder(humanEncoderConfig()),
				Ws:  getWriteSyncer("./debug/log.log"),
				Lvl: zap.NewAtomicLevelAt(zapcore.LevelOf(
					zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
						return lev == zap.DebugLevel
					}))),
			},
			{
				Enc: zapcore.NewJSONEncoder(humanEncoderConfig()),
				Ws:  getWriteSyncer("./info/log.log"),
				Lvl: zap.NewAtomicLevelAt(zapcore.LevelOf(
					zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
						return lev == zap.InfoLevel
					}))),
			},
			{
				Enc: zapcore.NewJSONEncoder(humanEncoderConfig()),
				Ws:  getWriteSyncer("./warn/log.log"),
				Lvl: zap.NewAtomicLevelAt(zapcore.LevelOf(
					zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
						return lev == zap.WarnLevel
					}))),
			},
			{
				Enc: zapcore.NewJSONEncoder(humanEncoderConfig()),
				Ws:  getWriteSyncer("./error/log.log"),
				Lvl: zap.NewAtomicLevelAt(zapcore.LevelOf(
					zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
						return lev >= zap.ErrorLevel
					}))),
			},
		}...),
	)
	defer logger.Sync()

	hlog.Infof("hello %s", "hertz")
}

// humanEncoderConfig copy from zap
func humanEncoderConfig() zapcore.EncoderConfig {
	cfg := testEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncodeDuration = zapcore.StringDurationEncoder
	return cfg
}

func getWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file,
		MaxSize:    10,
		MaxBackups: 50000,
		MaxAge:     1000,
		Compress:   true,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// testEncoderConfig encoder config for testing, copy from zap
func testEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "name",
		TimeKey:        "ts",
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     "\n",
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
```

## Attention

Method `SetLevel()` and `SetOutput()` only affect the first `zapcore` if imported multiple `zapcore`


> For some reason, zap don't support log context info yet.
