// Package hicon provides a client SDK for optimizing database queries.
//
// MIT License - see LICENSE file for details.
package hicon

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon-go/hicon-sm/sqlexecutor"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *Client) NewBulkWriteWithTx(operations ...*Operation) *BulkWriteWithTx {
	return &BulkWriteWithTx{
		operations: operations,
	}
}

func (s *BulkWriteWithTx) WithLockKey(lockKey string) *BulkWriteWithTx {
	s.lockKey = lockKey
	return s
}

func (s *BulkWriteWithTx) Exec(ctx context.Context) (r *sqlexecutor.BaseResponse, err error) {
	reqBytes, err := json.Marshal(&BaseRequest{Body: s})
	if err != nil {
		return
	}

	return sqlexecutor.NewSQLExecutorClient(client.conn).BulkWriteWithTx(ctx, &anypb.Any{Value: reqBytes})
}
