package dao

import (
	"context"
	dbsql "database/sql"
	"errors"
	"fmt"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/util/sftil"
	"sync"

	"github.com/uptrace/bun"
	"github.com/vothanhdo2602/hicon/external/util/commontil"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/internal/orm"
	"go.uber.org/zap"
)

const (
	FindByPrimaryKeysAction = "find_by_primary_keys_action"
	FindAllAction           = "find_all_action"
)

const (
	whereIDTmpl = "id = ?"
)

type dbInterface interface {
	String() string
	Scan(ctx context.Context, values ...any) error
	Model(model any) *bun.SelectQuery
}

type BaseInterface[T any] interface {
	FindByPrimaryKeys(ctx context.Context, id string, md *ModelParams) (m *T, err error, shared bool)
	//InsertWithTx(ctx context.Context, tx bun.IDB, m *T) error
	//UpdateWithTxById(ctx context.Context, tx bun.IDB, id string, m *T) error
}

type baseImpl[T any] struct {
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

func (s *baseImpl[T]) FindByPrimaryKeys(ctx context.Context, id string, md *ModelParams) (m *T, err error, shared bool) {
	var (
		key = fmt.Sprintf("%s:%s", FindByPrimaryKeysAction, id)
	)

	chResult := s.g.DoChan(key, func() (interface{}, error) {
		return s.FindByPrimaryKeysWithCacheOpts(ctx, id, md)
	})

	result := <-chResult

	if result.Err != nil {
		return nil, result.Err, shared
	}
	if result.Val == nil {
		return
	}

	return result.Val.(*T), err, result.Shared
}

func (s *baseImpl[T]) FindByPrimaryKeysWithCacheOpts(ctx context.Context, id string, md *ModelParams) (interface{}, error) {
	var (
		logger = log.WithCtx(ctx)
		db     = orm.GetDB()
		m      = config.GetModelRegistry().GetNewModel(md.Table)
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

	sql := db.NewSelect().Model(m).Where(whereIDTmpl, id)
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

func (s *baseImpl[T]) scan(ctx context.Context, sql dbInterface, values ...any) []interface{} {
	var (
		sqlKey = fmt.Sprintf("%s:%s", FindAllAction, sql.String())
	)
	models, _, _ := s.g.Do(sqlKey, func() (interface{}, error) {
		var (
			logger = log.WithCtx(ctx)
			models = make([]interface{}, 0)
		)

		err := sql.Model(&models).Scan(ctx, values...)
		if err != nil {
			logger.Error(err.Error(), zap.String("sql", sqlKey))
		}

		return models, nil
	})

	return models.([]interface{})
}

func (s *baseImpl[T]) scanOne(ctx context.Context, sql dbInterface, values ...any) interface{} {
	var (
		sqlKey = fmt.Sprintf("%s:%s", FindAllAction, sql.String())
		logger = log.WithCtx(ctx)
	)

	v, err, _ := s.g.Do(sqlKey, func() (interface{}, error) {
		var (
			m = commontil.InitStructFromGeneric[T]()
		)

		err := sql.Model(&m).Scan(ctx, values...)
		if err != nil {
			if !errors.Is(err, dbsql.ErrNoRows) {
				logger.Error("Resource not found", zap.String("sql", sql.String()))
			}
			return nil, err
		}

		//go rd.HSet(commontil.CopyContext(ctx), m)

		return &m, err
	})

	if err != nil {
		return nil
	}

	return v.(interface{})
}

func (s *baseImpl[T]) insert(ctx context.Context, model interface{}) (r dbsql.Result, err error) {
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

func (s *baseImpl[T]) insertAll(ctx context.Context, models []interface{}) (m []interface{}, r dbsql.Result, err error) {
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

//func (s *baseImpl[T]) InsertWithTx(ctx context.Context, tx bun.IDB, m interface{}) error {
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

//func (s *baseImpl[T]) UpdateWithTxById(ctx context.Context, tx bun.IDB, id string, m interface{}) error {
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
