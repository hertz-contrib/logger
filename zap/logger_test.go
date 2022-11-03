// Copyright 2022 CloudWeGo Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zap

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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

// humanEncoderConfig copy from zap
func humanEncoderConfig() zapcore.EncoderConfig {
	cfg := testEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncodeDuration = zapcore.StringDurationEncoder
	return cfg
}

// TestLogger test logger work with hertz
func TestLogger(t *testing.T) {
	buf := new(bytes.Buffer)

	logger := NewLogger(WithZapOptions(zap.WithFatalHook(zapcore.WriteThenPanic)))
	defer logger.Sync()

	hlog.SetLogger(logger)
	hlog.SetOutput(buf)
	hlog.SetLevel(hlog.LevelDebug)

	type logMap map[string]string

	logTestSlice := []logMap{
		{
			"logMessage":       "this is a trace log",
			"formatLogMessage": "this is a trace log: %s",
			"logLevel":         "Trace",
			"zapLogLevel":      "debug",
		},
		{
			"logMessage":       "this is a debug log",
			"formatLogMessage": "this is a debug log: %s",
			"logLevel":         "Debug",
			"zapLogLevel":      "debug",
		},
		{
			"logMessage":       "this is a info log",
			"formatLogMessage": "this is a info log: %s",
			"logLevel":         "Info",
			"zapLogLevel":      "info",
		},
		{
			"logMessage":       "this is a notice log",
			"formatLogMessage": "this is a notice log: %s",
			"logLevel":         "Notice",
			"zapLogLevel":      "warn",
		},
		{
			"logMessage":       "this is a warn log",
			"formatLogMessage": "this is a warn log: %s",
			"logLevel":         "Warn",
			"zapLogLevel":      "warn",
		},
		{
			"logMessage":       "this is a error log",
			"formatLogMessage": "this is a error log: %s",
			"logLevel":         "Error",
			"zapLogLevel":      "error",
		},
		{
			"logMessage":       "this is a fatal log",
			"formatLogMessage": "this is a fatal log: %s",
			"logLevel":         "Fatal",
			"zapLogLevel":      "fatal",
		},
	}

	testHertzLogger := reflect.ValueOf(logger)

	for _, v := range logTestSlice {
		t.Run(v["logLevel"], func(t *testing.T) {
			if v["logLevel"] == "Fatal" {
				defer func() {
					assert.Equal(t, "this is a fatal log", recover())
				}()
			}
			logFunc := testHertzLogger.MethodByName(v["logLevel"])
			logFunc.Call([]reflect.Value{
				reflect.ValueOf(v["logMessage"]),
			})
			assert.Contains(t, buf.String(), v["logMessage"])
			assert.Contains(t, buf.String(), v["zapLogLevel"])

			buf.Reset()

			logfFunc := testHertzLogger.MethodByName(fmt.Sprintf("%sf", v["logLevel"]))
			logfFunc.Call([]reflect.Value{
				reflect.ValueOf(v["formatLogMessage"]),
				reflect.ValueOf(v["logLevel"]),
			})
			assert.Contains(t, buf.String(), fmt.Sprintf(v["formatLogMessage"], v["logLevel"]))
			assert.Contains(t, buf.String(), v["zapLogLevel"])

			buf.Reset()

			ctx := context.Background()
			ctxLogfFunc := testHertzLogger.MethodByName(fmt.Sprintf("Ctx%sf", v["logLevel"]))
			ctxLogfFunc.Call([]reflect.Value{
				reflect.ValueOf(ctx),
				reflect.ValueOf(v["formatLogMessage"]),
				reflect.ValueOf(v["logLevel"]),
			})
			assert.Contains(t, buf.String(), fmt.Sprintf(v["formatLogMessage"], v["logLevel"]))
			assert.Contains(t, buf.String(), v["zapLogLevel"])

			buf.Reset()
		})
	}
}

// TestLogLevel test SetLevel
func TestLogLevel(t *testing.T) {
	buf := new(bytes.Buffer)

	logger := NewLogger()
	defer logger.Sync()

	// output to buffer
	logger.SetOutput(buf)

	logger.Debug("this is a debug log")
	assert.False(t, strings.Contains(buf.String(), "this is a debug log"))

	logger.SetLevel(hlog.LevelDebug)

	logger.Debugf("this is a debug log %s", "msg")
	assert.True(t, strings.Contains(buf.String(), "this is a debug log"))
}

// TestCoreOption test zapcore config option
func TestCoreOption(t *testing.T) {
	buf := new(bytes.Buffer)

	logger := NewLogger(
		WithCoreEnc(zapcore.NewConsoleEncoder(humanEncoderConfig())),
		WithCoreLevel(zap.NewAtomicLevelAt(zapcore.WarnLevel)),
		WithCoreWs(zapcore.AddSync(buf)),
	)
	defer logger.Sync()

	logger.SetOutput(buf)

	logger.Debug("this is a debug log")
	// test log level
	assert.False(t, strings.Contains(buf.String(), "this is a debug log"))

	logger.Error("this is a warn log")
	// test log level
	assert.True(t, strings.Contains(buf.String(), "this is a warn log"))
	// test console encoder result
	assert.True(t, strings.Contains(buf.String(), "\tERROR\t"))
}

// TestCoreOption test zapcore config option
func TestZapOption(t *testing.T) {
	buf := new(bytes.Buffer)

	logger := NewLogger(
		WithZapOptions(zap.AddCaller()),
	)
	defer logger.Sync()

	logger.SetOutput(buf)

	logger.Debug("this is a debug log")
	assert.False(t, strings.Contains(buf.String(), "this is a debug log"))

	logger.Error("this is a warn log")
	// test caller in log result
	assert.True(t, strings.Contains(buf.String(), "caller"))
}
