package service

import (
	"context"
	"fmt"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/model/entity"
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
	Connect(ctx context.Context, req *requestmodel.Credential) *responsemodel.BaseResponse
	UpsertConfig(ctx context.Context, req *requestmodel.UpsertConfig) *responsemodel.BaseResponse
	FindByPK(ctx context.Context, req *requestmodel.FindByPK) *responsemodel.BaseResponse
	FindOne(ctx context.Context, req *requestmodel.FindOne) *responsemodel.BaseResponse
	FindAll(ctx context.Context, req *requestmodel.FindAll) *responsemodel.BaseResponse
	Exec(ctx context.Context, req *requestmodel.Exec) *responsemodel.BaseResponse
	BulkInsert(ctx context.Context, req *requestmodel.BulkInsert) *responsemodel.BaseResponse
	UpdateByPK(ctx context.Context, req *requestmodel.UpdateByPK) *responsemodel.BaseResponse
	UpdateAll(ctx context.Context, req *requestmodel.UpdateAll) *responsemodel.BaseResponse
	BulkUpdateByPK(ctx context.Context, req *requestmodel.BulkUpdateByPK) *responsemodel.BaseResponse
	DeleteByPK(ctx context.Context, req *requestmodel.DeleteByPK) *responsemodel.BaseResponse
	BulkWriteWithTx(ctx context.Context, req *requestmodel.BulkWriteWithTx) *responsemodel.BaseResponse
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

func (s *sqlExecutorImpl) Connect(ctx context.Context, req *requestmodel.Credential) *responsemodel.BaseResponse {
	if err := req.Validate(); err != nil {
		return responsemodel.R400(err.Error())
	}

	return responsemodel.R200(nil, false)
}

func (s *sqlExecutorImpl) UpsertConfig(ctx context.Context, req *requestmodel.UpsertConfig) *responsemodel.BaseResponse {
	if err := req.Validate(); err != nil {
		return responsemodel.R400(err.Error())
	}

	dbCfg, err := config.NewDBConfig(req)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	// build struct
	var (
		models    = map[string][]reflect.StructField{}
		ptrModels = map[string][]reflect.StructField{}
		refModels = map[string][]reflect.StructField{}
	)

	for _, v := range dbCfg.ModelRegistry.TableConfigs {
		instance, ptrInstance, refInstance, err := orm.BuildEntity(v)
		if err != nil {
			return responsemodel.R400(err.Error())
		}
		models[v.Name] = instance
		ptrModels[v.Name] = ptrInstance
		refModels[v.Name] = refInstance
	}

	for _, v := range dbCfg.ModelRegistry.TableConfigs {
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

	config.SetDBConfig(dbCfg)
	config.SetRedisConfiguration(rdCfg)

	return responsemodel.R200(nil, false)
}

func (s *sqlExecutorImpl) FindByPK(ctx context.Context, req *requestmodel.FindByPK) *responsemodel.BaseResponse {
	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	var (
		d     = dao.SQLExecutor()
		dbCfg = config.GetENV().DB.DBConfig
		mp    = &entity.ModelParams{
			Database:            dbCfg.GetDatabaseName(),
			Table:               req.Table,
			DisableCache:        isDisableCache(req.DisableCache),
			ModeType:            config.DefaultModelType,
			WhereAllWithDeleted: req.WhereAllWithDeleted,
		}
		tableRegistry = config.GetModelRegistry().TableConfigs[mp.Table]
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

	if len(req.Selects) > 0 {
		newModel, err := config.TransformModel(req.Table, req.Selects, data, config.PtrModelType)
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
		dbCfg = config.GetENV().DB.DBConfig
		mp    = &entity.ModelParams{
			Database:            dbCfg.GetDatabaseName(),
			Table:               req.Table,
			DisableCache:        isDisableCache(req.DisableCache),
			WhereAllWithDeleted: req.WhereAllWithDeleted,
		}
		m   = config.GetModelRegistry().GetNewModel(mp.Table)
		sql = orm.GetDB().NewSelect().Model(m)
	)

	if err := req.Validate(); err != nil {
		return responsemodel.R400(err.Error())
	}

	sql.Limit(1)
	sql.Offset(req.Offset)

	if mp.WhereAllWithDeleted {
		sql.WhereAllWithDeleted()
	}

	for _, v := range req.Where {
		sql.Where(v.Query, v.Args...)
	}

	for _, v := range req.Relations {
		sql.Relation(pstring.Title(v))
	}

	for _, v := range req.Joins {
		sql.Join(v.Join, v.Args...)
	}

	for _, v := range req.OrderBy {
		sql.Order(v)
	}

	data, err, shared := d.FindOne(ctx, sql, m, mp)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	if len(req.Selects) > 0 {
		newModel, err := config.TransformModel(req.Table, req.Selects, data, config.DefaultModelType)
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
		dbCfg = config.GetENV().DB.DBConfig
		mp    = &entity.ModelParams{
			Database:            dbCfg.GetDatabaseName(),
			Table:               req.Table,
			DisableCache:        isDisableCache(req.DisableCache),
			WhereAllWithDeleted: req.WhereAllWithDeleted,
		}
		models = config.GetModelRegistry().GetNewSliceModel(mp.Table)
		sql    = orm.GetDB().NewSelect().Model(models)
	)

	if err := req.Validate(); err != nil {
		return responsemodel.R400(err.Error())
	}

	sql.Limit(req.Limit)
	sql.Offset(req.Offset)

	if mp.WhereAllWithDeleted {
		sql.WhereAllWithDeleted()
	}

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

	if len(req.Selects) > 0 {
		newModels, err := config.TransformModels(req.Table, req.Selects, data, config.DefaultModelType)
		if err != nil {
			return responsemodel.R400(err.Error())
		}

		return responsemodel.R200(newModels, shared)
	}

	return responsemodel.R200(data, shared)
}

func (s *sqlExecutorImpl) Exec(ctx context.Context, req *requestmodel.Exec) *responsemodel.BaseResponse {
	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	data, err, shared := s.exec(ctx, nil, req)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	return responsemodel.R200(data, shared)
}

func (s *sqlExecutorImpl) BulkInsert(ctx context.Context, req *requestmodel.BulkInsert) *responsemodel.BaseResponse {
	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	r, err, shared := s.bulkInsert(ctx, nil, req)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	return responsemodel.R200(r, shared)
}

func (s *sqlExecutorImpl) UpdateByPK(ctx context.Context, req *requestmodel.UpdateByPK) *responsemodel.BaseResponse {
	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	r, err, shared := s.updateByPK(ctx, nil, req)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	return responsemodel.R200(r, shared)
}

func (s *sqlExecutorImpl) UpdateAll(ctx context.Context, req *requestmodel.UpdateAll) *responsemodel.BaseResponse {
	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	r, err, shared := s.updateAll(ctx, nil, req)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	return responsemodel.R200(r, shared)
}

func (s *sqlExecutorImpl) BulkUpdateByPK(ctx context.Context, req *requestmodel.BulkUpdateByPK) *responsemodel.BaseResponse {
	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	r, err, shared := s.bulkUpdateByPK(ctx, nil, req)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	return responsemodel.R200(r, shared)
}

func (s *sqlExecutorImpl) DeleteByPK(ctx context.Context, req *requestmodel.DeleteByPK) *responsemodel.BaseResponse {
	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	r, err, shared := s.deleteByPK(ctx, nil, req)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	return responsemodel.R200(r, shared)
}

func (s *sqlExecutorImpl) BulkWriteWithTx(ctx context.Context, req *requestmodel.BulkWriteWithTx) *responsemodel.BaseResponse {
	if err := config.ConfigurationUpdated(); err != nil {
		return responsemodel.R400(err.Error())
	}

	fn := func() (interface{}, error) {
		tx, err := orm.BeginTx(ctx)
		defer orm.HandleTxErr(ctx, tx, err)
		if err != nil {
			return nil, err
		}

		if err = req.Validate(); err != nil {
			return nil, err
		}

		err = s.processBulkWrite(ctx, &tx, req)

		return req.Operations, err
	}

	if req.LockKey == "" {
		r, err := fn()
		if err != nil {
			return responsemodel.R400(err.Error())
		}

		return responsemodel.R200(r, false)
	}

	var (
		key = dao.GetBulkWriteWithTxLockKey(req.LockKey)
	)

	r, err, shared := s.g.Do(key, fn)
	if err != nil {
		return responsemodel.R400(err.Error())
	}

	return responsemodel.R200(r, shared)
}
