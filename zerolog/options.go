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
	cwzerolog "github.com/cloudwego-contrib/cwgo-pkg/log/logging/zerolog"
	"io"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/rs/zerolog"
)

type (
	Opt = cwzerolog.Opt
)

// WithOutput allows to specify the output of the logger. By default, it is set to os.Stdout.
func WithOutput(out io.Writer) Opt {
	return cwzerolog.WithOutput(out)
}

// WithLevel allows to specify the level of the logger. By default, it is set to WarnLevel.
func WithLevel(level hlog.Level) Opt {
	//lvl := matchHlogLevel(level)
	return cwzerolog.WithLevel(level)
}

// WithField adds a field to the logger's context
func WithField(name string, value interface{}) Opt {
	return cwzerolog.WithField(name, value)
}

// WithFields adds fields to the logger's context
func WithFields(fields map[string]interface{}) Opt {
	return cwzerolog.WithFields(fields)
}

// WithTimestamp adds a timestamp field to the logger's context
func WithTimestamp() Opt {
	return cwzerolog.WithTimestamp()
}

// WithFormattedTimestamp adds a formatted timestamp field to the logger's context
func WithFormattedTimestamp(format string) Opt {
	//zerolog.TimeFieldFormat = format
	return cwzerolog.WithFormattedTimestamp(format)
}

// WithCaller adds a caller field to the logger's context
func WithCaller() Opt {
	return cwzerolog.WithCaller()
}

// WithCallerSkipFrameCount adds a caller field to the logger's context
// The specified skipFrameCount int will override the global CallerSkipFrameCount for this context's respective logger.
// If set to -1 the global CallerSkipFrameCount will be used.
func WithCallerSkipFrameCount(skipFrameCount int) Opt {
	return cwzerolog.WithCallerSkipFrameCount(skipFrameCount)
}

// WithHook adds a hook to the logger's context
func WithHook(hook zerolog.Hook) Opt {
	return cwzerolog.WithHook(hook)
}

// WithHookFunc adds hook function to the logger's context
func WithHookFunc(hook zerolog.HookFunc) Opt {
	return cwzerolog.WithHookFunc(hook)
}
