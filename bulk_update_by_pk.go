package hicon

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon-sm/constant"
	"github.com/vothanhdo2602/hicon-sm/sqlexecutor"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *Client) NewBulkUpdateByPK(table string) *BulkUpdateByPK {
	return &BulkUpdateByPK{
		Table: table,
	}
}

func (s *BulkUpdateByPK) Cache(isCache bool) *BulkUpdateByPK {
	s.DisableCache = !isCache
	return s
}

func (s *BulkUpdateByPK) WithLockKey(lockKey string) *BulkUpdateByPK {
	s.LockKey = lockKey
	return s
}

func (s *BulkUpdateByPK) WithSet(columns ...string) *BulkUpdateByPK {
	s.Set = append(s.Set, columns...)
	return s
}

func (s *BulkUpdateByPK) WithWhere(columns ...string) *BulkUpdateByPK {
	s.Where = append(s.Where, columns...)
	return s
}

func (s *BulkUpdateByPK) WithData(data []interface{}) *BulkUpdateByPK {
	s.Data = data
	return s
}

func (s *BulkUpdateByPK) Exec(ctx context.Context, opts *ExecOptions) (r *BaseResponse, err error) {
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

	respBytes, err := sqlexecutor.NewSQLExecutorClient(client.conn).BulkUpdateByPK(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		return
	}

	result := &BaseResponse{}
	err = json.Unmarshal(respBytes.Value, result)
	return result, err
}
