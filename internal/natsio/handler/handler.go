package handler

import (
	"github.com/nats-io/nats.go"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/util/commontil"
	"github.com/vothanhdo2602/hicon/external/util/natstil"
	"github.com/vothanhdo2602/hicon/external/util/pjson"
)

func FindByPrimaryKeys(msg *nats.Msg) {
	var (
		ctx  = natstil.GetContext(msg)
		data requestmodel.FindByPrimaryKeys
	)

	defer commontil.Recover(ctx)
	defer natstil.Response(ctx, msg)

	if err := config.ConfigurationUpdated(); err != nil {
		msg.Data = natstil.R400(err.Error())
		return
	}

	if err := pjson.Unmarshal(ctx, msg.Data, &data); err != nil {
		msg.Data = natstil.R400(err.Error())
		return
	}

	if err := data.Validate(); err != nil {
		msg.Data = natstil.R400(err.Error())
		return
	}

	//var (
	//	svc = service.SQLExecutor[interface{}]()
	//)

	//_, err, _ := svc.FindByPrimaryKeys(ctx, &data)
	//if err != nil {
	//	msg.Data = natstil.R400(err.Error())
	//	return
	//}
}

func FindOne(msg *nats.Msg) {
	var (
		ctx  = natstil.GetContext(msg)
		data requestmodel.FindOne
	)

	defer commontil.Recover(ctx)
	defer natstil.Response(ctx, msg)

	if err := config.ConfigurationUpdated(); err != nil {
		msg.Data = natstil.R400(err.Error())
		return
	}

	if err := pjson.Unmarshal(ctx, msg.Data, &data); err != nil {
		msg.Data = natstil.R400(err.Error())
		return
	}

	if err := data.Validate(); err != nil {
		msg.Data = natstil.R400(err.Error())
		return
	}
}
