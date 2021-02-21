// Copyright 2021 glepnir. All rights reserved.

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package logger

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger *zerolog.Logger
}

var (
	loggerCtxKey struct{}
	traceIdKey   struct{}
)

func NewTraceIdContext(ctx context.Context, traceID string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, traceIdKey, traceID)
}

func NewContextLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerCtxKey, logger)
}

func FromContext(ctx context.Context) *Logger {
	logger, ok := ctx.Value(loggerCtxKey).(*Logger)
	if !ok {
		return nil
	}
	return logger
}

func NewLogger(isDebug bool, logDir string) (*Logger, error) {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "t"
	zerolog.MessageFieldName = "m"
	zerolog.CallerFieldName = "c"
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	return &Logger{logger: &logger}, nil
}

func (l *Logger) Debug() *zerolog.Event {
	return l.logger.Debug()
}

func (l *Logger) Name(fname string) *zerolog.Event {
	return l.Debug().Str("Name", fname)
}
