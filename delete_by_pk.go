package hicon

import (
	"context"
	"github.com/goccy/go-json"
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

func (s *DeleteByPK) Exec(ctx context.Context) (r *sqlexecutor.BaseResponse, err error) {
	reqBytes, err := json.Marshal(&BaseRequest{Body: s})
	if err != nil {
		return
	}

	return sqlexecutor.NewSQLExecutorClient(client.conn).DeleteByPK(ctx, &anypb.Any{Value: reqBytes})
}
