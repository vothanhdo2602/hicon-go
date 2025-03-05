package natstil

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/nats-io/nats.go"
	"github.com/vothanhdo2602/hicon/external/constant"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/external/util/mongotil"
	"go.uber.org/zap"
	"net/http"
)

func GetContext(msg *nats.Msg) context.Context {
	var (
		ctx       = context.Background()
		requestID = msg.Header.Get(constant.HeaderXRequestId)
	)

	if requestID == "" {
		requestID = mongotil.NewHexID()
	}

	return log.NewCtx(
		ctx,
		zap.String(constant.HeaderXRequestId, requestID),
	)
}

func Response(ctx context.Context, msg *nats.Msg) {
	if err := msg.RespondMsg(msg); err != nil {
		log.WithCtx(ctx).Error(err.Error())
	}
}

func R200(data interface{}) []byte {
	r := &BaseResponse[any]{
		Data:    data,
		Status:  http.StatusOK,
		Success: true,
	}
	respBytes, _ := json.Marshal(r)
	return respBytes
}

func R400(msg string) []byte {
	r := &BaseResponse[any]{
		Message: msg,
		Status:  http.StatusBadRequest,
		Success: false,
	}
	respBytes, _ := json.Marshal(r)
	return respBytes
}
