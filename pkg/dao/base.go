package dao

import (
	"context"
	dbsql "database/sql"
	"errors"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/external/util/sftil"
	"github.com/vothanhdo2602/hicon/internal/orm"
	"go.uber.org/zap"
	"sync"
)

const (
	FindByPrimaryKeysAction = "find_by_primary_keys_action"
	FindAllAction           = "find_all_action"
)

type dbInterface interface {
	String() string
	Scan(ctx context.Context, values ...any) error
	Model(model any) *bun.SelectQuery
}

type BaseInterface interface {
	FindByPrimaryKeys(ctx context.Context, primaryKeys map[string]interface{}, id string, md *ModelParams) (m interface{}, err error, shared bool)
	ScanOne(ctx context.Context, sql dbInterface, model interface{}, md *ModelParams, values ...any) (m interface{}, err error, shared bool)
	//InsertWithTx(ctx context.Context, tx bun.IDB, m *T) error
	//UpdateWithTxById(ctx context.Context, tx bun.IDB, id string, m *T) error
}

type baseImpl struct {
	g  sftil.Group
	mu sync.Mutex
}

type CacheOpts struct {
	SetCache bool
	GetCache bool
}

type ModelParams struct {
	Database     string
	Table        string
	DisableCache bool
	LockKey      string
}

func (s *baseImpl) FindByPrimaryKeys(ctx context.Context, primaryKeys map[string]interface{}, id string, md *ModelParams) (interface{}, error, bool) {
	var (
		key = fmt.Sprintf("%s:%s", FindByPrimaryKeysAction, id)
	)

	v, err, shared := s.g.Do(key, func() (interface{}, error) {
		var (
			m = config.GetModelRegistry().GetNewModel(md.Table)
		)
		return s.FindByPrimaryKeysWithCacheOpts(ctx, primaryKeys, m, md)
	})

	return v, err, shared
}

func (s *baseImpl) FindByPrimaryKeysWithCacheOpts(ctx context.Context, primaryKeys map[string]interface{}, m interface{}, md *ModelParams) (interface{}, error) {
	var (
		logger = log.WithCtx(ctx)
		db     = orm.GetDB()
		sql    = db.NewSelect().Model(m)
	)

	if !md.DisableCache {
		//cacheModel := rd.HGet[T](ctx, id)
		//if cacheModel != nil {
		//	if opts.SetCache {
		//		go rd.HSet(bgCtx, *cacheModel)
		//	}
		//
		//	return cacheModel, nil
		//}
	}

	for k, v := range primaryKeys {
		sql.Where("? = ?", bun.Ident(k), v)
	}

	err := sql.Scan(ctx)
	if err != nil {
		if errors.Is(err, dbsql.ErrNoRows) {
			return nil, nil
		}

		logger.Error(err.Error())
		return nil, err
	}

	if !md.DisableCache {
		//go rd.HSet(bgCtx, m)
	}
	return &m, nil
}

func (s *baseImpl) Scan(ctx context.Context, sql dbInterface, md *ModelParams, values ...any) (interface{}, error, bool) {
	var (
		sqlKey = fmt.Sprintf("%s:%s", FindAllAction, sql.String())
	)

	v, err, shared := s.g.Do(sqlKey, func() (interface{}, error) {
		var (
			logger = log.WithCtx(ctx)
			models = []interface{}{config.GetModelRegistry().GetNewModel(md.Table)}
		)

		err := sql.Model(&models).Scan(ctx, values...)
		if err != nil {
			if errors.Is(err, dbsql.ErrNoRows) {
				return nil, nil
			}

			logger.Error(err.Error())
			return nil, err
		}

		return models, nil
	})

	if err != nil {
		return nil, err, shared
	}

	return v, err, shared
}

func (s *baseImpl) ScanOne(ctx context.Context, sql dbInterface, m interface{}, md *ModelParams, values ...any) (interface{}, error, bool) {
	var (
		sqlKey = fmt.Sprintf("%s:%s", FindAllAction, sql.String())
	)

	v, err, shared := s.g.Do(sqlKey, func() (interface{}, error) {
		var (
			logger = log.WithCtx(ctx)
		)

		err := sql.Scan(ctx, values...)
		if err != nil {
			if errors.Is(err, dbsql.ErrNoRows) {
				return nil, nil
			}

			logger.Error(err.Error())
			return nil, err
		}

		if !md.DisableCache {
			//go rd.HSet(commontil.CopyContext(ctx), m)
		}

		return &m, err
	})

	if err != nil {
		return nil, err, shared
	}

	return v, err, shared
}

func (s *baseImpl) insert(ctx context.Context, model interface{}) (r dbsql.Result, err error) {
	var (
		db        = orm.GetDB()
		logger    = log.WithCtx(ctx)
		insertSQL = db.NewInsert().Model(model).Returning("*")
	)

	r, err = insertSQL.Exec(ctx)
	if err != nil {
		logger.Error(err.Error(), zap.String("sql", insertSQL.String()))
		return
	}

	//go rd.HSet(commontil.CopyContext(ctx), *model)

	return r, err
}

func (s *baseImpl) insertAll(ctx context.Context, models []interface{}) (m []interface{}, r dbsql.Result, err error) {
	if len(models) == 0 {
		return
	}

	var (
		db        = orm.GetDB()
		logger    = log.WithCtx(ctx)
		insertSQL = db.NewInsert().Model(&models).Returning("*")
	)

	r, err = insertSQL.Exec(ctx)
	if err != nil {
		logger.Error(err.Error(), zap.String("sql", insertSQL.String()))
		return
	}

	//go rd.HMSet(commontil.CopyContext(ctx), models)

	return models, r, err
}

//func (s *baseImpl) InsertWithTx(ctx context.Context, tx bun.IDB, m interface{}) error {
//	var (
//		logger = log.WithCtx(ctx)
//	)
//
//	_, err := tx.NewInsert().Model(m).Returning("*").Exec(ctx)
//	if err != nil {
//		logger.Error(err.Error(), zap.Any("model", m))
//		return err
//	}
//	//go rd.HSet(commontil.CopyContext(ctx), m)
//	return err
//}

//func (s *baseImpl) UpdateWithTxById(ctx context.Context, tx bun.IDB, id string, m interface{}) error {
//	var (
//		logger = log.WithCtx(ctx)
//	)
//
//	_, err := tx.NewUpdate().Model(m).Where(whereIDTmpl, id).Returning("*").Exec(ctx)
//	if err != nil {
//		logger.Error(err.Error(), zap.Any("model", m))
//		return err
//	}
//	//go rd.HSet(commontil.CopyContext(ctx), *m)
//	return err
//}
