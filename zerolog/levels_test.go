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

package zerolog

import (
	"testing"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestMatchHlogLevel(t *testing.T) {
	assert.Equal(t, zerolog.TraceLevel, matchHlogLevel(hlog.LevelTrace))
	assert.Equal(t, zerolog.DebugLevel, matchHlogLevel(hlog.LevelDebug))
	assert.Equal(t, zerolog.InfoLevel, matchHlogLevel(hlog.LevelInfo))
	assert.Equal(t, zerolog.WarnLevel, matchHlogLevel(hlog.LevelWarn))
	assert.Equal(t, zerolog.ErrorLevel, matchHlogLevel(hlog.LevelError))
	assert.Equal(t, zerolog.FatalLevel, matchHlogLevel(hlog.LevelFatal))
}

func TestMatchZerologLevel(t *testing.T) {
	assert.Equal(t, hlog.LevelTrace, matchZerologLevel(zerolog.TraceLevel))
	assert.Equal(t, hlog.LevelDebug, matchZerologLevel(zerolog.DebugLevel))
	assert.Equal(t, hlog.LevelInfo, matchZerologLevel(zerolog.InfoLevel))
	assert.Equal(t, hlog.LevelWarn, matchZerologLevel(zerolog.WarnLevel))
	assert.Equal(t, hlog.LevelError, matchZerologLevel(zerolog.ErrorLevel))
	assert.Equal(t, hlog.LevelFatal, matchZerologLevel(zerolog.FatalLevel))
}
