package hicon

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon-sm/constant"
	"github.com/vothanhdo2602/hicon-sm/sqlexecutor"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *Client) NewDeleteByPK(table string) *DeleteByPK {
	return &DeleteByPK{
		Table: table,
	}
}

func (s *DeleteByPK) Cache(isCache bool) *DeleteByPK {
	s.DisableCache = !isCache
	return s
}

func (s *DeleteByPK) WithLockKey(lockKey string) *DeleteByPK {
	s.LockKey = lockKey
	return s
}

func (s *DeleteByPK) WithForceDelete(forceDelete bool) *DeleteByPK {
	s.ForceDelete = forceDelete
	return s
}

func (s *DeleteByPK) WithData(data interface{}) *DeleteByPK {
	s.Data = data
	return s
}

func (s *DeleteByPK) WithWhere(query string, args ...interface{}) *DeleteByPK {
	s.Where = append(s.Where, &QueryWithArgs{Query: query, Args: args})
	return s
}

func (s *DeleteByPK) Exec(ctx context.Context, opts *ExecOptions) (r *BaseResponse, err error) {
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

	respBytes, err := sqlexecutor.NewSQLExecutorClient(client.conn).DeleteByPK(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		return
	}

	result := &BaseResponse{}
	err = json.Unmarshal(respBytes.Value, result)
	return result, err
}
