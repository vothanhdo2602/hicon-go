package hicon

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon-sm/constant"
	"github.com/vothanhdo2602/hicon-sm/sqlexecutor"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *Client) NewUpdateByPK(table string) *UpdateByPK {
	return &UpdateByPK{
		Table: table,
	}
}

func (s *UpdateByPK) Cache(isCache bool) *UpdateByPK {
	s.DisableCache = !isCache
	return s
}

func (s *UpdateByPK) WithLockKey(lockKey string) *UpdateByPK {
	s.LockKey = lockKey
	return s
}

func (s *UpdateByPK) WithData(data interface{}) *UpdateByPK {
	s.Data = data
	return s
}

func (s *UpdateByPK) WithWhere(query string, args ...interface{}) *UpdateByPK {
	s.Where = append(s.Where, &QueryWithArgs{Query: query, Args: args})
	return s
}

func (s *UpdateByPK) Exec(ctx context.Context, opts *ExecOptions) (r *BaseResponse, err error) {
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

	respBytes, err := sqlexecutor.NewSQLExecutorClient(client.conn).UpdateByPK(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		return
	}

	result := &BaseResponse{}
	err = json.Unmarshal(respBytes.Value, result)
	return result, err
}
