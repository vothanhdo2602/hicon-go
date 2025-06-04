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
		table: table,
	}
}

func (s *DeleteByPK) WithDisableCache() *DeleteByPK {
	s.disableCache = true
	return s
}

func (s *DeleteByPK) WithLockKey(lockKey string) *DeleteByPK {
	s.lockKey = lockKey
	return s
}

func (s *DeleteByPK) Data(data interface{}) *DeleteByPK {
	s.data = data
	return s
}

func (s *DeleteByPK) Where(query string, args ...interface{}) *DeleteByPK {
	s.where = append(s.where, &QueryWithArgs{Query: query, Args: args})
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
