// Copyright 2022 CloudWeGo Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slog

import (
	"io"
	"log/slog"

	cwslog "github.com/cloudwego-contrib/cwgo-pkg/log/logging/slog"
)

type Option = cwslog.Option

type config struct {
	options []cwslog.Option
}

func defaultConfig() *config {

	return &config{
		options: []cwslog.Option{},
	}
}

func WithLevel(lvl *slog.LevelVar) Option {
	return cwslog.WithLevel(lvl)
}

func WithHandlerOptions(opts *slog.HandlerOptions) Option {
	return cwslog.WithHandlerOptions(opts)
}

func WithOutput(writer io.Writer) Option {
	return cwslog.WithOutput(writer)
}
