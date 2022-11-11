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
	"io"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/rs/zerolog"
)

type (
	Options struct {
		context zerolog.Context
		level   zerolog.Level
	}

	Opt func(opts *Options)
)

func newOptions(log zerolog.Logger, options []Opt) *Options {
	opts := &Options{
		context: log.With(),
		level:   log.GetLevel(),
	}

	for _, set := range options {
		set(opts)
	}

	return opts
}

// WithOutput allows to specify the output of the logger. By default, it is set to os.Stdout.
func WithOutput(out io.Writer) Opt {
	return func(opts *Options) {
		opts.context = opts.context.Logger().Output(out).With()
	}
}

// WithLevel allows to specify the level of the logger. By default, it is set to WarnLevel.
func WithLevel(level hlog.Level) Opt {
	lvl := matchHlogLevel(level)
	return func(opts *Options) {
		opts.context = opts.context.Logger().Level(lvl).With()
		opts.level = lvl
	}
}

// WithField adds a field to the logger's context
func WithField(name string, value interface{}) Opt {
	return func(opts *Options) {
		opts.context = opts.context.Interface(name, value)
	}
}

// WithFields adds fields to the logger's context
func WithFields(fields map[string]interface{}) Opt {
	return func(opts *Options) {
		opts.context = opts.context.Fields(fields)
	}
}

// WithTimestamp adds a timestamp field to the logger's context
func WithTimestamp() Opt {
	return func(opts *Options) {
		opts.context = opts.context.Timestamp()
	}
}

// WithFormattedTimestamp adds a formatted timestamp field to the logger's context
func WithFormattedTimestamp(format string) Opt {
	zerolog.TimeFieldFormat = format
	return func(opts *Options) {
		opts.context = opts.context.Timestamp()
	}
}

// WithCaller adds a caller field to the logger's context
func WithCaller() Opt {
	return func(opts *Options) {
		opts.context = opts.context.Caller()
	}
}

// WithHook adds a hook to the logger's context
func WithHook(hook zerolog.Hook) Opt {
	return func(opts *Options) {
		opts.context = opts.context.Logger().Hook(hook).With()
	}
}

// WithHookFunc adds hook function to the logger's context
func WithHookFunc(hook zerolog.HookFunc) Opt {
	return func(opts *Options) {
		opts.context = opts.context.Logger().Hook(hook).With()
	}
}
