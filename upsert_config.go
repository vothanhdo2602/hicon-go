package hicon

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon-go/hicon-sm/sqlexecutor"
	"google.golang.org/protobuf/types/known/anypb"
)

// UpsertConfigOption is a functional configuration option type
type UpsertConfigOption func(config *UpsertConfig)

// WithDebug enables or disables debug mode
func WithDebug(debug bool) UpsertConfigOption {
	return func(c *UpsertConfig) {
		c.Debug = debug
	}
}

// WithDisableCache toggles cache functionality
func WithDisableCache(disable bool) UpsertConfigOption {
	return func(c *UpsertConfig) {
		c.DisableCache = disable
	}
}

// WithDBConfig WithDB Database configuration methods
func WithDBConfig(dbConfig *DBConfig) UpsertConfigOption {
	return func(c *UpsertConfig) {
		c.DBConfig = dbConfig
	}
}

// WithRedis Redis configuration methods
func WithRedis(redis *Redis) UpsertConfigOption {
	return func(c *UpsertConfig) {
		c.Redis = redis
	}
}

// WithTable adds a table configuration
func WithTable(tbl *TableConfig) UpsertConfigOption {
	return func(c *UpsertConfig) {
		c.TableConfigs = append(c.TableConfigs, tbl)
	}
}

// Build finalizes the configuration
func (s *UpsertConfig) build(opts ...UpsertConfigOption) *UpsertConfig {
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Exec finalize the configuration
func (s *UpsertConfig) Exec(ctx context.Context) (r *sqlexecutor.BaseResponse, err error) {
	reqBytes, err := json.Marshal(&BaseRequest{Body: s})
	if err != nil {
		return
	}

	return sqlexecutor.NewSQLExecutorClient(client.conn).UpsertConfig(ctx, &anypb.Any{Value: reqBytes})
}
