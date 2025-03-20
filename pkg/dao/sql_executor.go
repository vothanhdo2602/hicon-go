package dao

import (
	"context"
	dbsql "database/sql"
	"github.com/vothanhdo2602/hicon/external/util/log"
)

var (
	sqlExecutor *sqlExecutorImpl
)

type SQLExecutorInterface interface {
	BaseInterface
}
type sqlExecutorImpl struct {
	baseImpl
}

func SQLExecutor() SQLExecutorInterface {
	if sqlExecutor == nil {
		sqlExecutor = &sqlExecutorImpl{}
	}
	return sqlExecutor
}

func (s *baseImpl) scanRows(ctx context.Context, rows *dbsql.Rows) ([]interface{}, error) {
	var (
		models []interface{}
		logger = log.WithCtx(ctx)
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
			logger.Error(err.Error())
			return nil, err
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
