package hicon

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon-sm/constant"
	"github.com/vothanhdo2602/hicon-sm/sqlexecutor"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *Client) NewBulkInsert(table string) *BulkInsert {
	return &BulkInsert{
		Table: table,
	}
}

func (s *BulkInsert) Cache(isCache bool) *BulkInsert {
	s.DisableCache = !isCache
	return s
}

func (s *BulkInsert) WithLockKey(lockKey string) *BulkInsert {
	s.LockKey = lockKey
	return s
}

func (s *BulkInsert) WithData(data []interface{}) *BulkInsert {
	s.Data = data
	return s
}

func (s *BulkInsert) Exec(ctx context.Context, opts *ExecOptions) (r *BaseResponse, err error) {
	headers := map[string]string{}
	if opts != nil {
		if opts.RequestID != "" {
			headers[constant.HeaderXRequestId] = opts.RequestID
		}
	}

	reqBytes, err := json.Marshal(&BaseRequest{Body: s, Headers: headers})
	if err != nil {
		return
	}

	respBytes, err := sqlexecutor.NewSQLExecutorClient(client.conn).BulkInsert(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		return
	}

	result := &BaseResponse{}
	err = json.Unmarshal(respBytes.Value, result)
	return result, err
}
