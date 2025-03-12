package service

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/model/responsemodel"
	"github.com/vothanhdo2602/hicon/external/util/pstring"
	"github.com/vothanhdo2602/hicon/internal/orm"
	"github.com/vothanhdo2602/hicon/internal/rd"
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
	FindOne(ctx context.Context, req *requestmodel.FindOne) *responsemodel.BaseResponse
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
			Entities:            map[string]interface{}{},
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
				Join:     col.Join,
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
			Database:     dbCfg.GetDatabaseName(),
			Table:        req.Table,
			DisableCache: isDisableCache(req.DisableCache),
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

	if len(req.Select) > 0 {
		newModel, err := config.TransformModel(req.Table, req.Select, data)
		if err != nil {
			return responsemodel.R400(err.Error())
		}

		return responsemodel.R200(newModel, shared)
	}

	return responsemodel.R200(data, shared)
}

func (s *sqlExecutorImpl) FindOne(ctx context.Context, req *requestmodel.FindOne) *responsemodel.BaseResponse {
	var (
		d     = dao.SQLExecutor()
		dbCfg = config.GetENV().DB.DBConfiguration
		md    = &dao.ModelParams{
			Database:     dbCfg.GetDatabaseName(),
			Table:        req.Table,
			DisableCache: isDisableCache(req.DisableCache),
		}
		m   = config.GetModelRegistry().GetNewModel(md.Table)
		sql = orm.GetDB().NewSelect().Model(m)
	)

	sql.Offset(req.Offset)

	for _, v := range req.Where {
		sql.Where(v.Condition, v.Args...)
	}

	for _, v := range req.Relations {
		sql.Relation(v)
	}

	for _, v := range req.OrderBy {
		sql.Order(v)
	}

	data, err, shared := d.ScanOne(ctx, sql, m, md)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	if len(req.Select) > 0 {
		newModel, err := config.TransformModel(req.Table, req.Select, data)
		if err != nil {
			return responsemodel.R400(err.Error())
		}

		return responsemodel.R200(newModel, shared)
	}

	return responsemodel.R200(data, shared)
}

func (s *sqlExecutorImpl) FindAll(ctx context.Context, req *requestmodel.FindAll) *responsemodel.BaseResponse {
	var (
		d     = dao.SQLExecutor()
		dbCfg = config.GetENV().DB.DBConfiguration
		md    = &dao.ModelParams{
			Database:     dbCfg.GetDatabaseName(),
			Table:        req.Table,
			DisableCache: isDisableCache(req.DisableCache),
		}
		m   = config.GetModelRegistry().GetNewModel(md.Table)
		sql = orm.GetDB().NewSelect().Model(m)
	)

	if req.Limit > 0 {
		sql.Limit(req.Limit)
	}

	sql.Offset(req.Offset)

	for _, v := range req.Where {
		sql.Where(v.Condition, v.Args...)
	}

	for _, v := range req.Relations {
		sql.Relation(v)
	}

	for _, v := range req.OrderBy {
		sql.Order(v)
	}

	data, err, shared := d.ScanOne(ctx, sql, m, md)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	if len(req.Select) > 0 {
		newModel, err := config.TransformModel(req.Table, req.Select, data)
		if err != nil {
			return responsemodel.R400(err.Error())
		}

		return responsemodel.R200(newModel, shared)
	}

	return responsemodel.R200(data, shared)
}

func isDisableCache(localCache bool) bool {
	var (
		env = config.GetENV()
	)
	if !env.DB.DBConfiguration.DisableCache || rd.GetRedis() == nil {
		return env.DB.DBConfiguration.DisableCache
	}

	return localCache
}

type User struct {
	bun.BaseModel `bun:"users"`
	ID            string   `bun:"id,pk" redis:"id" json:"id"`
	Type          string   `bun:"type,pk" redis:"type" json:"type"`
	Profile       *Profile `json:"profile,omitempty" bun:"rel:has-one,join:id=user_id"`
}

type Profile struct {
	ID   string `bun:"id,pk" redis:"id" json:"id"`
	Name string `bun:"name,pk" redis:"name" json:"name"`
}
