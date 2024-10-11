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
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/rs/zerolog"
	"io"
)

var _ hlog.FullLogger = (*Logger)(nil)

// Logger is a wrapper around `zerolog.Logger` that provides an implementation of `hlog.FullLogger` interface
type Logger = cwzerolog.Logger

// ConsoleWriter parses the JSON input and writes it in an
// (optionally) colorized, human-friendly format to Out.
type ConsoleWriter = cwzerolog.ConsoleWriter

// MultiLevelWriter may be used to send the log message to multiple outputs.
func MultiLevelWriter(writers ...io.Writer) zerolog.LevelWriter {
	return cwzerolog.MultiLevelWriter(writers...)
}

// New returns a new Logger instance
func New(options ...Opt) *Logger {
	return cwzerolog.New(options...)
}

// From returns a new Logger instance using an existing logger
func From(log zerolog.Logger, options ...Opt) *Logger {
	return cwzerolog.From(log, options...)
}

// GetLogger returns the default logger instance
func GetLogger() (Logger, error) {

	return cwzerolog.GetLogger()
}
