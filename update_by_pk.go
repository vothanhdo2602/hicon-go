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
		table: table,
	}
}

func (s *UpdateByPK) WithDisableCache() *UpdateByPK {
	s.disableCache = true
	return s
}

func (s *UpdateByPK) WithLockKey(lockKey string) *UpdateByPK {
	s.lockKey = lockKey
	return s
}

func (s *UpdateByPK) Data(data interface{}) *UpdateByPK {
	s.data = data
	return s
}

func (s *UpdateByPK) Where(query string, args ...interface{}) *UpdateByPK {
	s.where = append(s.where, &QueryWithArgs{Query: query, Args: args})
	return s
}

func (s *UpdateByPK) Exec(ctx context.Context, opts ExecOptions) (r *BaseResponse, err error) {
	headers := map[string]string{}
	if opts.RequestID != "" {
		headers[constant.HeaderXRequestId] = opts.RequestID
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
