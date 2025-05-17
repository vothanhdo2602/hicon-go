// Package hicon provides a client SDK for optimizing database queries through hicon query proxy.
//
// MIT License - see LICENSE file for details.
package hicon

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon-go/hicon-sm/sqlexecutor"
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

func (s *UpdateAll) Exec(ctx context.Context) (r *sqlexecutor.BaseResponse, err error) {
	reqBytes, err := json.Marshal(&BaseRequest{Body: s})
	if err != nil {
		return
	}

	return sqlexecutor.NewSQLExecutorClient(client.conn).UpdateAll(ctx, &anypb.Any{Value: reqBytes})
}
