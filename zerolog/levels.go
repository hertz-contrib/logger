/*
 * Copyright 2022 CloudWeGo Authors.
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
 */

package zerolog

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/rs/zerolog"
)

var (
	zerologLevels = map[hlog.Level]zerolog.Level{
		hlog.LevelTrace:  zerolog.TraceLevel,
		hlog.LevelDebug:  zerolog.DebugLevel,
		hlog.LevelInfo:   zerolog.InfoLevel,
		hlog.LevelWarn:   zerolog.WarnLevel,
		hlog.LevelNotice: zerolog.WarnLevel,
		hlog.LevelError:  zerolog.ErrorLevel,
		hlog.LevelFatal:  zerolog.FatalLevel,
	}

	hlogLevels = map[zerolog.Level]hlog.Level{
		zerolog.TraceLevel: hlog.LevelTrace,
		zerolog.DebugLevel: hlog.LevelDebug,
		zerolog.InfoLevel:  hlog.LevelInfo,
		zerolog.WarnLevel:  hlog.LevelWarn,
		zerolog.ErrorLevel: hlog.LevelError,
		zerolog.FatalLevel: hlog.LevelFatal,
	}
)

// matchHlogLevel map hlog.Level to zerolog.Level
func matchHlogLevel(level hlog.Level) zerolog.Level {
	zlvl, found := zerologLevels[level]

	if found {
		return zlvl
	}

	return zerolog.WarnLevel // Default level
}

// matchZerologLevel map zerolog.Level to hlog.Level
func matchZerologLevel(level zerolog.Level) hlog.Level {
	hlvl, found := hlogLevels[level]

	if found {
		return hlvl
	}

	return hlog.LevelWarn // Default level
}
