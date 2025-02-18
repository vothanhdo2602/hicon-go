package dao

import (
	"context"
	dbsql "database/sql"
	"errors"
	"fmt"
	"github.com/vothanhdo2602/hicon/external/util/sftil"
	"strings"
	"sync"

	"github.com/uptrace/bun"
	"github.com/vothanhdo2602/hicon/external/model/entity"
	"github.com/vothanhdo2602/hicon/external/util/commontil"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/internal/orm"
	"github.com/vothanhdo2602/hicon/internal/rd"
	"go.uber.org/zap"
)

const (
	findByIDAction  = "find-by-id-action"
	findByIDsAction = "find_by_ids_action"
	findAllAction   = "find_all_action"
)

const (
	whereIDTmpl = "id = ?"
)

func FindByIDPrefixKey(dbName string) string {
	return fmt.Sprintf("%s:%s", dbName, findByIDAction)
}

func FindByIDsPrefixKey(dbName string) string {
	return fmt.Sprintf("%s:%s", dbName, findByIDsAction)
}

func FindAllPrefixKey(dbName string) string {
	return fmt.Sprintf("%s:%s", dbName, findAllAction)
}

type dbInterface interface {
	String() string
	Scan(ctx context.Context, values ...any) error
	Model(model any) *bun.SelectQuery
}

type BaseInterface[T any] interface {
	FindByID(ctx context.Context, id string) *T
	FindByIDs(ctx context.Context, ids []string) []T
	FindAllByIDs(ctx context.Context, ids []string) []T
	InsertWithTx(ctx context.Context, tx bun.IDB, m *T) error
	UpdateWithTxById(ctx context.Context, tx bun.IDB, id string, m *T) error
}

type baseImpl[T entity.BaseInterface[T]] struct {
	g  sftil.Group
	mu sync.Mutex
}

type baseCacheOpts struct {
	SetCache bool
	GetCache bool
}

func (s *baseImpl[T]) FindByID(ctx context.Context, id string) *T {
	var (
		key = fmt.Sprintf("%s:%s", findByIDPrefixKey, id)
	)

	if id == "" {
		return nil
	}

	v, err, _ := s.g.Do(key, func() (interface{}, error) {
		return s.FindByIDWithCacheOpts(ctx, id, baseCacheOpts{SetCache: true, GetCache: true})
	})

	if err != nil {
		return nil
	}

	return v.(*T)
}

func (s *baseImpl[T]) FindByIDWithCacheOpts(ctx context.Context, id string, opts baseCacheOpts) (*T, error) {
	var (
		logger = log.WithCtx(ctx)
		m      = commontil.InitStructFromGeneric[T]()
		bgCtx  = commontil.CopyContext(ctx)
		db     = orm.GetDB()
	)

	if opts.GetCache {
		cacheModel := rd.HGet[T](ctx, id)
		if cacheModel != nil {
			if opts.SetCache {
				//go rd.HSet(bgCtx, *cacheModel)
			}

			return cacheModel, nil
		}
	}

	sql := db.NewSelect().Model(&m).Where(whereIDTmpl, id)
	err := sql.Scan(ctx)
	if err != nil {
		if errors.Is(err, dbsql.ErrNoRows) {
			// logical error, for production tracking
			logger.Warn(err.Error(), zap.String("sql", sql.String()))
		} else {
			logger.Error(err.Error())
		}
		return nil, err
	}

	if opts.SetCache {
		//go rd.HSet(bgCtx, m)
	}
	return &m, nil
}

func (s *baseImpl[T]) FindByIDs(ctx context.Context, ids []string) []T {
	var (
		key = fmt.Sprintf("%s:%s", findByIDsPrefixKey, strings.Join(ids, ","))
	)
	v, _, _ := s.g.Do(key, func() (interface{}, error) {
		var (
			models []T
			wg     sync.WaitGroup
		)

		for i := range ids {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				m := s.FindByID(ctx, ids[i])
				if m != nil {
					s.mu.Lock()
					models = append(models, *m)
					s.mu.Unlock()
				}
			}(i)
		}
		wg.Wait()

		return models, nil
	})

	return v.([]T)
}

func (s *baseImpl[T]) FindAllByIDs(ctx context.Context, ids []string) []T {
	var (
		key = fmt.Sprintf("%s:%s", findByIDsPrefixKey, strings.Join(ids, ","))
	)
	v, _, _ := s.g.Do(key, func() (interface{}, error) {
		var (
			models = make([]T, len(ids))
			wg     sync.WaitGroup
		)

		for i := range ids {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				m := s.FindByID(ctx, ids[i])
				if m != nil {
					models[i] = *m
				}
			}(i)
		}
		wg.Wait()

		return models, nil
	})

	return v.([]T)
}

func (s *baseImpl[T]) scan(ctx context.Context, sql dbInterface, values ...any) []T {
	var (
		sqlKey = fmt.Sprintf("%s:%s", findAllPrefixKey, sql.String())
	)
	models, _, _ := s.g.Do(sqlKey, func() (interface{}, error) {
		var (
			logger = log.WithCtx(ctx)
			models = make([]T, 0)
		)

		err := sql.Model(&models).Scan(ctx, values...)
		if err != nil {
			logger.Error(err.Error(), zap.String("sql", sqlKey))
		}

		return models, nil
	})

	return models.([]T)
}

func (s *baseImpl[T]) scanOne(ctx context.Context, sql dbInterface, values ...any) *T {
	var (
		sqlKey = fmt.Sprintf("%s:%s", findAllPrefixKey, sql.String())
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

	return v.(*T)
}

func (s *baseImpl[T]) insert(ctx context.Context, model *T) (r dbsql.Result, err error) {
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

func (s *baseImpl[T]) insertAll(ctx context.Context, models []T) (m []T, r dbsql.Result, err error) {
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

	go rd.HMSet(commontil.CopyContext(ctx), models)

	return models, r, err
}

func (s *baseImpl[T]) InsertWithTx(ctx context.Context, tx bun.IDB, m *T) error {
	var (
		logger = log.WithCtx(ctx)
	)

	_, err := tx.NewInsert().Model(m).Returning("*").Exec(ctx)
	if err != nil {
		logger.Error(err.Error(), zap.Any("model", m))
		return err
	}
	//go rd.HSet(commontil.CopyContext(ctx), m)
	return err
}

func (s *baseImpl[T]) UpdateWithTxById(ctx context.Context, tx bun.IDB, id string, m *T) error {
	var (
		logger = log.WithCtx(ctx)
	)

	_, err := tx.NewUpdate().Model(m).Where(whereIDTmpl, id).Returning("*").Exec(ctx)
	if err != nil {
		logger.Error(err.Error(), zap.Any("model", m))
		return err
	}
	//go rd.HSet(commontil.CopyContext(ctx), *m)
	return err
}
