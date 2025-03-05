package log

import (
	"context"
	"go.uber.org/zap"
)

const (
	loggerKey = "zap-ctx-logger"
)

var (
	logger *zap.Logger
)

func Init() {
	logger, _ = zap.NewDevelopment()
}

func NewCtx(ctx context.Context, fields ...zap.Field) context.Context {
	return context.WithValue(ctx, loggerKey, WithCtx(ctx).With(fields...))
}

func WithCtx(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return logger
	}
	if ctxLogger, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return ctxLogger
	}
	return logger
}
