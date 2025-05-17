// Package hicon provides a client SDK for optimizing database queries through hicon query proxy.
//
// MIT License - see LICENSE file for details.
package hicon

import (
	"context"
	"github.com/goccy/go-json"
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

func (s *UpdateByPK) Exec(ctx context.Context) (r *sqlexecutor.BaseResponse, err error) {
	reqBytes, err := json.Marshal(&BaseRequest{Body: s})
	if err != nil {
		return
	}

	return sqlexecutor.NewSQLExecutorClient(client.conn).UpdateByPK(ctx, &anypb.Any{Value: reqBytes})
}
