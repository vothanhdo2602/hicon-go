package natsio

import (
	"github.com/nats-io/nats.go"
	dynamicstruct "github.com/ompluscator/dynamic-struct"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/model/reqmodel"
	"github.com/vothanhdo2602/hicon/external/util/commontil"
	"github.com/vothanhdo2602/hicon/external/util/natstil"
	"github.com/vothanhdo2602/hicon/external/util/pjson"
	"github.com/vothanhdo2602/hicon/internal/orm"
)

func UpdateConfig(msg *nats.Msg) {
	var (
		ctx  = natstil.GetContext(msg)
		data reqmodel.UpdateConfig
	)

	defer commontil.Recover(ctx)
	defer natstil.Response(ctx, msg)

	if err := pjson.Unmarshal(ctx, msg.Data, &data); err != nil {
		msg.Data = natstil.R400(err.Error())
		return
	}

	if err := data.Validate(); err != nil {
		msg.Data = natstil.R400(err.Error())
		return
	}

	dbCfg := &config.DBConfig{
		Type:         data.DBConfig.Type,
		Host:         data.DBConfig.Host,
		Port:         data.DBConfig.Port,
		Username:     data.DBConfig.Username,
		Password:     data.DBConfig.Password,
		Database:     data.DBConfig.Database,
		MaxCons:      data.DBConfig.MaxCons,
		DisableCache: data.DisableCache,
		Debug:        data.Debug,
	}
	if data.DBConfig.TLS != nil {
		dbCfg.TLS = &config.TLS{
			RootCAPEM: data.DBConfig.TLS.RootCAPEM,
		}
	}

	// build struct
	builderEntities := map[string]dynamicstruct.Builder{}
	for _, v := range dbCfg.TableConfigs {
		instance, err := orm.BuildEntity(v)
		if err != nil {
			msg.Data = natstil.R400(err.Error())
			return
		}
		builderEntities[v.Name] = instance
	}

	for name, v := range builderEntities {
		dbCfg.TableStructs[name] = v.Build()
	}

	config.SetDBConfig(dbCfg)
	if err := Init(ctx, nil); err != nil {
		msg.Data = natstil.R400(err.Error())
		return
	}
}
