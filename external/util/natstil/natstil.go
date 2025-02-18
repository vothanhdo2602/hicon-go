package natstil

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/nats-io/nats.go"
	"github.com/vothanhdo2602/hicon/external/constant"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/external/util/mongotil"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

const (
	stream = "hicon"
)

func GenerateReqrepSubject(channel string) string {
	return fmt.Sprintf("%s.reqrep.%s", stream, channel)
}

func GenerateJetstreamSubject(channel, action string) string {
	return fmt.Sprintf("%s.jetstream.%s.%s", stream, channel, action)
}

func GenerateQueueNameFromSubject(subject string) string {
	return strings.ReplaceAll(subject, ".", "_")
}

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

func R400(msg string) []byte {
	r := &IResponse[any]{
		Message: msg,
		Status:  http.StatusOK,
		Success: true,
	}
	respBytes, _ := json.Marshal(r)
	return respBytes
}
