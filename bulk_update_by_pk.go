package hicon

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon-go/hicon-sm/sqlexecutor"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *Client) NewBulkUpdateByPK(table string) *BulkUpdateByPK {
	return &BulkUpdateByPK{
		table: table,
	}
}

func (s *BulkUpdateByPK) WithDisableCache() *BulkUpdateByPK {
	s.disableCache = true
	return s
}

func (s *BulkUpdateByPK) WithLockKey(lockKey string) *BulkUpdateByPK {
	s.lockKey = lockKey
	return s
}

func (s *BulkUpdateByPK) Set(columns ...string) *BulkUpdateByPK {
	s.set = append(s.set, columns...)
	return s
}

func (s *BulkUpdateByPK) Where(columns ...string) *BulkUpdateByPK {
	s.where = append(s.where, columns...)
	return s
}

func (s *BulkUpdateByPK) Data(data []interface{}) *BulkUpdateByPK {
	s.data = data
	return s
}

func (s *BulkUpdateByPK) Exec(ctx context.Context) (r *sqlexecutor.BaseResponse, err error) {
	reqBytes, err := json.Marshal(&BaseRequest{Body: s})
	if err != nil {
		return
	}

	return sqlexecutor.NewSQLExecutorClient(client.conn).BulkUpdateByPK(ctx, &anypb.Any{Value: reqBytes})
}
