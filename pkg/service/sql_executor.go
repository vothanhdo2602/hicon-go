package service

import (
	"context"
	"fmt"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/model/responsemodel"
	"github.com/vothanhdo2602/hicon/external/util/pstring"
	"github.com/vothanhdo2602/hicon/external/util/sftil"
	"github.com/vothanhdo2602/hicon/internal/orm"
	"github.com/vothanhdo2602/hicon/pkg/dao"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect"
	"strings"
)

var (
	sqlExecutor *sqlExecutorImpl
)

type SQLExecutorInterface interface {
	UpsertConfiguration(ctx context.Context, req *requestmodel.UpsertConfiguration) *responsemodel.BaseResponse
	FindByPrimaryKeys(ctx context.Context, req *requestmodel.FindByPrimaryKeys) *responsemodel.BaseResponse
	FindOne(ctx context.Context, req *requestmodel.FindOne) *responsemodel.BaseResponse
	FindAll(ctx context.Context, req *requestmodel.FindAll) *responsemodel.BaseResponse
	Exec(ctx context.Context, req *requestmodel.Exec) *responsemodel.BaseResponse
	BulkInsert(ctx context.Context, req *requestmodel.BulkInsert) *responsemodel.BaseResponse
}

type sqlExecutorImpl struct {
	g   sftil.Group
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

func (s *sqlExecutorImpl) UpsertConfiguration(ctx context.Context, req *requestmodel.UpsertConfiguration) *responsemodel.BaseResponse {
	if err := req.Validate(); err != nil {
		return responsemodel.R400(err.Error())
	}

	dbCfg, err := config.NewDBConfiguration(req)
	if err != nil {
		return responsemodel.R400(err.Error())
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

	if err := req.Validate(); err != nil {
		return responsemodel.R400(err.Error())
	}

	if tableRegistry == nil {
		return responsemodel.R400(fmt.Sprintf("table %s not found in registerd model", req.Table))
	}

	for k := range tableRegistry.PrimaryColumns {
		v, ok := req.Data[k]
		if !ok {
			return responsemodel.R400(fmt.Sprintf("primary key column %v not found in registered model, please check func update configuration", k))
		}
		arrKeys = append(arrKeys, pstring.InterfaceToString(v))
	}

	id := strings.Join(arrKeys, ";")

	data, err, shared := d.FindByPrimaryKeys(ctx, req.Data, id, md)
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

	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	if err := req.Validate(); err != nil {
		return responsemodel.R400(err.Error())
	}

	sql.Offset(req.Offset)

	for _, v := range req.Where {
		sql.Where(v.Condition, v.Args...)
	}

	for _, v := range req.Relations {
		sql.Relation(cases.Title(language.English).String(v))
	}

	for _, v := range req.Join {
		sql.Join(v.Join, v.Args...)
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
		m   = config.GetModelRegistry().GetNewSliceModel(md.Table)
		sql = orm.GetDB().NewSelect().Model(m)
	)

	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	if err := req.Validate(); err != nil {
		return responsemodel.R400(err.Error())
	}

	if req.Limit > 0 {
		sql.Limit(req.Limit)
	}

	sql.Offset(req.Offset)

	for _, v := range req.Where {
		sql.Where(v.Condition, v.Args...)
	}

	for _, v := range req.Relations {
		sql.Relation(cases.Title(language.English).String(v))
	}

	for _, v := range req.OrderBy {
		sql.Order(v)
	}

	data, err, shared := d.Scan(ctx, sql, m, md)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	if len(req.Select) > 0 {
		newModels, err := config.TransformModels(req.Table, req.Select, data)
		if err != nil {
			return responsemodel.R400(err.Error())
		}

		return responsemodel.R200(newModels, shared)
	}

	return responsemodel.R200(data, shared)
}

func (s *sqlExecutorImpl) Exec(ctx context.Context, req *requestmodel.Exec) *responsemodel.BaseResponse {
	var (
		d     = dao.SQLExecutor()
		dbCfg = config.GetENV().DB.DBConfiguration
		md    = &dao.ModelParams{
			Database: dbCfg.GetDatabaseName(),
			//Table:        req.Table,
			//DisableCache: isDisableCache(req.DisableCache),
		}
	)

	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	data, err, shared := d.Exec(ctx, req.SQL, req.LockKey, md, req.Args...)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	return responsemodel.R200(data, shared)
}

func (s *sqlExecutorImpl) BulkInsert(ctx context.Context, req *requestmodel.BulkInsert) *responsemodel.BaseResponse {
	var (
		d     = dao.SQLExecutor()
		dbCfg = config.GetENV().DB.DBConfiguration
		md    = &dao.ModelParams{
			Database:     dbCfg.GetDatabaseName(),
			Table:        req.Table,
			DisableCache: isDisableCache(req.DisableCache),
		}
	)

	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	if err := req.Validate(); err != nil {
		return responsemodel.R400(err.Error())
	}

	models, err := config.TransformModels(req.Table, []string{}, req.Data)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	fn := func() (interface{}, error) {
		return d.BulkInsert(ctx, models, md)
	}

	if req.LockKey == "" {
		r, err := fn()
		if err != nil {
			return responsemodel.R400(err.Error())
		}

		return responsemodel.R200(r, false)
	}

	var (
		key = dao.GetCustomLockKey(req.LockKey)
	)

	r, err, shared := s.g.Do(key, fn)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	return responsemodel.R200(r, shared)
}
