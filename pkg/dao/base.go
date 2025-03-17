package dao

import (
	"context"
	dbsql "database/sql"
	"errors"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/util/commontil"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/external/util/sftil"
	"github.com/vothanhdo2602/hicon/internal/orm"
	"github.com/vothanhdo2602/hicon/internal/rd"
	"go.uber.org/zap"
	"sync"
)

const (
	FindByPrimaryKeysAction = "find_by_primary_keys_action"
	FindAllAction           = "find_all_action"
	CustomLockKey           = "custom_lock_key_action"
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
	FindByPrimaryKeys(ctx context.Context, primaryKeys map[string]interface{}, id string, md *ModelParams) (m interface{}, err error, shared bool)
	ScanOne(ctx context.Context, sql dbInterface, model interface{}, md *ModelParams, values ...any) (m interface{}, err error, shared bool)
	Scan(ctx context.Context, sql dbInterface, models interface{}, md *ModelParams, values ...any) (interface{}, error, bool)
	Exec(ctx context.Context, stringSQL, lockKey string, md *ModelParams, args ...any) (interface{}, error, bool)
	BulkInsert(ctx context.Context, models interface{}, md *ModelParams) (m interface{}, err error)
	UpdateByPrimaryKeys(ctx context.Context, m interface{}, md *ModelParams) (r interface{}, err error)
	//InsertWithTx(ctx context.Context, tx bun.IDB, m *T) error
	//UpdateWithTxById(ctx context.Context, tx bun.IDB, id string, m *T) error
}

type baseImpl struct {
	g  sftil.Group
	mu sync.Mutex
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
		newModel, _ := config.TransformModel(md.Table, nil, primaryKeys, false)
		return s.findByPrimaryKeys(ctx, newModel, md)
	})

	return v, err, shared
}

func (s *baseImpl) findByPrimaryKeys(ctx context.Context, m interface{}, md *ModelParams) (interface{}, error) {
	var (
		logger = log.WithCtx(ctx)
		db     = orm.GetDB()
		sql    = db.NewSelect().Model(m)
	)

	if !md.DisableCache {
		cacheModel := rd.HGet(ctx, md.Database, md.Table, m)
		if cacheModel != nil {
			go rd.HSet(commontil.CopyContext(ctx), md.Database, md.Table, cacheModel)

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

	if !md.DisableCache {
		//go rd.HSet(bgCtx, m)
	}
	return &m, nil
}

func (s *baseImpl) Scan(ctx context.Context, sql dbInterface, models interface{}, md *ModelParams, values ...any) (interface{}, error, bool) {
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

func (s *baseImpl) Exec(ctx context.Context, stringSQL, lockKey string, md *ModelParams, args ...any) (interface{}, error, bool) {
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

		return s.scanRows(rows)
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

func (s *baseImpl) BulkInsert(ctx context.Context, models interface{}, md *ModelParams) (m interface{}, err error) {
	var (
		db     = orm.GetDB()
		logger = log.WithCtx(ctx)
		sql    = db.NewInsert().Model(models)
	)

	if !md.DisableCache {
		sql.Returning("*")
	} else {
		sql.Returning("NULL")
	}

	_, err = sql.Exec(ctx)
	if err != nil {
		logger.Error(err.Error(), zap.String("sql", sql.String()))
		return nil, err
	}

	//go rd.HMSet(commontil.CopyContext(ctx), models)
	return models, nil
}

func (s *baseImpl) UpdateByPrimaryKeys(ctx context.Context, m interface{}, md *ModelParams) (r interface{}, err error) {
	var (
		db     = orm.GetDB()
		logger = log.WithCtx(ctx)
		sql    = db.NewUpdate().Model(m).WherePK().OmitZero()
	)

	if !md.DisableCache {
		sql.Returning("*")
	} else {
		sql.Returning("NULL")
	}

	_, err = sql.Exec(ctx)
	if err != nil {
		logger.Error(err.Error(), zap.String("sql", sql.String()))
		return nil, err
	}

	//go rd.HSet(commontil.CopyContext(ctx), *m)
	return m, err
}

func (s *baseImpl) scanRows(rows *dbsql.Rows) ([]interface{}, error) {
	var (
		models []interface{}
	)

	columns, err := rows.Columns()
	if err != nil {
		return models, err
	}

	values := make([]interface{}, len(columns))
	// Create a slice of pointers to those interfaces
	scanArgs := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		var (
			m = map[string]interface{}{}
		)

		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err)
		}

		for i, col := range columns {
			val := values[i]

			// Handle null values
			if val == nil {
				models = append(models, nil)
				continue
			}

			// The actual type will depend on your database driver
			// Common types like string, int64, float64, bool, []byte will be returned
			switch v := val.(type) {
			case []byte:
				// Convert []byte to string for readability
				m[col] = string(v)
			default:
				m[col] = v
			}
		}

		models = append(models, m)
	}

	return models, nil
}
