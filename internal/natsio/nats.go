package natsio

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/internal/natsio/handler"
	"sync"
	"time"
)

var (
	nc    *nats.Conn
	jsCtx nats.JetStreamContext
	js    jetstream.JetStream
)

func Init(ctx context.Context, wg *sync.WaitGroup) (err error) {
	var (
		logger = log.WithCtx(ctx)
	)

	if wg != nil {
		defer wg.Done()
	}

	nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	jsCtx, err = nc.JetStream(nats.PublishAsyncMaxPending(10000))
	if err != nil {
		logger.Error(err.Error())
		return
	}

	js, err = jetstream.New(nc)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	err = RegisterReqrep(ctx)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info(fmt.Sprintf("⚡️[natsio]: connected to %s", nats.DefaultURL))

	return
}

func GetNC() *nats.Conn {
	return nc
}

func GetJs() jetstream.JetStream {
	return js
}

func RegisterReqrep(ctx context.Context) (err error) {
	if err = ReqrepQueueSubscribe(ctx, GetFindByPrimaryKeysSubject(), handler.FindByPrimaryKeys); err != nil {
		return
	}
	return
}

func ReqrepQueueSubscribe(ctx context.Context, subject string, cb func(msg *nats.Msg)) error {
	var (
		logger = log.WithCtx(ctx)
		queue  = GenerateQueueNameFromSubject(subject)
	)

	_, err := nc.QueueSubscribe(subject, queue, cb)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func JetstreamQueueSubscribe(ctx context.Context, stream, channel string) error {
	var (
		logger = log.WithCtx(ctx)
		sub    = GenerateJetstreamSubject(stream, channel)
	)

	jsCfg := jetstream.StreamConfig{
		Name:     stream,
		Subjects: []string{sub},
	}
	if _, err := js.CreateOrUpdateStream(ctx, jsCfg); err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func GracefulStop(ctx context.Context) {
	time.Sleep(20 * time.Second)
	nc.Close()
}
