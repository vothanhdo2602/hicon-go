package service

import (
	"context"
	"errors"
	"github.com/uptrace/bun"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/constant"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/util/pjson"
	"github.com/vothanhdo2602/hicon/internal/orm"
	"github.com/vothanhdo2602/hicon/pkg/dao"
)

func (s *sqlExecutorImpl) bulkInsert(ctx context.Context, tx bun.IDB, req *requestmodel.BulkInsert) (interface{}, error, bool) {
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

		if tx == nil {
			tx = orm.GetDB()
		}

		return d.BulkInsert(ctx, tx, models, mp)
	}

	if req.LockKey == "" {
		r, err := fn()
		return r, err, false
	}

	var (
		key = dao.GetBulkInsertLockKey(req.LockKey)
	)

	r, err, shared := s.g.Do(key, fn)
	return r, err, shared
}

func (s *sqlExecutorImpl) updateByPK(ctx context.Context, tx bun.IDB, req *requestmodel.UpdateByPK) (interface{}, error, bool) {

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

		if tx == nil {
			tx = orm.GetDB()
		}
		sql := tx.NewUpdate()

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
		return r, err, false
	}

	var (
		key = dao.GetUpdateByPKLockKey(req.LockKey)
	)

	r, err, shared := s.g.Do(key, fn)
	return r, err, shared
}

func (s *sqlExecutorImpl) updateAll(ctx context.Context, tx bun.IDB, req *requestmodel.UpdateAll) (interface{}, error, bool) {
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

		if tx == nil {
			tx = orm.GetDB()
		}
		sql := tx.NewUpdate()

		for _, v := range req.Where {
			sql.Where(v.Query, v.Args...)
		}

		for _, v := range req.Set {
			sql.Set(v.Query, v.Args...)
		}

		if req.WhereAllWithDeleted {
			sql.WhereAllWithDeleted()
		}

		m := config.GetModelRegistry().GetNewModel(req.Table)

		return d.UpdateAll(ctx, sql, m, mp)
	}

	if req.LockKey == "" {
		r, err := fn()
		return r, err, false
	}

	var (
		key = dao.GetUpdateByPKLockKey(req.LockKey)
	)

	r, err, shared := s.g.Do(key, fn)
	return r, err, shared
}

func (s *sqlExecutorImpl) bulkUpdateByPK(ctx context.Context, tx bun.IDB, req *requestmodel.BulkUpdateByPK) (interface{}, error, bool) {
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

		m, err := config.TransformModels(req.Table, nil, req.Data, config.PtrModelType)
		if err != nil {
			return nil, err
		}

		if tx == nil {
			tx = orm.GetDB()
		}
		sql := tx.NewUpdate().Column(req.Set...)

		return d.BulkUpdateByPK(ctx, sql, m, mp)
	}

	if req.LockKey == "" {
		r, err := fn()
		return r, err, false
	}

	var (
		key = dao.GetBulkUpdateByPKLockKey(req.LockKey)
	)

	r, err, shared := s.g.Do(key, fn)
	return r, err, shared
}

func (s *sqlExecutorImpl) deleteByPK(ctx context.Context, tx bun.IDB, req *requestmodel.DeleteByPK) (interface{}, error, bool) {
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

		if tx == nil {
			tx = orm.GetDB()
		}
		sql := tx.NewDelete()

		if req.ForceDelete {
			sql.ForceDelete()
		}

		for _, v := range req.Where {
			sql.Where(v.Query, v.Args...)
		}

		m, err := config.TransformModel(req.Table, nil, req.Data, config.PtrModelType)
		if err != nil {
			return nil, err
		}

		return d.DeleteByPK(ctx, sql, m, mp)
	}

	if req.LockKey == "" {
		r, err := fn()
		return r, err, false
	}

	var (
		key = dao.GetUpdateByPKLockKey(req.LockKey)
	)

	r, err, shared := s.g.Do(key, fn)
	return r, err, shared
}

func (s *sqlExecutorImpl) processBulkWrite(ctx context.Context, tx *bun.Tx, req *requestmodel.BulkWriteWithTx) (err error) {
	for _, o := range req.Operations {
		switch o.Name {
		case constant.BWOperationBulkInsert:
			err = s.processBulkInsert(ctx, tx, o)
		case constant.BWOperationUpdateByPK:
			err = s.processUpdateByPK(ctx, tx, o)
		case constant.BWOperationUpdateAll:
			err = s.processUpdateAll(ctx, tx, o)
		case constant.BWOperationBulkUpdateByPK:
			err = s.processBulkUpdateByPK(ctx, tx, o)
		case constant.BWOperationDeleteByPK:
			err = s.processDeleteByPK(ctx, tx, o)
		default:
			return errors.New("operation not supported")
		}

		if err != nil {
			return err
		}
	}

	return err
}

func (s *sqlExecutorImpl) processBulkInsert(ctx context.Context, tx *bun.Tx, o *requestmodel.Operation) error {
	data, err := pjson.ConvertWithType[requestmodel.BulkInsert](o.Data)
	if err != nil {
		return err
	}
	data.LockKey = ""
	_, err, _ = s.bulkInsert(ctx, tx, &data)
	return err
}

func (s *sqlExecutorImpl) processUpdateByPK(ctx context.Context, tx *bun.Tx, o *requestmodel.Operation) error {
	data, err := pjson.ConvertWithType[requestmodel.UpdateByPK](o.Data)
	if err != nil {
		return err
	}

	data.LockKey = ""
	_, err, _ = s.updateByPK(ctx, tx, &data)
	return err
}

func (s *sqlExecutorImpl) processUpdateAll(ctx context.Context, tx *bun.Tx, o *requestmodel.Operation) error {
	data, err := pjson.ConvertWithType[requestmodel.UpdateAll](o.Data)
	if err != nil {
		return err
	}
	data.LockKey = ""
	_, err, _ = s.updateAll(ctx, tx, &data)
	return err
}

func (s *sqlExecutorImpl) processBulkUpdateByPK(ctx context.Context, tx *bun.Tx, o *requestmodel.Operation) error {
	data, err := pjson.ConvertWithType[requestmodel.BulkUpdateByPK](o.Data)
	if err != nil {
		return err
	}
	data.LockKey = ""
	_, err, _ = s.bulkUpdateByPK(ctx, tx, &data)
	return err
}
func (s *sqlExecutorImpl) processDeleteByPK(ctx context.Context, tx *bun.Tx, o *requestmodel.Operation) error {
	data, err := pjson.ConvertWithType[requestmodel.DeleteByPK](o.Data)
	if err != nil {
		return err
	}
	data.LockKey = ""
	_, err, _ = s.deleteByPK(ctx, tx, &data)
	return err
}
