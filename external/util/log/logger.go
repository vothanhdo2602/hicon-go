package log

import (
	"context"
	"github.com/vothanhdo2602/hicon/external/util/mongotil"
	"go.uber.org/zap"
)

const (
	loggerKey  = "zap-ctx-logger"
	XRequestID = "X-Request-ID"
)

var (
	logger *zap.Logger
)

func Init() {
	logger = zap.Must(zap.NewDevelopment())
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

func GetContext(ctx context.Context, headers map[string]string) context.Context {
	requestID := headers[XRequestID]
	if requestID == "" {
		requestID = mongotil.NewHexID()
	}

	return NewCtx(
		ctx,
		zap.String(XRequestID, requestID),
	)
}
