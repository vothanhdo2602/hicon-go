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
		table: table,
	}
}

func (s *UpdateAll) WithDisableCache() *UpdateAll {
	s.disableCache = true
	return s
}

func (s *UpdateAll) WithLockKey(lockKey string) *UpdateAll {
	s.lockKey = lockKey
	return s
}

func (s *UpdateAll) Set(query string, args ...interface{}) *UpdateAll {
	s.set = append(s.set, &QueryWithArgs{Query: query, Args: args})
	return s
}

func (s *UpdateAll) Where(query string, args ...interface{}) *UpdateAll {
	s.where = append(s.where, &QueryWithArgs{Query: query, Args: args})
	return s
}

func (s *UpdateAll) WhereAllWithDeleted() *UpdateAll {
	s.whereAllWithDeleted = true
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
