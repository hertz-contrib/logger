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
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/rs/zerolog"
)

var _ hlog.FullLogger = (*Logger)(nil)

// Logger is a wrapper around `zerolog.Logger` that provides an implementation of `hlog.FullLogger` interface
type Logger struct {
	log     zerolog.Logger
	out     io.Writer
	level   zerolog.Level
	options []Opt
}

// ConsoleWriter parses the JSON input and writes it in an
// (optionally) colorized, human-friendly format to Out.
type ConsoleWriter = zerolog.ConsoleWriter

// MultiLevelWriter may be used to send the log message to multiple outputs.
func MultiLevelWriter(writers ...io.Writer) zerolog.LevelWriter {
	return zerolog.MultiLevelWriter(writers...)
}

// New returns a new Logger instance
func New(options ...Opt) *Logger {
	return newLogger(zerolog.New(os.Stdout), options)
}

// From returns a new Logger instance using an existing logger
func From(log zerolog.Logger, options ...Opt) *Logger {
	return newLogger(log, options)
}

// GetLogger returns the default logger instance
func GetLogger() (Logger, error) {
	defaultLogger := hlog.DefaultLogger()

	if logger, ok := defaultLogger.(*Logger); ok {
		return *logger, nil
	}

	return Logger{}, errors.New("hlog.DefaultLogger is not a zerolog logger")
}

// SetLevel setting logging level for logger
func (l *Logger) SetLevel(level hlog.Level) {
	lvl := matchHlogLevel(level)
	l.level = lvl
	l.log = l.log.Level(lvl)
}

// SetOutput setting output for logger
func (l *Logger) SetOutput(writer io.Writer) {
	l.out = writer
	l.log = l.log.Output(writer)
}

// WithContext returns context with logger attached
func (l *Logger) WithContext(ctx context.Context) context.Context {
	return l.log.WithContext(ctx)
}

// WithField appends a field to the logger
func (l *Logger) WithField(key string, value interface{}) Logger {
	l.log = l.log.With().Interface(key, value).Logger()
	return *l
}

// Unwrap returns the underlying zerolog logger
func (l *Logger) Unwrap() zerolog.Logger {
	return l.log
}

// Log log using zerolog logger with specified level
func (l *Logger) Log(level hlog.Level, kvs ...interface{}) {
	switch level {
	case hlog.LevelTrace, hlog.LevelDebug:
		l.log.Debug().Msg(fmt.Sprint(kvs...))
	case hlog.LevelInfo:
		l.log.Info().Msg(fmt.Sprint(kvs...))
	case hlog.LevelNotice, hlog.LevelWarn:
		l.log.Warn().Msg(fmt.Sprint(kvs...))
	case hlog.LevelError:
		l.log.Error().Msg(fmt.Sprint(kvs...))
	case hlog.LevelFatal:
		l.log.Fatal().Msg(fmt.Sprint(kvs...))
	default:
		l.log.Warn().Msg(fmt.Sprint(kvs...))
	}
}

// Logf log using zerolog logger with specified level and formatting
func (l *Logger) Logf(level hlog.Level, format string, kvs ...interface{}) {
	switch level {
	case hlog.LevelTrace, hlog.LevelDebug:
		l.log.Debug().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelInfo:
		l.log.Info().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelNotice, hlog.LevelWarn:
		l.log.Warn().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelError:
		l.log.Error().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelFatal:
		l.log.Fatal().Msg(fmt.Sprintf(format, kvs...))
	default:
		l.log.Warn().Msg(fmt.Sprintf(format, kvs...))
	}
}

// CtxLogf log with logger associated with context.
// If no logger is associated, DefaultContextLogger is used, unless DefaultContextLogger is nil, in which case a disabled logger is used.
func (l *Logger) CtxLogf(level hlog.Level, ctx context.Context, format string, kvs ...interface{}) {
	logger := zerolog.Ctx(ctx)
	switch level {
	case hlog.LevelTrace, hlog.LevelDebug:
		logger.Debug().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelInfo:
		logger.Info().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelNotice, hlog.LevelWarn:
		logger.Warn().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelError:
		logger.Error().Msg(fmt.Sprintf(format, kvs...))
	case hlog.LevelFatal:
		logger.Fatal().Msg(fmt.Sprintf(format, kvs...))
	default:
		logger.Warn().Msg(fmt.Sprintf(format, kvs...))
	}
}

// Trace logs a message at trace level.
func (l *Logger) Trace(v ...interface{}) {
	l.Log(hlog.LevelTrace, v...)
}

// Debug logs a message at debug level.
func (l *Logger) Debug(v ...interface{}) {
	l.Log(hlog.LevelDebug, v...)
}

// Info logs a message at info level.
func (l *Logger) Info(v ...interface{}) {
	l.Log(hlog.LevelInfo, v...)
}

// Notice logs a message at notice level.
func (l *Logger) Notice(v ...interface{}) {
	l.Log(hlog.LevelNotice, v...)
}

// Warn logs a message at warn level.
func (l *Logger) Warn(v ...interface{}) {
	l.Log(hlog.LevelWarn, v...)
}

// Error logs a message at error level.
func (l *Logger) Error(v ...interface{}) {
	l.Log(hlog.LevelError, v...)
}

// Fatal logs a message at fatal level.
func (l *Logger) Fatal(v ...interface{}) {
	l.Log(hlog.LevelFatal, v...)
}

// Tracef logs a formatted message at trace level.
func (l *Logger) Tracef(format string, v ...interface{}) {
	l.Logf(hlog.LevelTrace, format, v...)
}

// Debugf logs a formatted message at debug level.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Logf(hlog.LevelDebug, format, v...)
}

// Infof logs a formatted message at info level.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Logf(hlog.LevelInfo, format, v...)
}

// Noticef logs a formatted message at notice level.
func (l *Logger) Noticef(format string, v ...interface{}) {
	l.Logf(hlog.LevelWarn, format, v...)
}

// Warnf logs a formatted message at warn level.
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Logf(hlog.LevelWarn, format, v...)
}

// Errorf logs a formatted message at error level.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Logf(hlog.LevelError, format, v...)
}

// Fatalf logs a formatted message at fatal level.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Logf(hlog.LevelError, format, v...)
}

// CtxTracef logs a message at trace level with logger associated with context.
// If no logger is associated, DefaultContextLogger is used, unless DefaultContextLogger is nil, in which case a disabled logger is used.
func (l *Logger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelTrace, ctx, format, v...)
}

// CtxDebugf logs a message at debug level with logger associated with context.
// If no logger is associated, DefaultContextLogger is used, unless DefaultContextLogger is nil, in which case a disabled logger is used.
func (l *Logger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelDebug, ctx, format, v...)
}

// CtxInfof logs a message at info level with logger associated with context.
// If no logger is associated, DefaultContextLogger is used, unless DefaultContextLogger is nil, in which case a disabled logger is used.
func (l *Logger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelInfo, ctx, format, v...)
}

// CtxNoticef logs a message at notice level with logger associated with context.
// If no logger is associated, DefaultContextLogger is used, unless DefaultContextLogger is nil, in which case a disabled logger is used.
func (l *Logger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelNotice, ctx, format, v...)
}

// CtxWarnf logs a message at warn level with logger associated with context.
// If no logger is associated, DefaultContextLogger is used, unless DefaultContextLogger is nil, in which case a disabled logger is used.
func (l *Logger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelWarn, ctx, format, v...)
}

// CtxErrorf logs a message at error level with logger associated with context.
// If no logger is associated, DefaultContextLogger is used, unless DefaultContextLogger is nil, in which case a disabled logger is used.
func (l *Logger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelError, ctx, format, v...)
}

// CtxFatalf logs a message at fatal level with logger associated with context.
// If no logger is associated, DefaultContextLogger is used, unless DefaultContextLogger is nil, in which case a disabled logger is used.
func (l *Logger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	l.CtxLogf(hlog.LevelFatal, ctx, format, v...)
}

func newLogger(log zerolog.Logger, options []Opt) *Logger {
	opts := newOptions(log, options)

	return &Logger{
		log:     opts.context.Logger(),
		out:     nil,
		level:   opts.level,
		options: options,
	}
}
