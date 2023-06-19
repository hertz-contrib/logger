/*
 * Copyright 2022 CloudWeGo Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package zerolog

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestWithOutput(t *testing.T) {
	b := &bytes.Buffer{}
	l := New(WithOutput(b))

	l.Info("foobar")

	assert.Equal(
		t,
		`{"level":"info","message":"foobar"}
`,
		b.String(),
	)
}

func TestWithCaller(t *testing.T) {
	b := &bytes.Buffer{}
	l := New(WithCaller())
	l.SetOutput(b)

	l.Info("foobar")

	type Log struct {
		Level   string `json:"level"`
		Caller  string `json:"caller"`
		Message string `json:"message"`
	}

	log := &Log{}

	err := json.Unmarshal(b.Bytes(), log)

	assert.NoError(t, err)

	segments := strings.Split(log.Caller, ":")
	filePath := filepath.Base(segments[0])

	assert.Equal(t, filePath, "logger.go")
}

func TestWithCallerSkipFrameCount(t *testing.T) {
	b := &bytes.Buffer{}
	l := New(WithCallerSkipFrameCount(5))
	l.SetOutput(b)
	hlog.SetLogger(l)
	hlog.Info("foobar")

	type Log struct {
		Level   string `json:"level"`
		Caller  string `json:"caller"`
		Message string `json:"message"`
	}

	log := &Log{}

	err := json.Unmarshal(b.Bytes(), log)

	assert.NoError(t, err)

	segments := strings.Split(log.Caller, ":")
	filePath := filepath.Base(segments[0])

	assert.Equal(t, filePath, "options_test.go")
}

func TestWithField(t *testing.T) {
	b := &bytes.Buffer{}
	l := New(WithField("service", "logging"))
	l.SetOutput(b)

	l.Info("foobar")

	type Log struct {
		Level   string `json:"level"`
		Service string `json:"service"`
		Message string `json:"message"`
	}

	log := &Log{}

	err := json.Unmarshal(b.Bytes(), log)

	assert.NoError(t, err)
	assert.Equal(t, log.Service, "logging")
}

func TestWithFields(t *testing.T) {
	b := &bytes.Buffer{}
	l := New(WithFields(map[string]interface{}{
		"host": "localhost",
		"port": 8080,
	}))
	l.SetOutput(b)

	l.Info("foobar")

	type Log struct {
		Level   string `json:"level"`
		Host    string `json:"host"`
		Port    int    `json:"port"`
		Message string `json:"message"`
	}

	log := &Log{}

	err := json.Unmarshal(b.Bytes(), log)

	assert.NoError(t, err)
	assert.Equal(t, log.Host, "localhost")
	assert.Equal(t, log.Port, 8080)
}

type (
	Hook struct {
		logs []HookLog
	}

	HookLog struct {
		level   zerolog.Level
		message string
	}
)

func (h *Hook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	h.logs = append(h.logs, HookLog{
		level:   level,
		message: message,
	})
}

func TestWithHook(t *testing.T) {
	b := &bytes.Buffer{}
	h := &Hook{}
	l := New(WithHook(h))
	l.SetOutput(b)

	l.Info("Foo")
	l.Warn("Bar")

	assert.Len(t, h.logs, 2)
	assert.Equal(t, h.logs[0].level, zerolog.InfoLevel)
	assert.Equal(t, h.logs[0].message, "Foo")
	assert.Equal(t, h.logs[1].level, zerolog.WarnLevel)
	assert.Equal(t, h.logs[1].message, "Bar")
}

func TestWithHookFunc(t *testing.T) {
	b := &bytes.Buffer{}
	logs := make([]HookLog, 0, 2)
	l := New(WithHookFunc(func(e *zerolog.Event, level zerolog.Level, message string) {
		logs = append(logs, HookLog{
			level:   level,
			message: message,
		})
	}))
	l.SetOutput(b)

	l.Info("Foo")
	l.Warn("Bar")

	assert.Len(t, logs, 2)
	assert.Equal(t, logs[0].level, zerolog.InfoLevel)
	assert.Equal(t, logs[0].message, "Foo")
	assert.Equal(t, logs[1].level, zerolog.WarnLevel)
	assert.Equal(t, logs[1].message, "Bar")
}

func TestWithLevel(t *testing.T) {
	b := &bytes.Buffer{}
	l := New(WithLevel(hlog.LevelInfo))
	l.SetOutput(b)

	l.Debug("Test")

	assert.Equal(t, b.String(), "")

	l.Info("foobar")

	assert.Equal(t, b.String(), `{"level":"info","message":"foobar"}
`)
}

type Log struct {
	Level   string    `json:"level"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func TestWithTimestamp(t *testing.T) {
	b := &bytes.Buffer{}
	l := New(WithTimestamp())
	l.SetOutput(b)

	l.Info("foobar")

	log := &Log{}

	err := json.Unmarshal(b.Bytes(), log)

	assert.NoError(t, err)
	assert.NotEmpty(t, log.Time)
}

func TestWithFormattedTimestamp(t *testing.T) {
	b := &bytes.Buffer{}
	l := New(WithFormattedTimestamp(time.RFC3339Nano))
	l.SetOutput(b)

	l.Info("foobar")

	log := &Log{}
	err := json.Unmarshal(b.Bytes(), log)

	assert.NoError(t, err)
	assert.NotEmpty(t, log.Time)
}
