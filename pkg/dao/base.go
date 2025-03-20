package dao

import (
	"context"
	dbsql "database/sql"
	"errors"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/constant"
	"github.com/vothanhdo2602/hicon/external/util/commontil"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/external/util/sftil"
	"github.com/vothanhdo2602/hicon/internal/orm"
	"github.com/vothanhdo2602/hicon/internal/rd"
	"go.uber.org/zap"
	"sync"
)

const (
	FindByPKAction = "find_by_primary_keys_action"
	FindAllAction  = "find_all_action"
	CustomLockKey  = "custom_lock_key_action"
)

func GetCustomLockKey(lockKey string) string {
	return fmt.Sprintf("%s:%s", CustomLockKey, lockKey)
}

type dbInterface interface {
	String() string
	Scan(ctx context.Context, values ...any) error
	Model(model any) *bun.SelectQuery
}

type BaseInterface interface {
	FindByPK(ctx context.Context, pk interface{}, id string, mp *constant.ModelParams) (m interface{}, err error, shared bool)
	FindOne(ctx context.Context, sql dbInterface, model interface{}, mp *constant.ModelParams, values ...any) (m interface{}, err error, shared bool)
	FindAll(ctx context.Context, sql dbInterface, models interface{}, mp *constant.ModelParams, values ...any) (interface{}, error, bool)
	Exec(ctx context.Context, stringSQL, lockKey string, args ...any) (interface{}, error, bool)
	BulkInsert(ctx context.Context, models interface{}, mp *constant.ModelParams) (m interface{}, err error)
	UpdateByPK(ctx context.Context, sql *bun.UpdateQuery, m interface{}, mp *constant.ModelParams) (r interface{}, err error)
	BulkUpdateByPK(ctx context.Context, sql *bun.UpdateQuery, m interface{}, mp *constant.ModelParams) (r interface{}, err error)
}

type baseImpl struct {
	g  sftil.Group
	mu sync.Mutex
}

func (s *baseImpl) FindByPK(ctx context.Context, m interface{}, id string, mp *constant.ModelParams) (interface{}, error, bool) {
	var (
		key = fmt.Sprintf("%s:%s", FindByPKAction, id)
	)

	v, err, shared := s.g.Do(key, func() (interface{}, error) {
		newModel, _ := config.TransformModel(mp.Table, nil, m, mp.ModeType)
		return s.findByPK(ctx, newModel, mp)
	})

	return v, err, shared
}

func (s *baseImpl) findByPK(ctx context.Context, m interface{}, mp *constant.ModelParams) (interface{}, error) {
	var (
		logger = log.WithCtx(ctx)
		db     = orm.GetDB()
		sql    = db.NewSelect().Model(m)
	)

	if !mp.DisableCache {
		cacheModel := rd.HGet(ctx, mp, m)
		if cacheModel != nil {
			return cacheModel, nil
		}
	}

	err := sql.WherePK().Scan(ctx)
	if err != nil {
		if errors.Is(err, dbsql.ErrNoRows) {
			return nil, nil
		}

		logger.Error(err.Error())
		return nil, err
	}

	if !mp.DisableCache {
		rd.HSet(ctx, mp, m)
	}
	return m, nil
}

func (s *baseImpl) FindAll(ctx context.Context, sql dbInterface, models interface{}, mp *constant.ModelParams, values ...any) (interface{}, error, bool) {
	var (
		sqlStr = sql.String()
		sqlKey = fmt.Sprintf("%s:%s", FindAllAction, sqlStr)
	)

	v, err, shared := s.g.Do(sqlKey, func() (interface{}, error) {
		var (
			logger = log.WithCtx(ctx)
		)

		if !mp.DisableCache {
			cacheModel := rd.HGetAllSQL(ctx, mp, sqlStr, models, s.FindByPK)
			if cacheModel != nil {
				return cacheModel, nil
			}
		}

		err := sql.Scan(ctx, values...)
		if err != nil {
			if errors.Is(err, dbsql.ErrNoRows) {
				return nil, nil
			}

			logger.Error(err.Error())
			return nil, err
		}

		if !mp.DisableCache {
			rd.HSetAllSQL(ctx, mp, sqlStr, models)
		}

		return models, nil
	})

	if err != nil {
		return nil, err, shared
	}

	return v, err, shared
}

func (s *baseImpl) FindOne(ctx context.Context, sql dbInterface, m interface{}, mp *constant.ModelParams, values ...any) (interface{}, error, bool) {
	var (
		sqlStr = sql.String()
		sqlKey = fmt.Sprintf("%s:%s", FindAllAction, sqlStr)
	)

	v, err, shared := s.g.Do(sqlKey, func() (interface{}, error) {
		var (
			logger = log.WithCtx(ctx)
		)

		if !mp.DisableCache {
			cacheModel := rd.HGetSQL(ctx, mp, sqlStr, m, s.FindByPK)
			if cacheModel != nil {
				return cacheModel, nil
			}
		}

		err := sql.Scan(ctx, values...)
		if err != nil {
			if errors.Is(err, dbsql.ErrNoRows) {
				return nil, nil
			}

			logger.Error(err.Error())
			return nil, err
		}

		if !mp.DisableCache {
			rd.HSetSQL(ctx, mp, sqlStr, m)
		}

		return &m, err
	})

	if err != nil {
		return nil, err, shared
	}

	return v, err, shared
}

func (s *baseImpl) Exec(ctx context.Context, stringSQL, lockKey string, args ...any) (interface{}, error, bool) {
	fn := func() (interface{}, error) {
		var (
			logger = log.WithCtx(ctx)
			db     = orm.GetDB()
		)

		rows, err := db.QueryContext(ctx, stringSQL, args...)
		if err != nil {
			if errors.Is(err, dbsql.ErrNoRows) {
				return nil, nil
			}

			logger.Error(err.Error())
			return nil, err
		}
		defer rows.Close()

		return s.scanRows(ctx, rows)
	}

	if lockKey == "" {
		v, err := fn()
		return v, err, false
	}

	var (
		sqlKey = GetCustomLockKey(lockKey)
	)

	return s.g.Do(sqlKey, fn)
}

func (s *baseImpl) BulkInsert(ctx context.Context, models interface{}, mp *constant.ModelParams) (m interface{}, err error) {
	var (
		db     = orm.GetDB()
		logger = log.WithCtx(ctx)
		sql    = db.NewInsert().Model(models)
	)

	if !mp.DisableCache {
		sql.Returning("*")
	} else {
		sql.Returning("NULL")
	}

	_, err = sql.Exec(ctx)
	if err != nil {
		logger.Error(err.Error(), zap.String("sql", sql.String()))
		return nil, err
	}

	return models, nil
}

func (s *baseImpl) UpdateByPK(ctx context.Context, sql *bun.UpdateQuery, m interface{}, mp *constant.ModelParams) (r interface{}, err error) {
	var (
		logger = log.WithCtx(ctx)
	)

	sql = sql.Model(m).WherePK().OmitZero().Returning("*")

	if !mp.DisableCache {
		sql.Returning("*")
	} else {
		sql.Returning("NULL")
	}

	_, err = sql.Exec(ctx)
	if err != nil {
		logger.Error(err.Error(), zap.String("sql", sql.String()))
		return nil, err
	}

	if !mp.DisableCache {
		go rd.HSet(commontil.CopyContext(ctx), mp, m)
	} else {
		go rd.HDel(commontil.CopyContext(ctx), mp, m)
	}

	return m, err
}

func (s *baseImpl) BulkUpdateByPK(ctx context.Context, sql *bun.UpdateQuery, models interface{}, mp *constant.ModelParams) (r interface{}, err error) {
	var (
		logger = log.WithCtx(ctx)
	)

	sql = sql.Model(models).WherePK().Bulk()

	if !mp.DisableCache {
		sql.Returning("*")
	} else {
		sql.Returning("NULL")
	}

	_, err = sql.Exec(ctx)
	if err != nil {
		logger.Error(err.Error(), zap.String("sql", sql.String()))
		return nil, err
	}

	if !mp.DisableCache {
		go rd.HMSet(commontil.CopyContext(ctx), mp, models)
	} else {
		go rd.HMDel(commontil.CopyContext(ctx), mp, models)
	}

	return models, err
}
