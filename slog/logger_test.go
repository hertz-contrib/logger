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

package slog

import (
	"bufio"
	"bytes"
	"log/slog"
	"os"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/stretchr/testify/assert"
)

const (
	traceMsg    = "this is a trace log"
	debugMsg    = "this is a debug log"
	infoMsg     = "this is a info log"
	warnMsg     = "this is a warn log"
	noticeMsg   = "this is a notice log"
	errorMsg    = "this is a error log"
	fatalMsg    = "this is a fatal log"
	logFileName = "hertz.log"
)

// TestLogger test logger work with hertz
func TestLogger(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := NewLogger()

	hlog.SetLogger(logger)
	hlog.SetOutput(buf)
	hlog.SetLevel(hlog.LevelError)

	hlog.Info(infoMsg)
	assert.Equal(t, "", buf.String())

	hlog.Error(errorMsg)
	// test SetLevel
	assert.Contains(t, buf.String(), errorMsg)

	buf.Reset()
	f, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		t.Error(err)
	}

	defer os.Remove(logFileName)

	hlog.SetOutput(f)

	hlog.Info(infoMsg)
	hlog.Error(errorMsg)
	_ = f.Sync()

	readF, err := os.OpenFile(logFileName, os.O_RDONLY, 0o400)
	if err != nil {
		t.Error(err)
	}
	line, _ := bufio.NewReader(readF).ReadString('\n')

	// test SetOutput
	assert.Contains(t, line, errorMsg)
}

func TestWithLevel(t *testing.T) {
	buf := new(bytes.Buffer)
	lvl := &slog.LevelVar{}
	lvl.Set(slog.LevelError)
	logger := NewLogger(WithLevel(lvl))

	hlog.SetLogger(logger)
	hlog.SetOutput(buf)

	hlog.Notice(infoMsg)
	assert.Equal(t, "", buf.String())

	hlog.Error(errorMsg)
	assert.Contains(t, buf.String(), errorMsg)

	buf.Reset()
	hlog.SetLevel(hlog.LevelDebug)
	hlog.Debug(debugMsg)

	assert.Contains(t, buf.String(), debugMsg)
}

func TestWithHandlerOptions(t *testing.T) {
	buf := new(bytes.Buffer)
	lvl := &slog.LevelVar{}
	lvl.Set(slog.LevelError)
	logger := NewLogger(WithHandlerOptions(&slog.HandlerOptions{Level: lvl}))

	hlog.SetLogger(logger)
	hlog.SetOutput(buf)

	hlog.Warn(warnMsg)
	assert.Equal(t, "", buf.String())

	hlog.SetLevel(hlog.LevelInfo)

	hlog.Debug(debugMsg)
	assert.Equal(t, "", buf.String())

	hlog.Info(infoMsg)
	assert.Contains(t, buf.String(), infoMsg)
}
