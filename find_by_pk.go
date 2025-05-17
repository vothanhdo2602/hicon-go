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

func (s *Client) NewFindByPK(table string) *FindByPK {
	return &FindByPK{
		table: table,
	}
}

func (s *FindByPK) WithDisableCache() *FindByPK {
	s.disableCache = true
	return s
}

func (s *FindByPK) Selects(columns ...string) *FindByPK {
	s.selects = append(s.selects, columns...)
	return s
}

func (s *FindByPK) Data(data interface{}) *FindByPK {
	s.data = data
	return s
}

func (s *FindByPK) WhereAllWithDeleted() *FindByPK {
	s.whereAllWithDeleted = true
	return s
}

func (s *FindByPK) Exec(ctx context.Context) (r *sqlexecutor.BaseResponse, err error) {
	reqBytes, err := json.Marshal(&BaseRequest{Body: s})
	if err != nil {
		return
	}

	return sqlexecutor.NewSQLExecutorClient(client.conn).FindByPK(ctx, &anypb.Any{Value: reqBytes})
}
