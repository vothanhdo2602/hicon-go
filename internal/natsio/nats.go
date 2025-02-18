package natsio

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/external/util/natstil"
	"sync"
)

var (
	nc *nats.Conn
	js jetstream.JetStream
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

	js, err = jetstream.New(nc)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	return
}

func GetNC() *nats.Conn {
	return nc
}

func GetJs() jetstream.JetStream {
	return js
}

func Subscribe(ctx context.Context) (err error) {
	if err = ReqrepQueueSubscribe(ctx, updateConfigSubject, UpdateConfig); err != nil {
		return
	}
	if err = ReqrepQueueSubscribe(ctx, updateConfigSubject, UpdateConfig); err != nil {
		return
	}

	return
}

func ReqrepQueueSubscribe(ctx context.Context, channel string, cb func(msg *nats.Msg)) error {
	var (
		logger = log.WithCtx(ctx)
		sub    = natstil.GenerateReqrepSubject(channel)
		queue  = natstil.GenerateQueueNameFromSubject(sub)
	)

	if _, err := nc.QueueSubscribe(sub, queue, cb); err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func JetstreamQueueSubscribe(ctx context.Context, stream, channel string) error {
	var (
		logger = log.WithCtx(ctx)
		sub    = natstil.GenerateJetstreamSubject(stream, channel)
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
