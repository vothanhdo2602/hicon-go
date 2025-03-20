package service

import (
	"context"
	"fmt"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/constant"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/model/responsemodel"
	"github.com/vothanhdo2602/hicon/external/util/pstring"
	"github.com/vothanhdo2602/hicon/external/util/sftil"
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
	UpsertConfiguration(ctx context.Context, req *requestmodel.UpsertConfiguration) *responsemodel.BaseResponse
	FindByPK(ctx context.Context, req *requestmodel.FindByPK) *responsemodel.BaseResponse
	FindOne(ctx context.Context, req *requestmodel.FindOne) *responsemodel.BaseResponse
	FindAll(ctx context.Context, req *requestmodel.FindAll) *responsemodel.BaseResponse
	Exec(ctx context.Context, req *requestmodel.Exec) *responsemodel.BaseResponse
	BulkInsert(ctx context.Context, req *requestmodel.BulkInsert) *responsemodel.BaseResponse
	UpdateByPK(ctx context.Context, req *requestmodel.UpdateByPK) *responsemodel.BaseResponse
	BulkUpdateByPK(ctx context.Context, req *requestmodel.BulkUpdateByPK) *responsemodel.BaseResponse
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
	var (
		models    = map[string][]reflect.StructField{}
		ptrModels = map[string][]reflect.StructField{}
		refModels = map[string][]reflect.StructField{}
	)

	for _, v := range dbCfg.ModelRegistry.TableConfigurations {
		instance, ptrInstance, refInstance, err := orm.BuildEntity(v)
		if err != nil {
			return responsemodel.R400(err.Error())
		}
		models[v.Name] = instance
		ptrModels[v.Name] = ptrInstance
		refModels[v.Name] = refInstance
	}

	for _, v := range dbCfg.ModelRegistry.TableConfigurations {
		if err = orm.MapRelationToEntity(v, models, refModels); err != nil {
			return responsemodel.R400(err.Error())
		}
	}

	dbCfg.ModelRegistry.Models = models
	dbCfg.ModelRegistry.PtrModels = ptrModels
	dbCfg.ModelRegistry.RefModels = refModels

	if err = orm.Init(ctx, nil, dbCfg); err != nil {
		return responsemodel.R400(err.Error())
	}

	rdCfg := config.NewRedisConfiguration(req)
	if err = rd.Init(ctx, nil, rdCfg); err != nil {
		return responsemodel.R400(err.Error())
	}

	config.SetDBConfiguration(dbCfg)
	config.SetRedisConfiguration(rdCfg)

	return responsemodel.R200(nil, false)
}

func (s *sqlExecutorImpl) FindByPK(ctx context.Context, req *requestmodel.FindByPK) *responsemodel.BaseResponse {
	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	var (
		d     = dao.SQLExecutor()
		dbCfg = config.GetENV().DB.DBConfiguration
		mp    = &constant.ModelParams{
			Database:     dbCfg.GetDatabaseName(),
			Table:        req.Table,
			DisableCache: isDisableCache(req.DisableCache),
			ModeType:     config.DefaultModelType,
		}
		tableRegistry = config.GetModelRegistry().TableConfigurations[mp.Table]
		arrKeys       []string
	)

	if err := req.Validate(); err != nil {
		return responsemodel.R400(err.Error())
	}

	if tableRegistry == nil {
		return responsemodel.R400(fmt.Sprintf("table %s not found in registerd model", req.Table))
	}

	mapData := req.Data.(map[string]interface{})
	for k := range tableRegistry.PrimaryColumns {
		v, ok := mapData[k]
		if !ok {
			return responsemodel.R400(fmt.Sprintf("primary key column %v not found in registered model, please check func update configuration", k))
		}
		arrKeys = append(arrKeys, pstring.InterfaceToString(v))
	}

	id := strings.Join(arrKeys, ";")

	data, err, shared := d.FindByPK(ctx, req.Data, id, mp)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	if len(req.Select) > 0 {
		newModel, err := config.TransformModel(req.Table, req.Select, data, config.PtrModelType)
		if err != nil {
			return responsemodel.R400(err.Error())
		}

		return responsemodel.R200(newModel, shared)
	}

	return responsemodel.R200(data, shared)
}

func (s *sqlExecutorImpl) FindOne(ctx context.Context, req *requestmodel.FindOne) *responsemodel.BaseResponse {
	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	var (
		d     = dao.SQLExecutor()
		dbCfg = config.GetENV().DB.DBConfiguration
		mp    = &constant.ModelParams{
			Database:     dbCfg.GetDatabaseName(),
			Table:        req.Table,
			DisableCache: isDisableCache(req.DisableCache),
		}
		m   = config.GetModelRegistry().GetNewModel(mp.Table)
		sql = orm.GetDB().NewSelect().Model(m)
	)

	if err := req.Validate(); err != nil {
		return responsemodel.R400(err.Error())
	}

	sql.Offset(req.Offset)

	for _, v := range req.Where {
		sql.Where(v.Query, v.Args...)
	}

	for _, v := range req.Relations {
		sql.Relation(pstring.Title(v))
	}

	for _, v := range req.Join {
		sql.Join(v.Join, v.Args...)
	}

	for _, v := range req.OrderBy {
		sql.Order(v)
	}

	data, err, shared := d.FindOne(ctx, sql, m, mp)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	if len(req.Select) > 0 {
		newModel, err := config.TransformModel(req.Table, req.Select, data, config.DefaultModelType)
		if err != nil {
			return responsemodel.R400(err.Error())
		}

		return responsemodel.R200(newModel, shared)
	}

	return responsemodel.R200(data, shared)
}

func (s *sqlExecutorImpl) FindAll(ctx context.Context, req *requestmodel.FindAll) *responsemodel.BaseResponse {
	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	var (
		d     = dao.SQLExecutor()
		dbCfg = config.GetENV().DB.DBConfiguration
		mp    = &constant.ModelParams{
			Database:     dbCfg.GetDatabaseName(),
			Table:        req.Table,
			DisableCache: isDisableCache(req.DisableCache),
		}
		models = config.GetModelRegistry().GetNewSliceModel(mp.Table)
		sql    = orm.GetDB().NewSelect().Model(models)
	)

	if err := req.Validate(); err != nil {
		return responsemodel.R400(err.Error())
	}

	if req.Limit > 0 {
		sql.Limit(req.Limit)
	}

	sql.Offset(req.Offset)

	for _, v := range req.Where {
		sql.Where(v.Query, v.Args...)
	}

	for _, v := range req.Relations {
		sql.Relation(pstring.Title(v))
	}

	for _, v := range req.OrderBy {
		sql.Order(v)
	}

	data, err, shared := d.FindAll(ctx, sql, models, mp)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	if len(req.Select) > 0 {
		newModels, err := config.TransformModels(req.Table, req.Select, data, config.DefaultModelType)
		if err != nil {
			return responsemodel.R400(err.Error())
		}

		return responsemodel.R200(newModels, shared)
	}

	return responsemodel.R200(data, shared)
}

func (s *sqlExecutorImpl) Exec(ctx context.Context, req *requestmodel.Exec) *responsemodel.BaseResponse {
	var (
		d = dao.SQLExecutor()
	)

	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	data, err, shared := d.Exec(ctx, req.SQL, req.LockKey, req.Args...)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	return responsemodel.R200(data, shared)
}

func (s *sqlExecutorImpl) BulkInsert(ctx context.Context, req *requestmodel.BulkInsert) *responsemodel.BaseResponse {
	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	fn := func() (interface{}, error) {
		var (
			d     = dao.SQLExecutor()
			dbCfg = config.GetENV().DB.DBConfiguration
			mp    = &constant.ModelParams{
				Database:     dbCfg.GetDatabaseName(),
				Table:        req.Table,
				DisableCache: isDisableCache(req.DisableCache),
			}
		)

		if err := req.Validate(); err != nil {
			return nil, err
		}

		models, err := config.TransformModels(req.Table, nil, req.Data, config.PtrModelType)
		if err != nil {
			return nil, err
		}

		return d.BulkInsert(ctx, models, mp)
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

func (s *sqlExecutorImpl) UpdateByPK(ctx context.Context, req *requestmodel.UpdateByPK) *responsemodel.BaseResponse {
	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	fn := func() (interface{}, error) {
		var (
			d     = dao.SQLExecutor()
			dbCfg = config.GetENV().DB.DBConfiguration
			mp    = &constant.ModelParams{
				Database:     dbCfg.GetDatabaseName(),
				Table:        req.Table,
				DisableCache: isDisableCache(req.DisableCache),
			}
			sql = orm.GetDB().NewUpdate()
		)

		if err := req.Validate(); err != nil {
			return nil, err
		}

		for _, v := range req.Where {
			sql.Where(v.Query, v.Args...)
		}

		m, err := config.TransformModel(req.Table, nil, req.Data, config.PtrModelType)
		if err != nil {
			return nil, err
		}

		return d.UpdateByPK(ctx, sql, m, mp)
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

func (s *sqlExecutorImpl) BulkUpdateByPK(ctx context.Context, req *requestmodel.BulkUpdateByPK) *responsemodel.BaseResponse {
	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	fn := func() (interface{}, error) {
		var (
			d     = dao.SQLExecutor()
			dbCfg = config.GetENV().DB.DBConfiguration
			mp    = &constant.ModelParams{
				Database:     dbCfg.GetDatabaseName(),
				Table:        req.Table,
				DisableCache: isDisableCache(req.DisableCache),
			}
			sql = orm.GetDB().NewUpdate().Column(req.Set...)
		)

		if err := req.Validate(); err != nil {
			return nil, err
		}

		m, err := config.TransformModels(req.Table, nil, req.Data, config.PtrModelType)
		if err != nil {
			return nil, err
		}

		return d.BulkUpdateByPK(ctx, sql, m, mp)
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
