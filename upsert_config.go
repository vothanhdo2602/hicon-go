package hicon

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon-sm/constant"
	"github.com/vothanhdo2602/hicon-sm/sqlexecutor"
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
func (s *UpsertConfig) Exec(ctx context.Context, opts ExecOptions) (r *BaseResponse, err error) {
	h := map[string]string{}
	if opts.RequestID != "" {
		h[constant.HeaderXRequestId] = opts.RequestID
	}

	reqBytes, err := json.Marshal(&BaseRequest{Body: s, Headers: h})
	if err != nil {
		return
	}

	respBytes, err := sqlexecutor.NewSQLExecutorClient(client.conn).UpsertConfig(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		return
	}

	result := &BaseResponse{}
	err = json.Unmarshal(respBytes.Value, result)
	return result, err
}
