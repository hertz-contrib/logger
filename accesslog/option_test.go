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
	"fmt"
	"testing"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestDefaultOption(t *testing.T) {
	opts := newOptions()
	assert.DeepEqual(t, defaultTagFormat, opts.format)
	assert.DeepEqual(t, "15:04:05", opts.timeFormat)
	assert.DeepEqual(t, time.Local, opts.timeZoneLocation)
	assert.DeepEqual(t, 500*time.Millisecond, opts.timeInterval)
	assert.DeepEqual(t, fmt.Sprintf("%p", hlog.CtxInfof), fmt.Sprintf("%p", opts.logFunc))
}

func TestOption(t *testing.T) {
	opts := newOptions(
		WithFormat("[${time}] ${status} - ${latency} ${method} ${path} ${referer}"),
		WithTimeFormat(time.UnixDate),
		WithTimeInterval(100*time.Millisecond),
		WithAccessLogFunc(hlog.CtxDebugf),
	)
	assert.DeepEqual(t, "[${time}] ${status} - ${latency} ${method} ${path} ${referer}", opts.format)
	assert.DeepEqual(t, time.UnixDate, opts.timeFormat)
	assert.DeepEqual(t, 100*time.Millisecond, opts.timeInterval)
	assert.DeepEqual(t, fmt.Sprintf("%p", hlog.CtxDebugf), fmt.Sprintf("%p", opts.logFunc))
	assert.DeepEqual(t, false, opts.enableLatency)
}
