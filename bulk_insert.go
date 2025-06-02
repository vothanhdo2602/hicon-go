package hicon

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon-sm/sqlexecutor"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *Client) NewBulkInsert(table string) *BulkInsert {
	return &BulkInsert{
		table: table,
	}
}

func (s *BulkInsert) WithDisableCache() *BulkInsert {
	s.disableCache = true
	return s
}

func (s *BulkInsert) WithLockKey(lockKey string) *BulkInsert {
	s.lockKey = lockKey
	return s
}

func (s *BulkInsert) Data(data []interface{}) *BulkInsert {
	s.data = data
	return s
}

func (s *BulkInsert) Exec(ctx context.Context, opts ExecOptions) (r *sqlexecutor.BaseResponse, err error) {
	headers := map[string]string{}
	if opts.RequestID != "" {
		headers["X-Request-ID"] = opts.RequestID
	}

	reqBytes, err := json.Marshal(&BaseRequest{Body: s, Headers: headers})
	if err != nil {
		return
	}

	return sqlexecutor.NewSQLExecutorClient(client.conn).BulkInsert(ctx, &anypb.Any{Value: reqBytes})
}
