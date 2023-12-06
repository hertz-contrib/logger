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
 * The MIT License (MIT)
 *
 * Copyright (c) 2015-present Aliaksandr Valialkin, VertaMedia, Kirill Danshin, Erik Dubbelboer, FastHTTP Authors
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
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
	"strings"
	"testing"

	"github.com/cloudwego/hertz/pkg/protocol"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/bytebufferpool"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/cloudwego/hertz/pkg/route"
)

func TestLogger(t *testing.T) {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	hlog.SetOutput(buf)
	engine := route.NewEngine(config.NewOptions([]config.Option{}))
	engine.Use(New(WithFormat("${route}")))
	engine.GET("/", func(ctx context.Context, c *app.RequestContext) {})
	request := ut.PerformRequest(engine, "GET", "/", nil)
	w := request.Result()
	assert.DeepEqual(t, w.StatusCode(), 200)
	assert.DeepEqual(t, "/\n", buf.String()[len(buf.String())-2:])
}

func TestLoggerAll(t *testing.T) {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	hlog.SetOutput(buf)
	engine := route.NewEngine(config.NewOptions([]config.Option{}))
	engine.Use(New(WithFormat("${pid}${reqHeaders}${resHeaders}${referer}${protocol}${ip}${ips}" +
		"${host}${url}${ua}${body}${route}")))
	request := ut.PerformRequest(engine, "GET", "/?foo=bar", nil)
	w := request.Result()
	assert.DeepEqual(t, 404, w.StatusCode())
	assert.True(t, strings.Contains(buf.String(), fmt.Sprintf("%vContent-Type=text/plain; charset=utf-8http/?foo=bar/\n",
		os.Getpid())))
}

func TestQueryParams(t *testing.T) {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	hlog.SetOutput(buf)
	engine := route.NewEngine(config.NewOptions([]config.Option{}))
	engine.Use(New(WithFormat("${queryParams}")))
	request := ut.PerformRequest(engine, "GET", "/?foo=bar&baz=moz", nil)
	result := request.Result()
	assert.DeepEqual(t, 404, result.StatusCode())
	assert.True(t, strings.Contains(buf.String(), "foo=bar&baz=moz\n"))
}

func TestRespBody(t *testing.T) {
	const getBody = "Sample response body"
	const postBody = "Post in test"
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	hlog.SetOutput(buf)
	engine := route.NewEngine(config.NewOptions([]config.Option{}))
	engine.Use(New(WithFormat("${resBody}")))
	engine.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.String(200, getBody)
	})
	engine.POST("/test", func(ctx context.Context, c *app.RequestContext) {
		c.String(200, postBody)
	})
	request := ut.PerformRequest(engine, "GET", "/", nil)
	w := request.Result()
	assert.DeepEqual(t, 200, w.StatusCode())
	assert.DeepEqual(t, getBody+"\n", buf.String()[len(buf.String())-len(getBody)-1:])
	buf.Reset()
	request = ut.PerformRequest(engine, "POST", "/test", nil)
	w = request.Result()
	assert.DeepEqual(t, 200, w.StatusCode())
	assert.DeepEqual(t, postBody+"\n", buf.String()[len(buf.String())-len(postBody)-1:])
}

// go test -v -run=^$ -bench=Benchmark_Logger -benchmem -count=4
func Benchmark_AccessLog(b *testing.B) {
	hlog.SetOutput(io.Discard)

	benchSetup := func(b *testing.B, engine *route.Engine) {
		b.Helper()

		ctx := engine.NewContext()
		req := protocol.NewRequest("GET", "/", nil)

		b.ReportAllocs()
		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			req.CopyTo(&ctx.Request)
			engine.ServeHTTP(context.Background(), ctx)
			ctx.Reset()
		}
	}

	b.Run("Base", func(bb *testing.B) {
		engine := route.NewEngine(config.NewOptions([]config.Option{}))
		engine.Use(New(WithFormat("${bytesReceived} ${bytesSent} ${status}")))
		engine.GET("/", func(c context.Context, ctx *app.RequestContext) {
			ctx.Response.Header.Set("test", "test")

			ctx.String(200, "hello world")
		})
		benchSetup(bb, engine)
	})

	b.Run("DefaultFormat", func(bb *testing.B) {
		engine := route.NewEngine(config.NewOptions([]config.Option{}))
		engine.Use(New())
		engine.GET("/", func(c context.Context, ctx *app.RequestContext) {
			ctx.String(200, "hello world")
		})
	})
}
