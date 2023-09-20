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
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var _ hlog.FullLogger = (*Logger)(nil)

func NewLogger(opts ...Option) *Logger {
	config := defaultConfig()

	// apply options
	for _, opt := range opts {
		opt.apply(config)
	}
	if config.withLevel {
		config.handlerOptions.Level = config.level
	}

	if config.withHandlerOptions {
		lvl := &slog.LevelVar{}
		lvl.Set(config.handlerOptions.Level.Level())
		config.handlerOptions.Level = lvl
		config.level = lvl
	}

	return &Logger{
		l:   slog.New(slog.NewJSONHandler(os.Stdout, config.handlerOptions)),
		cfg: config,
	}
}

// Logger slog impl
type Logger struct {
	l   *slog.Logger
	cfg *config
}

func (l *Logger) Logger() *slog.Logger {
	return l.l
}

func (l *Logger) log(level hlog.Level, v ...any) {
	lvl := hLevelToSLevel(level)
	l.l.Log(context.TODO(), lvl, fmt.Sprint(v...))
}

func (l *Logger) logf(level hlog.Level, format string, kvs ...any) {
	lvl := hLevelToSLevel(level)
	l.l.Log(context.TODO(), lvl, fmt.Sprintf(format, kvs...))
}

func (l *Logger) ctxLogf(level hlog.Level, ctx context.Context, format string, v ...any) {
	lvl := hLevelToSLevel(level)
	l.l.Log(ctx, lvl, fmt.Sprintf(format, v...))
}

func (l *Logger) Trace(v ...any) {
	l.log(hlog.LevelTrace, v...)
}

func (l *Logger) Debug(v ...any) {
	l.log(hlog.LevelDebug, v...)
}

func (l *Logger) Info(v ...any) {
	l.log(hlog.LevelInfo, v...)
}

func (l *Logger) Notice(v ...any) {
	l.log(hlog.LevelNotice, v...)
}

func (l *Logger) Warn(v ...any) {
	l.log(hlog.LevelWarn, v...)
}

func (l *Logger) Error(v ...any) {
	l.log(hlog.LevelError, v...)
}

func (l *Logger) Fatal(v ...any) {
	l.log(hlog.LevelFatal, v...)
}

func (l *Logger) Tracef(format string, v ...any) {
	l.logf(hlog.LevelTrace, format, v...)
}

func (l *Logger) Debugf(format string, v ...any) {
	l.logf(hlog.LevelDebug, format, v...)
}

func (l *Logger) Infof(format string, v ...any) {
	l.logf(hlog.LevelInfo, format, v...)
}

func (l *Logger) Noticef(format string, v ...any) {
	l.logf(hlog.LevelNotice, format, v...)
}

func (l *Logger) Warnf(format string, v ...any) {
	l.logf(hlog.LevelWarn, format, v...)
}

func (l *Logger) Errorf(format string, v ...any) {
	l.logf(hlog.LevelError, format, v...)
}

func (l *Logger) Fatalf(format string, v ...any) {
	l.logf(hlog.LevelFatal, format, v...)
}

func (l *Logger) CtxTracef(ctx context.Context, format string, v ...any) {
	l.ctxLogf(hlog.LevelDebug, ctx, format, v...)
}

func (l *Logger) CtxDebugf(ctx context.Context, format string, v ...any) {
	l.ctxLogf(hlog.LevelDebug, ctx, format, v...)
}

func (l *Logger) CtxInfof(ctx context.Context, format string, v ...any) {
	l.ctxLogf(hlog.LevelInfo, ctx, format, v...)
}

func (l *Logger) CtxNoticef(ctx context.Context, format string, v ...any) {
	l.ctxLogf(hlog.LevelNotice, ctx, format, v...)
}

func (l *Logger) CtxWarnf(ctx context.Context, format string, v ...any) {
	l.ctxLogf(hlog.LevelWarn, ctx, format, v...)
}

func (l *Logger) CtxErrorf(ctx context.Context, format string, v ...any) {
	l.ctxLogf(hlog.LevelError, ctx, format, v...)
}

func (l *Logger) CtxFatalf(ctx context.Context, format string, v ...any) {
	l.ctxLogf(hlog.LevelFatal, ctx, format, v...)
}

func (l *Logger) SetLevel(level hlog.Level) {
	lvl := hLevelToSLevel(level)
	l.cfg.level.Set(lvl)
}

func (l *Logger) SetOutput(writer io.Writer) {
	l.cfg.output = writer
	l.l = slog.New(slog.NewJSONHandler(writer, l.cfg.handlerOptions))
}
