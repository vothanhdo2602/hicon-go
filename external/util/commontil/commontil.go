package commontil

import (
	"context"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"go.uber.org/zap"
)

func Recover(ctx context.Context) {
	if r := recover(); r != nil {
		log.WithCtx(ctx).Error("Recover from panic", zap.Any("error", r))
		return
	}
}
