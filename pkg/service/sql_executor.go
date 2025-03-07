package service

import (
	"context"
	"fmt"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/model/responsemodel"
	"github.com/vothanhdo2602/hicon/external/util/pstring"
	"github.com/vothanhdo2602/hicon/internal/orm"
	"github.com/vothanhdo2602/hicon/pkg/dao"
	"reflect"
	"strings"
)

var (
	sqlExecutor *sqlExecutorImpl
)

type SQLExecutorInterface interface {
	UpdateConfiguration(ctx context.Context, req *requestmodel.UpsertConfiguration) *responsemodel.BaseResponse
	FindByPrimaryKeys(ctx context.Context, req *requestmodel.FindByPrimaryKeys) *responsemodel.BaseResponse
}

type sqlExecutorImpl struct {
	dao dao.SQLExecutorInterface
}

func SQLExecutor() SQLExecutorInterface {
	if sqlExecutor == nil {
		sqlExecutor = &sqlExecutorImpl{
			dao: dao.SQLExecutor(),
		}
	}
	return sqlExecutor
}

func (s *sqlExecutorImpl) UpdateConfiguration(ctx context.Context, req *requestmodel.UpsertConfiguration) *responsemodel.BaseResponse {
	if err := req.Validate(); err != nil {
		return responsemodel.R400(err.Error())
	}

	dbCfg := &config.DBConfiguration{
		Type:         req.DBConfiguration.Type,
		Host:         req.DBConfiguration.Host,
		Port:         req.DBConfiguration.Port,
		Username:     req.DBConfiguration.Username,
		Password:     req.DBConfiguration.Password,
		Database:     req.DBConfiguration.Database,
		MaxCons:      req.DBConfiguration.MaxCons,
		DisableCache: req.DisableCache,
		Debug:        req.Debug,
		ModelRegistry: &config.ModelRegistry{
			TableConfigurations: map[string]*config.TableConfiguration{},
			Models:              map[string][]reflect.StructField{},
		},
	}

	if req.DBConfiguration.TLS != nil {
		dbCfg.TLS = &config.TLS{
			RootCAPEM: req.DBConfiguration.TLS.RootCAPEM,
		}
	}

	for _, t := range req.TableConfigurations {
		tblCfg := &config.TableConfiguration{
			Name:                  t.Name,
			ColumnConfigs:         map[string]*config.ColumnConfig{},
			PrimaryColumns:        map[string]interface{}{},
			RelationColumnConfigs: map[string]*config.RelationColumnConfig{},
		}

		for _, col := range t.ColumnConfigs {
			tblCfg.ColumnConfigs[col.Name] = &config.ColumnConfig{
				Type:         col.Type,
				Nullable:     col.Nullable,
				IsPrimaryKey: col.IsPrimaryKey,
			}

			if col.IsPrimaryKey {
				tblCfg.PrimaryColumns[col.Name] = col.Name
			}
		}

		if len(tblCfg.PrimaryColumns) == 0 {
			return responsemodel.R400(fmt.Sprintf("Table %s has no primary columns", tblCfg.Name))
		}

		for _, col := range t.RelationColumnConfigs {
			tblCfg.RelationColumnConfigs[col.Name] = &config.RelationColumnConfig{
				Name:     col.Name,
				RefTable: col.RefTable,
				Type:     col.Type,
			}
		}

		dbCfg.ModelRegistry.TableConfigurations[t.Name] = tblCfg
	}

	// build struct
	builderEntities := map[string][]reflect.StructField{}
	for _, v := range dbCfg.ModelRegistry.TableConfigurations {
		instance, err := orm.BuildEntity(v)
		if err != nil {
			return responsemodel.R400(err.Error())
		}
		builderEntities[v.Name] = instance
	}

	for _, v := range dbCfg.ModelRegistry.TableConfigurations {
		if err := orm.MapRelationToEntity(v, builderEntities); err != nil {
			return responsemodel.R400(err.Error())
		}

		dbCfg.ModelRegistry.Models[v.Name] = builderEntities[v.Name]
	}

	if err := orm.Init(ctx, nil, dbCfg); err != nil {
		return responsemodel.R400(err.Error())
	}

	config.SetDBConfiguration(dbCfg)

	return responsemodel.R200(nil, false)
}

func (s *sqlExecutorImpl) FindByPrimaryKeys(ctx context.Context, req *requestmodel.FindByPrimaryKeys) *responsemodel.BaseResponse {
	var (
		d     = dao.SQLExecutor()
		dbCfg = config.GetENV().DB.DBConfiguration
		md    = &dao.ModelParams{
			Database: dbCfg.Database,
			Table:    req.Table,
		}
		tableRegistry = dbCfg.ModelRegistry.TableConfigurations[md.Table]
		arrKeys       []string
	)

	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	//if err := req.Validate(); err != nil {
	//	return responsemodel.R400(err.Error())
	//}

	if tableRegistry == nil {
		return responsemodel.R400(fmt.Sprintf("table %s not found in registerd model", req.Table))
	}

	if dbCfg.DisableCache || req.DisableCache {
		md.DisableCache = req.DisableCache
	}

	for k := range tableRegistry.PrimaryColumns {
		v, ok := req.PrimaryKeys[k]
		if !ok {
			return responsemodel.R400(fmt.Sprintf("primary key column %v not found in registered model, please check func update configuration", k))
		}
		arrKeys = append(arrKeys, pstring.InterfaceToString(v))
	}

	id := strings.Join(arrKeys, ";")

	data, err, shared := d.FindByPrimaryKeys(ctx, req.PrimaryKeys, id, md)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	return responsemodel.R200(data, shared)
}
