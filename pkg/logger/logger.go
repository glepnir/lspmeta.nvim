// Copyright 2021 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"io"
	"os"
	"path"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	logger *zerolog.Logger
}

func newRollingFile(logDir string) io.Writer {
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Error().Err(err).Str("path", logDir).Msg("can't create log directory")
		return nil
	}

	return &lumberjack.Logger{
		Filename: path.Join(logDir, "lspmeta.log"),
	}
}
func NewLogger(isDebug bool, logDir string) (*Logger, error) {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}

	writer := newRollingFile(logDir)

	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(writer).With().Timestamp().Logger()

	return &Logger{logger: &logger}, nil
}

func (l *Logger) Print(err error) {
	l.logger.Print(err)
}
