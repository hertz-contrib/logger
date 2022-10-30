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

package logrus_test

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	ctx := context.Background()

	logger := hertzlogrus.NewLogger(hertzlogrus.WithLogger(logrus.New()))

	logger.Logger().Info("log from origin logrus")

	hlog.SetLogger(logger)
	hlog.SetLevel(hlog.LevelError)
	hlog.SetLevel(hlog.LevelWarn)
	hlog.SetLevel(hlog.LevelInfo)
	hlog.SetLevel(hlog.LevelDebug)
	hlog.SetLevel(hlog.LevelTrace)

	hlog.Trace("trace")
	hlog.Debug("debug")
	hlog.Info("info")
	hlog.Notice("notice")
	hlog.Warn("warn")
	hlog.Error("error")

	hlog.Tracef("log level: %s", "trace")
	hlog.Debugf("log level: %s", "debug")
	hlog.Infof("log level: %s", "info")
	hlog.Noticef("log level: %s", "notice")
	hlog.Warnf("log level: %s", "warn")
	hlog.Errorf("log level: %s", "error")

	hlog.CtxTracef(ctx, "log level: %s", "trace")
	hlog.CtxDebugf(ctx, "log level: %s", "debug")
	hlog.CtxInfof(ctx, "log level: %s", "info")
	hlog.CtxNoticef(ctx, "log level: %s", "notice")
	hlog.CtxWarnf(ctx, "log level: %s", "warn")
	hlog.CtxErrorf(ctx, "log level: %s", "error")
}
