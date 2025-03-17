package handler

import (
	"github.com/nats-io/nats.go"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/util/commontil"
	"github.com/vothanhdo2602/hicon/external/util/natstil"
	"github.com/vothanhdo2602/hicon/external/util/pjson"
	"github.com/vothanhdo2602/hicon/pkg/service"
)

//func UpsertConfiguration(msg *nats.Msg) {
//	var (
//		ctx  = context.Background()
//		data requestmodel.UpsertConfiguration
//	)
//
//	defer commontil.Recover(ctx)
//	defer natstil.Response(ctx, msg)
//
//	if err := pjson.Unmarshal(ctx, msg.Data, &data); err != nil {
//		msg.Data = natstil.R400(err.Error())
//		return
//	}
//
//	if err := data.Validate(); err != nil {
//		msg.Data = natstil.R400(err.Error())
//		return
//	}
//
//	dbCfg := &config.DBConfiguration{
//		Type:         data.DBConfiguration.Type,
//		Host:         data.DBConfiguration.Host,
//		Port:         data.DBConfiguration.Port,
//		Username:     data.DBConfiguration.Username,
//		Password:     data.DBConfiguration.Password,
//		Database:     data.DBConfiguration.Database,
//		MaxCons:      data.DBConfiguration.MaxCons,
//		DisableCache: data.DisableCache,
//		Debug:        data.Debug,
//		ModelRegistry: &config.ModelRegistry{
//			TableConfigurations: map[string]*config.TableConfiguration{},
//			Models:              map[string][]reflect.StructField{},
//		},
//	}
//
//	if data.DBConfiguration.TLS != nil {
//		dbCfg.TLS = &config.TLS{
//			RootCAPEM: data.DBConfiguration.TLS.RootCAPEM,
//		}
//	}
//
//	for _, t := range data.TableConfigurations {
//		tblCfg := &config.TableConfiguration{
//			Name:                  t.Name,
//			ColumnConfigs:         map[string]*config.ColumnConfig{},
//			PrimaryColumns:        map[string]interface{}{},
//			RelationColumnConfigs: map[string]*config.RelationColumnConfig{},
//		}
//
//		for _, col := range t.ColumnConfigs {
//			tblCfg.ColumnConfigs[col.Name] = &config.ColumnConfig{
//				Type:         col.Type,
//				Nullable:     col.Nullable,
//				IsPrimaryKey: col.IsPrimaryKey,
//			}
//
//			if col.IsPrimaryKey {
//				tblCfg.PrimaryColumns[col.Name] = col.Name
//			}
//		}
//
//		if len(tblCfg.PrimaryColumns) == 0 {
//			msg.Data = natstil.R400(fmt.Sprintf("Table %s has no primary columns", tblCfg.Name))
//			return
//		}
//
//		for _, col := range t.RelationColumnConfigs {
//			tblCfg.RelationColumnConfigs[col.Name] = &config.RelationColumnConfig{
//				Name:     col.Name,
//				RefTable: col.RefTable,
//				Type:     col.Type,
//			}
//		}
//
//		dbCfg.ModelRegistry.TableConfigurations[t.Name] = tblCfg
//	}
//
//	// build struct
//	models := map[string][]reflect.StructField{}
//	ptrModels := map[string][]reflect.StructField{}
//	for _, v := range dbCfg.ModelRegistry.TableConfigurations {
//		instance, ptrInstance, err := orm.BuildEntity(v)
//		if err != nil {
//			msg.Data = natstil.R400(err.Error())
//			return
//		}
//		models[v.Name] = instance
//		ptrModels[v.Name] = ptrInstance
//	}
//
//	for _, v := range dbCfg.ModelRegistry.TableConfigurations {
//		if err := orm.MapRelationToEntity(v, models); err != nil {
//			msg.Data = natstil.R400(err.Error())
//			return
//		}
//
//		dbCfg.ModelRegistry.Models[v.Name] = models[v.Name]
//		dbCfg.ModelRegistry.PtrModels[v.Name] = models[v.Name]
//	}
//
//	if err := orm.Init(ctx, nil, dbCfg); err != nil {
//		msg.Data = natstil.R400(err.Error())
//		return
//	}
//
//	config.SetDBConfiguration(dbCfg)
//
//	msg.Data = natstil.R200(nil)
//}

func FindByPrimaryKeys(msg *nats.Msg) {
	go func(msg *nats.Msg) {
		var (
			ctx  = natstil.GetContext(msg)
			data requestmodel.FindByPrimaryKeys
		)

		defer commontil.Recover(ctx)
		defer natstil.Response(ctx, msg)

		if err := pjson.Unmarshal(ctx, msg.Data, &data); err != nil {
			msg.Data = natstil.R400(err.Error())
			return
		}

		var (
			svc = service.SQLExecutor()
		)

		resp := svc.FindByPrimaryKeys(ctx, &data)
		msg.Data = natstil.R200(resp)
	}(msg)
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
