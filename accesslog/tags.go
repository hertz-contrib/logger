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
	"github.com/cloudwego-contrib/cwgo-pkg/log/accesslog"
)

const (
	TagPid               = accesslog.TagPid
	TagTime              = accesslog.TagTime
	TagReferer           = accesslog.TagReferer
	TagProtocol          = accesslog.TagProtocol
	TagPort              = accesslog.TagPort
	TagIP                = accesslog.TagIP
	TagIPs               = accesslog.TagIPs
	TagHost              = accesslog.TagHost
	TagClientIP          = accesslog.TagClientIP
	TagMethod            = accesslog.TagMethod
	TagPath              = accesslog.TagPath
	TagURL               = accesslog.TagURL
	TagUA                = accesslog.TagUA
	TagLatency           = accesslog.TagLatency
	TagStatus            = accesslog.TagStatus
	TagResBody           = accesslog.TagResBody
	TagReqHeaders        = accesslog.TagReqHeaders
	TagResHeaders        = accesslog.TagResHeaders
	TagQueryStringParams = accesslog.TagQueryStringParams
	TagBody              = accesslog.TagBody
	TagBytesSent         = accesslog.TagBytesSent
	TagBytesReceived     = accesslog.TagBytesReceived
	TagRoute             = accesslog.TagRoute
)

type LogFunc = accesslog.LogFunc

// Data is a struct to define some variables to use in custom logger function.
type Data = accesslog.Data

var Tags = accesslog.Tags
