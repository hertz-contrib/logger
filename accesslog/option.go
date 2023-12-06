/*
 * Copyright 2022 CloudWeGo Authors
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
 *
 * MIT License
 *
 * Copyright (c) 2019-present Fenny and Contributors
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.E SOFTWARE.
 *
 * This file may have been modified by CloudWeGo authors. All CloudWeGo
 * Modifications are Copyright 2022 CloudWeGo Authors.
 */

package accesslog

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type (
	logConditionFunc func(ctx context.Context, c *app.RequestContext) bool

	// options defines the config for middleware.
	options struct {
		// format defines the logging tags
		//
		// Optional. Default: [${time}] ${status} - ${latency} ${method} ${path}\n
		format string

		// timeFormat defines timestamp format  https://programming.guide/go/format-parse-string-time-date-example.html
		//
		// Optional. Default: 15:04:05
		timeFormat string

		// timeInterval is the delay before the timestamp is updated
		//
		// Optional. Default: 500 * time.Millisecond
		timeInterval time.Duration

		// logFunc custom define log function
		//
		// Optional. Default: hlog.CtxInfof
		logFunc func(ctx context.Context, format string, v ...interface{})

		// timeZoneLocation can be specified time zone
		//
		// Optional. Default: time.Local
		timeZoneLocation *time.Location
		enableLatency    bool
		logConditionFunc logConditionFunc
	}

	Option func(o *options)
)

var defaultTagFormat = "[${time}] ${status} - ${latency} ${method} ${path}"

func newOptions(opts ...Option) *options {
	cfg := &options{
		format:           defaultTagFormat,
		timeFormat:       "15:04:05",
		timeZoneLocation: time.Local,
		timeInterval:     500 * time.Millisecond,
		logFunc:          hlog.CtxInfof,
		logConditionFunc: func(ctx context.Context, c *app.RequestContext) bool {
			return true
		},
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

// WithFormat set log format
func WithFormat(s string) Option {
	return func(o *options) {
		o.format = s
	}
}

// WithTimeFormat set log time format
func WithTimeFormat(s string) Option {
	return func(o *options) {
		o.timeFormat = s
	}
}

// WithTimeInterval set timestamp refresh interval
func WithTimeInterval(t time.Duration) Option {
	return func(o *options) {
		o.timeInterval = t
	}
}

// WithAccessLogFunc set print log function
func WithAccessLogFunc(f func(ctx context.Context, format string, v ...interface{})) Option {
	return func(o *options) {
		o.logFunc = f
	}
}

// WithTimeZoneLocation set timestamp zone
func WithTimeZoneLocation(loc *time.Location) Option {
	return func(o *options) {
		o.timeZoneLocation = loc
	}
}

// WithLogConditionFunc set logConditionFunc
func WithLogConditionFunc(f logConditionFunc) Option {
	return func(o *options) {
		o.logConditionFunc = f
	}
}
