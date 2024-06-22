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
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InArray check if a string in a slice
func InArray(key ExtraKey, arr []ExtraKey) bool {
	for _, k := range arr {
		if k == key {
			return true
		}
	}
	return false
}

var hLevelToZapLevelMap = map[hlog.Level]zapcore.Level{
	hlog.LevelTrace:  zapcore.DebugLevel,
	hlog.LevelDebug:  zapcore.DebugLevel,
	hlog.LevelInfo:   zapcore.InfoLevel,
	hlog.LevelNotice: zapcore.WarnLevel,
	hlog.LevelWarn:   zapcore.WarnLevel,
	hlog.LevelError:  zapcore.ErrorLevel,
	hlog.LevelFatal:  zapcore.FatalLevel,
}

func hLevelToZapLevel(level hlog.Level) zapcore.Level {
	if zapLevel, ok := hLevelToZapLevelMap[level]; ok {
		return zapLevel
	}

	return zap.WarnLevel
}
