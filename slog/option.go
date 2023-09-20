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
	"os"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type Option interface {
	apply(cfg *config)
}

type option func(cfg *config)

func (fn option) apply(cfg *config) {
	fn(cfg)
}

type config struct {
	level              *slog.LevelVar
	withLevel          bool
	handlerOptions     *slog.HandlerOptions
	withHandlerOptions bool
	output             io.Writer
}

func defaultConfig() *config {
	lvl := &slog.LevelVar{}
	lvl.Set(hLevelToSLevel(hlog.LevelInfo))

	handlerOptions := &slog.HandlerOptions{
		Level: lvl,
	}
	return &config{
		level:              lvl,
		withLevel:          false,
		handlerOptions:     handlerOptions,
		withHandlerOptions: false,
		output:             os.Stdout,
	}
}

func WithLevel(lvl *slog.LevelVar) Option {
	return option(func(cfg *config) {
		cfg.level = lvl
		cfg.withLevel = true
	})
}

func WithHandlerOptions(opts *slog.HandlerOptions) Option {
	return option(func(cfg *config) {
		cfg.handlerOptions = opts
		cfg.withHandlerOptions = true
	})
}

func WithOutput(writer io.Writer) Option {
	return option(func(cfg *config) {
		cfg.output = writer
	})
}
