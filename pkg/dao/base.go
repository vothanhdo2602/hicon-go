package dao

import (
	"context"
	dbsql "database/sql"
	"errors"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/model/entity"
	"github.com/vothanhdo2602/hicon/external/util/commontil"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/external/util/sftil"
	"github.com/vothanhdo2602/hicon/internal/orm"
	"github.com/vothanhdo2602/hicon/internal/rd"
	"go.uber.org/zap"
	"sync"
)

const (
	FindByPKAction         = "find_by_primary"
	FindAllAction          = "find_all_action"
	CustomLockKey          = "custom_lock"
	BulkInsertLockKey      = "bulk_insert"
	UpdateByPKLockKey      = "update_by_pk"
	BulkUpdateByPKLockKey  = "bulk_update_by_pk"
	BulkWriteWithTxLockKey = "bulk_write_with_tx"
)

func GetCustomLockKey(lockKey string) string {
	return fmt.Sprintf("%s:%s", CustomLockKey, lockKey)
}

func GetBulkInsertLockKey(lockKey string) string {
	return fmt.Sprintf("%s:%s", BulkInsertLockKey, lockKey)
}

func GetUpdateByPKLockKey(lockKey string) string {
	return fmt.Sprintf("%s:%s", UpdateByPKLockKey, lockKey)
}

func GetBulkUpdateByPKLockKey(lockKey string) string {
	return fmt.Sprintf("%s:%s", BulkUpdateByPKLockKey, lockKey)
}

func GetBulkWriteWithTxLockKey(lockKey string) string {
	return fmt.Sprintf("%s:%s", BulkWriteWithTxLockKey, lockKey)
}

type dbInterface interface {
	String() string
	Scan(ctx context.Context, values ...any) error
	Model(model any) *bun.SelectQuery
}

type BaseInterface interface {
	FindByPK(ctx context.Context, pk interface{}, id string, mp *entity.ModelParams) (m interface{}, err error, shared bool)
	FindOne(ctx context.Context, sql dbInterface, model interface{}, mp *entity.ModelParams, values ...any) (m interface{}, err error, shared bool)
	FindAll(ctx context.Context, sql dbInterface, models interface{}, mp *entity.ModelParams, values ...any) (interface{}, error, bool)
	Exec(ctx context.Context, tx bun.IDB, stringSQL, lockKey string, args ...any) (interface{}, error, bool)
	BulkInsert(ctx context.Context, db bun.IDB, models interface{}, mp *entity.ModelParams) (m interface{}, err error)
	UpdateByPK(ctx context.Context, sql *bun.UpdateQuery, m interface{}, mp *entity.ModelParams) (r interface{}, err error)
	UpdateAll(ctx context.Context, sql *bun.UpdateQuery, m interface{}, mp *entity.ModelParams) (r interface{}, err error)
	BulkUpdateByPK(ctx context.Context, sql *bun.UpdateQuery, m interface{}, mp *entity.ModelParams) (r interface{}, err error)
	DeleteByPK(ctx context.Context, sql *bun.DeleteQuery, m interface{}, mp *entity.ModelParams) (r interface{}, err error)
}

type baseImpl struct {
	g  sftil.Group
	mu sync.Mutex
}

func (s *baseImpl) FindByPK(ctx context.Context, m interface{}, id string, mp *entity.ModelParams) (interface{}, error, bool) {
	var (
		key = fmt.Sprintf("%s:%s", FindByPKAction, id)
	)

	v, err, shared := s.g.Do(key, func() (interface{}, error) {
		newModel, _ := config.TransformModel(mp.Table, nil, m, mp.ModeType)
		return s.findByPK(ctx, newModel, mp)
	})

	if v != nil && !mp.WhereAllWithDeleted {
		var (
			cols = config.GetModelRegistry().GetTableConfig(mp.Table).SoftDeleteColumns
		)
		for _, c := range cols {
			if !entity.IsZeroValueField(v, c) {
				return nil, err, shared
			}
		}
	}

	return v, err, shared
}

func (s *baseImpl) findByPK(ctx context.Context, m interface{}, mp *entity.ModelParams) (interface{}, error) {
	var (
		logger = log.WithCtx(ctx)
		db     = orm.GetDB()
		sql    = db.NewSelect().Model(m)
	)

	if !mp.DisableCache {
		cachedModel := rd.HGet(ctx, mp, m)
		if cachedModel != nil {
			return cachedModel, nil
		}
	}

	err := sql.WherePK().WhereAllWithDeleted().Scan(ctx)
	if err != nil {
		if errors.Is(err, dbsql.ErrNoRows) {
			return nil, nil
		}

		logger.Error(err.Error())
		return nil, err
	}

	if !mp.DisableCache {
		go rd.HSet(commontil.CopyContext(ctx), mp, m)
	}
	return m, nil
}

func (s *baseImpl) FindAll(ctx context.Context, sql dbInterface, models interface{}, mp *entity.ModelParams, values ...any) (interface{}, error, bool) {
	var (
		sqlStr = sql.String()
		sqlKey = fmt.Sprintf("%s:%s", FindAllAction, sqlStr)
	)

	v, err, shared := s.g.Do(sqlKey, func() (interface{}, error) {
		var (
			logger = log.WithCtx(ctx)
		)

		if !mp.DisableCache {
			cachedModel := rd.HGetAllSQL(ctx, mp, sqlStr, models, s.FindByPK)
			if cachedModel != nil {
				return cachedModel, nil
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
			go rd.HSetAllSQL(commontil.CopyContext(ctx), mp, sqlStr, models)
		}

		return models, nil
	})

	return v, err, shared
}

func (s *baseImpl) FindOne(ctx context.Context, sql dbInterface, m interface{}, mp *entity.ModelParams, values ...any) (interface{}, error, bool) {
	var (
		sqlStr = sql.String()
		sqlKey = fmt.Sprintf("%s:%s", FindAllAction, sqlStr)
	)

	v, err, shared := s.g.Do(sqlKey, func() (interface{}, error) {
		var (
			logger = log.WithCtx(ctx)
		)

		if !mp.DisableCache {
			cachedModel := rd.HGetSQL(ctx, mp, sqlStr, m, s.FindByPK)
			if cachedModel != nil {
				return cachedModel, nil
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
			go rd.HSetSQL(commontil.CopyContext(ctx), mp, sqlStr, m)
		}

		return &m, err
	})

	if err != nil {
		return nil, err, shared
	}

	return v, err, shared
}

func (s *baseImpl) Exec(ctx context.Context, tx bun.IDB, stringSQL, lockKey string, args ...any) (interface{}, error, bool) {
	fn := func() (interface{}, error) {
		var (
			logger = log.WithCtx(ctx)
		)

		if tx == nil {
			tx = orm.GetDB()
		}

		rows, err := tx.QueryContext(ctx, stringSQL, args...)
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

func (s *baseImpl) BulkInsert(ctx context.Context, db bun.IDB, models interface{}, mp *entity.ModelParams) (m interface{}, err error) {
	var (
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

	go rd.HDelRefSQL(ctx, mp)

	return models, nil
}

func (s *baseImpl) UpdateByPK(ctx context.Context, sql *bun.UpdateQuery, m interface{}, mp *entity.ModelParams) (r interface{}, err error) {
	var (
		logger = log.WithCtx(ctx)
		bgCtx  = commontil.CopyContext(ctx)
	)

	sql = sql.Model(m).WherePK().WhereDeleted().OmitZero()

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
		go rd.HSet(bgCtx, mp, m)
	} else {
		go rd.HDel(bgCtx, mp, m)
	}

	return m, err
}

func (s *baseImpl) UpdateAll(ctx context.Context, sql *bun.UpdateQuery, m interface{}, mp *entity.ModelParams) (r interface{}, err error) {
	var (
		logger = log.WithCtx(ctx)
		bgCtx  = commontil.CopyContext(ctx)
	)

	sql = sql.Model(m).WhereDeleted().OmitZero()

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
		go rd.HMSet(bgCtx, mp, m)
	} else {
		go rd.HMDel(bgCtx, mp, m)
	}

	return m, err
}

func (s *baseImpl) DeleteByPK(ctx context.Context, sql *bun.DeleteQuery, m interface{}, mp *entity.ModelParams) (r interface{}, err error) {
	var (
		logger = log.WithCtx(ctx)
		bgCtx  = commontil.CopyContext(ctx)
	)

	sql = sql.Model(m).WherePK()

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
		go rd.HSet(bgCtx, mp, m)
	} else {
		go rd.HDel(bgCtx, mp, m)
	}

	return m, err
}

func (s *baseImpl) BulkUpdateByPK(ctx context.Context, sql *bun.UpdateQuery, models interface{}, mp *entity.ModelParams) (r interface{}, err error) {
	var (
		logger = log.WithCtx(ctx)
		bgCtx  = commontil.CopyContext(ctx)
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
		go rd.HMSet(bgCtx, mp, models)
	} else {
		go rd.HMDel(bgCtx, mp, models)
	}

	return models, err
}
