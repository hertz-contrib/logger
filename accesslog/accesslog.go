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
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/bytebufferpool"
	"github.com/hertz-contrib/logger/accesslog/internal/fasttemplate"
)

var defaultFormat = " %s | %3d | %7v | %-7s | %-s "

func New(opts ...Option) app.HandlerFunc {
	return new(context.Background(), opts...)
}

func NewWithContext(ctx context.Context, opts ...Option) app.HandlerFunc {
	return new(ctx, opts...)
}

func new(ctx context.Context, opts ...Option) app.HandlerFunc {
	cfg := newOptions(opts...)
	// Check if format contains latency
	cfg.enableLatency = strings.Contains(cfg.format, "${latency}")

	tmpl := fasttemplate.New(cfg.format, "${", "}")

	// Create correct time format
	var timestamp atomic.Value

	// Update date/time every 500 milliseconds in a separate go routine
	if strings.Contains(cfg.format, "${time}") {
		go func() {
			for {
				select {
				case <-time.After(cfg.timeInterval):
				case <-ctx.Done():
					return
				}
				timestamp.Store(time.Now().In(cfg.timeZoneLocation).Format(cfg.timeFormat))
			}
		}()
	}

	// Set PID once and add tag
	pid := strconv.Itoa(os.Getpid())
	Tags[TagPid] = func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return buf.WriteString(pid)
	}
	Tags[TagTime] = func(ctx context.Context, c *app.RequestContext, buf *bytebufferpool.ByteBuffer) (int, error) {
		return buf.WriteString(timestamp.Load().(string))
	}

	return func(ctx context.Context, c *app.RequestContext) {
		var start, stop time.Time

		// Set latency start time
		if cfg.enableLatency {
			start = time.Now()
			c.Set("start", start)
		}

		c.Next(ctx)

		if cfg.enableLatency {
			stop = time.Now()
			c.Set("stop", stop)
		}

		// Get new buffer
		buf := bytebufferpool.Get()
		defer bytebufferpool.Put(buf)

		if cfg.format == defaultTagFormat {
			// format log to buffer
			_, _ = buf.WriteString(fmt.Sprintf(defaultFormat,
				timestamp,
				c.Response.StatusCode(),
				stop.Sub(start),
				c.Method(),
				c.Path(),
			))

			cfg.logFunc(ctx, buf.String())
			return
		}

		_, err := tmpl.ExecuteFunc(buf, func(w io.Writer, tag string) (int, error) {
			if function, ok := Tags[tag]; ok {
				return function(ctx, c, buf)
			}
			return 0, nil
		})
		// Also write errors to the buffer
		if err != nil {
			_, _ = buf.WriteString(err.Error())
		}

		cfg.logFunc(ctx, buf.String())
	}
}

func appendInt(buf *bytebufferpool.ByteBuffer, v int) (int, error) {
	old := len(buf.B)
	buf.B = appendUint(buf.B, v)
	return len(buf.B) - old, nil
}

func appendUint(dst []byte, n int) []byte {
	if n < 0 {
		panic("BUG: int must be positive")
	}

	var b [20]byte
	buf := b[:]
	i := len(buf)
	var q int
	for n >= 10 {
		i--
		q = n / 10
		buf[i] = '0' + byte(n-q*10)
		n = q
	}
	i--
	buf[i] = '0' + byte(n)

	dst = append(dst, buf[i:]...)
	return dst
}
