package commontil

import (
	"context"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"go.uber.org/zap"
	"io"
	"net/http"
	"sync"
)

func InitStructFromGeneric[T any]() T {
	return make([]T, 1)[0]
}

func BackgroundRun(ctx context.Context, wg *sync.WaitGroup, fn func(ctx context.Context)) {
	if wg != nil {
		defer wg.Done()
	}

	defer Recover(ctx)

	fn(ctx)
}

func GetCurrentPublicIP(ctx context.Context) string {
	var (
		logger = log.WithCtx(ctx)
	)

	req, err := http.Get("https://icanhazip.com/")
	if err != nil {
		logger.Error(err.Error())
		return ""
	}

	defer func() {
		err := req.Body.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}()

	bytesResp, err := io.ReadAll(req.Body)
	if err != nil {
		logger.Error(err.Error())
		return ""
	}

	return string(bytesResp)
}

func Recover(ctx context.Context) {
	if r := recover(); r != nil {
		log.WithCtx(ctx).Error("Recover from panic", zap.Any("error", r))
		return
	}
}
