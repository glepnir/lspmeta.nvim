// Copyright 2021 glepnir. All rights reserved.

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package logger

import (
	"context"
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

func NewTraceIdContext(ctx context.Context, traceID string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, "traceID", traceID)
}

func newRollingFile(logDir string) io.Writer {
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Error().Err(
			err,
		).Str(
			"path",
			logDir,
		).Msg(
			"can't create log directory",
		)
		return nil
	}
	return &lumberjack.Logger{
		Filename:   path.Join(logDir, "lspmeta.log"),
		MaxSize:    0,
		MaxAge:     0,
		MaxBackups: 0,
		LocalTime:  false,
		Compress:   false,
	}
}

func NewLogger(isDebug bool, logDir string) (*Logger, error) {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	writer := newRollingFile(logDir)
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "t"
	zerolog.MessageFieldName = "m"
	zerolog.CallerFieldName = "c"
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(writer).With().Timestamp().Caller().Logger()
	return &Logger{logger: &logger}, nil
}

func WithContext(ctx context.Context) *Logger {
	if ctx == nil {
		logger, err := NewLogger(true, "ss")
		if err != nil {
			panic("init logger error")
		}
		return logger
	}
	if ctxlogger, ok := ctx.Value("loggerkey").(*Logger); ok {
		return ctxlogger
	}
	return nil
}

func (l *Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}
