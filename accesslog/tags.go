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
	"strings"
	"sync/atomic"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

const (
	TagPid               = "pid"
	TagTime              = "time"
	TagReferer           = "referer"
	TagProtocol          = "protocol"
	TagPort              = "port"
	TagIP                = "ip"
	TagIPs               = "ips"
	TagHost              = "host"
	TagClientIP          = "clientIP"
	TagMethod            = "method"
	TagPath              = "path"
	TagURL               = "url"
	TagUA                = "ua"
	TagLatency           = "latency"
	TagStatus            = "status"
	TagResBody           = "resBody"
	TagReqHeaders        = "reqHeaders"
	TagResHeaders        = "resHeaders"
	TagQueryStringParams = "queryParams"
	TagBody              = "body"
	TagBytesSent         = "bytesSent"
	TagBytesReceived     = "bytesReceived"
	TagRoute             = "route"
)

type LogFunc func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error)

// Data is a struct to define some variables to use in custom logger function.
type Data struct {
	Pid       string
	Start     time.Time
	Stop      time.Time
	Timestamp atomic.Value
}

var Tags = map[string]LogFunc{
	TagReferer: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		return output.WriteString(c.Request.Header.Get("Referer"))
	},
	TagProtocol: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		return output.WriteString(string(c.Request.URI().Scheme()))
	},
	TagPort: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		host := string(c.Request.URI().Host())
		split := strings.Split(host, ":")
		return output.WriteString(split[1])
	},
	TagIP: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		host := string(c.Request.URI().Host())
		split := strings.Split(host, ":")
		return output.WriteString(split[0])
	},
	TagIPs: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		return output.WriteString(c.Request.Header.Get("X-Forwarded-For"))
	},
	TagResBody: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		return output.WriteString(string(c.Response.Body()))
	},
	TagHost: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		return output.WriteString(string(c.Request.URI().Host()))
	},
	TagClientIP: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		return output.WriteString(c.ClientIP())
	},
	TagPath: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		return output.WriteString(string(c.Request.Path()))
	},
	TagURL: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		return output.WriteString(string(c.Request.Header.RequestURI()))
	},
	TagUA: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		return output.WriteString(c.Request.Header.Get("User-Agent"))
	},
	TagBody: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		return output.Write(c.Request.Body())
	},
	TagBytesSent: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		if c.Response.Header.ContentLength() < 0 {
			return appendInt(output, 0)
		}
		return appendInt(output, len(c.Response.Body()))
	},
	TagBytesReceived: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		return appendInt(output, len(c.Request.Body()))
	},
	TagRoute: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		return output.WriteString(string(c.Path()))
	},
	TagStatus: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		return appendInt(output, c.Response.StatusCode())
	},
	TagReqHeaders: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		reqHeaders := make([]string, 0)
		c.Request.Header.VisitAll(func(k, v []byte) {
			reqHeaders = append(reqHeaders, string(k)+"="+string(v))
		})
		return output.Write([]byte(strings.Join(reqHeaders, "&")))
	},
	TagResHeaders: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		resHeaders := make([]string, 0)
		c.Response.Header.VisitAll(func(k, v []byte) {
			resHeaders = append(resHeaders, string(k)+"="+string(v))
		})
		return output.Write([]byte(strings.Join(resHeaders, "&")))
	},
	TagQueryStringParams: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		return output.WriteString(c.Request.URI().QueryArgs().String())
	},
	TagMethod: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		return output.WriteString(string(c.Method()))
	},
	TagLatency: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		latency := data.Stop.Sub(data.Start)
		return output.WriteString(fmt.Sprintf("%13v", latency))
	},
	TagPid: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		return output.WriteString(data.Pid)
	},
	TagTime: func(output Buffer, c *app.RequestContext, data *Data, extraParam string) (int, error) {
		return output.WriteString(data.Timestamp.Load().(string))
	},
}
