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

package zap

import (
	cwzap "github.com/cloudwego-contrib/cwgo-pkg/log/logging/zap"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option = cwzap.Option

type ExtraKey = cwzap.ExtraKey

type CoreConfig = cwzap.CoreConfig

// WithCoreEnc zapcore encoder
func WithCoreEnc(enc zapcore.Encoder) Option {
	return cwzap.WithCoreEnc(enc)
}

// WithCoreWs zapcore write syncer
func WithCoreWs(ws zapcore.WriteSyncer) Option {
	return cwzap.WithCoreWs(ws)
}

// WithCoreLevel zapcore log level
func WithCoreLevel(lvl zap.AtomicLevel) Option {
	return cwzap.WithCoreLevel(lvl)
}

// WithCores zapcore
func WithCores(coreConfigs ...CoreConfig) Option {
	return cwzap.WithCores(coreConfigs...)
}

// WithZapOptions add origin zap option
func WithZapOptions(opts ...zap.Option) Option {
	return cwzap.WithZapOptions(opts...)
}

// WithExtraKeys allow you log extra values from context
func WithExtraKeys(keys []ExtraKey) Option {
	return cwzap.WithExtraKeys(keys)
}

// WithExtraKeyAsStr convert extraKey to a string type when retrieving value from context
// Not recommended for use, only for compatibility with certain situations
//
// For more information, refer to the documentation at
// `https://pkg.go.dev/context#WithValue`
func WithExtraKeyAsStr() Option {
	return cwzap.WithExtraKeyAsStr()
}
