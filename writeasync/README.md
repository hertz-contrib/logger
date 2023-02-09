# Hertz writeasync (This is a community driven project)

## Introduction

This is a library of tools that implement asynchronous logging

## Usage

Download and install it:

```go
go get github.com/hertz-contrib/logger/writeasync
```

Import it in your code:

```go
import github.com/hertz-contrib/logger/writeasync
```

Simple Example:

```go
package main

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/logger/writeasync"
	"os"
)

func main() {
	f, err := os.OpenFile("./output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    wa := writeasync.NewWriterAsync(f)
    defer wa.Close()

    hlog.SetOutput(wa)

	// ...

	hlog.Infof("hello %s", "hertz")
}
```

Other logger Examples:

`zap`

```go
package main

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/logger/writeasync"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
    wa := writeasync.NewWriterAsync(getWriteSyncer("./output.log"))
    defer wa.Close()

    logger := hertzzap.NewLogger(
		hertzzap.WithCores([]hertzzap.CoreConfig{
			{
				Enc: zapcore.NewJSONEncoder(humanEncoderConfig()),
				Ws:  zapcore.AddSync(wa),
				Lvl: zap.NewAtomicLevelAt(zapcore.DebugLevel),
			},
		}...),
	)

	hlog.SetLogger(logger)

	// ...

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
	//return zapcore.AddSync(lumberJackLogger)
	return lumberJackLogger
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

`zerolog`

```go
package main

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzZerolog "github.com/hertz-contrib/logger/zerolog"
)

func main() {
	wa := writeasync.NewWriterAsync(os.Stdout)
    defer wa.Close()

	hlog.SetLogger(hertzZerolog.New(
        hertzZerolog.WithOutput(wa), // allows to specify output
        hertzZerolog.WithLevel(hlog.LevelInfo), // option with log level
		hertzZerolog.WithTimestamp(), // option with timestamp
		hertzZerolog.WithCaller())) // option with caller

	// ...

	hlog.Infof("hello %s", "hertz")
}
```

## Performance

|  #     | 框架 | 场景 | Requests/sec |
|  ----  | ---- | ---- | ----------- |
| 1  | default | 同步写文件 | 93107.46 |
| 2  | default | 异步写文件 | 187069.88 |
| 3  | default | 同步控制台输出 | 20150.08 |
| 4  | default | 异步控制台输出 | 72374.32 |
| 5  | zap | 同步写文件 | 120954.57 |
| 6  | zap | 异步写文件 | 210114.05 |
| 7  | zap | 同步控制台输出 | 20972.39 |
| 8  | zap | 异步控制台输出 | 64265.20 |
| 9  | zerolog | 同步写文件 | 52941.38 |
| 10  | zerolog | 异步写文件 | 199147.68 |
| 11  | zerolog | 同步控制台输出 | 17475.38 |
| 12  | zerolog | 异步控制台输出 | 67763.88 |

```shell
[default]
同步写文件
$ wrk -t 5 -c 1000 http://192.168.1.131:8888/get
Running 10s test @ http://192.168.1.131:8888/get
  5 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    11.09ms    8.53ms  71.86ms   62.05%
    Req/Sec    18.87k   832.22    20.73k    74.00%
  938348 requests in 10.08s, 163.76MB read
Requests/sec:  93107.46
Transfer/sec:     16.25MB

异步写文件
$ wrk -t 5 -c 1000 http://192.168.1.131:8888/get
Running 10s test @ http://192.168.1.131:8888/get
  5 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     7.14ms    8.28ms 109.47ms   86.73%
    Req/Sec    38.04k     4.31k   59.59k    81.56%
  1889030 requests in 10.10s, 329.68MB read
Requests/sec: 187069.88
Transfer/sec:     32.65MB

同步控制台输出
$ wrk -t 5 -c 1000 http://192.168.1.131:8888/get
Running 10s test @ http://192.168.1.131:8888/get
  5 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    47.12ms   14.15ms 163.86ms   71.29%
    Req/Sec     4.09k     0.98k    6.53k    74.95%
  203087 requests in 10.08s, 35.44MB read
Requests/sec:  20150.08
Transfer/sec:      3.52MB

异步控制台输出
$ wrk -t 5 -c 1000 http://192.168.1.131:8888/get
Running 10s test @ http://192.168.1.131:8888/get
  5 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    16.56ms   16.00ms 122.66ms   53.28%
    Req/Sec    14.61k    14.41k   44.45k    70.00%
  726929 requests in 10.04s, 126.87MB read
Requests/sec:  72374.32
Transfer/sec:     12.63MB


[zap]
同步写文件
wrk -t 5 -c 1000 http://192.168.1.131:8888/get
Running 10s test @ http://192.168.1.131:8888/get
  5 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     8.13ms    4.99ms  67.06ms   71.90%
    Req/Sec    24.66k     1.73k   40.05k    89.56%
  1221613 requests in 10.10s, 213.20MB read
Requests/sec: 120954.57
Transfer/sec:     21.11MB

异步写文件
$ wrk -t 5 -c 1000 http://192.168.1.131:8888/get
Running 10s test @ http://192.168.1.131:8888/get
  5 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.52ms    4.73ms 230.22ms   94.59%
    Req/Sec    42.46k     7.25k   66.28k    68.89%
  2105784 requests in 10.02s, 367.51MB read
Requests/sec: 210114.05
Transfer/sec:     36.67MB

同步控制台输出
$ wrk -t 5 -c 1000 http://192.168.1.131:8888/get
Running 10s test @ http://192.168.1.131:8888/get
  5 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    49.54ms   36.16ms 412.63ms   67.11%
    Req/Sec     4.25k   662.07     5.86k    70.20%
  211719 requests in 10.10s, 36.95MB read
Requests/sec:  20972.39
Transfer/sec:      3.66MB

异步控制台输出
$ wrk -t 5 -c 1000 http://192.168.1.131:8888/get
Running 10s test @ http://192.168.1.131:8888/get
  5 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    37.59ms   50.29ms 350.01ms   81.38%
    Req/Sec    12.98k    18.26k   61.13k    77.78%
  645069 requests in 10.04s, 112.58MB read
Requests/sec:  64265.20
Transfer/sec:     11.22MB


[zerolog]
同步控制台输出
$ wrk -t 5 -c 1000 http://192.168.1.131:8888/get
Running 10s test @ http://192.168.1.131:8888/get
  5 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    58.59ms   41.95ms 459.14ms   62.42%
    Req/Sec     3.54k   320.46     4.39k    71.20%
  176284 requests in 10.09s, 30.77MB read
Requests/sec:  17475.38
Transfer/sec:      3.05MB

异步控制台输出
$ wrk -t 5 -c 1000 http://192.168.1.131:8888/get
Running 10s test @ http://192.168.1.131:8888/get
  5 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    23.51ms   27.49ms 276.24ms   81.59%
    Req/Sec    13.68k    16.93k   63.26k    76.97%
  681006 requests in 10.05s, 118.85MB read
Requests/sec:  67763.88
Transfer/sec:     11.83MB

同步写文件
$ wrk -t 5 -c 1000 http://192.168.1.131:8888/get
Running 10s test @ http://192.168.1.131:8888/get
  5 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    21.41ms   19.30ms 257.22ms   61.30%
    Req/Sec    10.73k     2.73k   15.51k    57.20%
  533871 requests in 10.08s, 93.17MB read
Requests/sec:  52941.38
Transfer/sec:      9.24MB

异步写文件
$ wrk -t 5 -c 1000 http://192.168.1.131:8888/get
Running 10s test @ http://192.168.1.131:8888/get
  5 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.19ms    6.47ms 223.09ms   95.70%
    Req/Sec    40.30k     7.79k   82.50k    75.50%
  2008631 requests in 10.09s, 350.55MB read
Requests/sec: 199147.68
Transfer/sec:     34.76MB
```
