package service

import (
	"context"
	"fmt"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/util/pstring"
	"github.com/vothanhdo2602/hicon/pkg/dao"
	"strings"
)

type SQLExecutorInterface[T any] interface {
	FindByPrimaryKeys(ctx context.Context, req *requestmodel.FindByPrimaryKeys) (m *T, err error, shared bool)
}

type sqlExecuteImpl[T any] struct {
	dao dao.SQLExecutorInterface[T]
}

func SQLExecutor[T any]() SQLExecutorInterface[T] {
	return &sqlExecuteImpl[T]{
		dao: dao.SQLExecutor[T](),
	}
}

func (s *sqlExecuteImpl[T]) FindByPrimaryKeys(ctx context.Context, req *requestmodel.FindByPrimaryKeys) (m *T, err error, shared bool) {
	var (
		d     = dao.SQLExecutor[T]()
		dbCfg = config.GetENV().DB.DBConfiguration
		md    = &dao.ModelParams{
			Database: dbCfg.Database,
			Table:    req.Table,
		}
		tableRegistry = dbCfg.ModelRegistry.TableConfigurations[md.Table]
		arrKeys       []string
	)

	if dbCfg.DisableCache || req.DisableCache {
		md.DisableCache = req.DisableCache
	}

	for k := range tableRegistry.PrimaryColumns {
		v, ok := req.PrimaryKeys[k]
		if !ok {
			return nil, fmt.Errorf("primary key column %v not found in registered model, please check func update configuration", k), shared
		}
		arrKeys = append(arrKeys, pstring.InterfaceToString(v))
	}

	id := strings.Join(arrKeys, ";")

	//for i := 0; i < 1000; i++ {
	//	go d.FindByPrimaryKeys(ctx, id, md)
	//}
	return d.FindByPrimaryKeys(ctx, id, md)
}
