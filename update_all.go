package hicon

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon-sm/constant"
	"github.com/vothanhdo2602/hicon-sm/sqlexecutor"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *Client) NewUpdateAll(table string) *UpdateAll {
	return &UpdateAll{
		Table: table,
	}
}

func (s *UpdateAll) Cache(isCache bool) *UpdateAll {
	s.DisableCache = !isCache
	return s
}

func (s *UpdateAll) WithLockKey(lockKey string) *UpdateAll {
	s.LockKey = lockKey
	return s
}

func (s *UpdateAll) WithSet(query string, args ...interface{}) *UpdateAll {
	s.Set = append(s.Set, &QueryWithArgs{Query: query, Args: args})
	return s
}

func (s *UpdateAll) WithWhere(query string, args ...interface{}) *UpdateAll {
	s.Where = append(s.Where, &QueryWithArgs{Query: query, Args: args})
	return s
}

func (s *UpdateAll) WithWhereAllWithDeleted() *UpdateAll {
	s.WhereAllWithDeleted = true
	return s
}

func (s *UpdateAll) Exec(ctx context.Context, opts *ExecOptions) (r *BaseResponse, err error) {
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

	respBytes, err := sqlexecutor.NewSQLExecutorClient(client.conn).UpdateAll(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		return
	}

	result := &BaseResponse{}
	err = json.Unmarshal(respBytes.Value, result)
	return result, err
}
