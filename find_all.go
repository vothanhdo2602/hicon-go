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

func (s *Client) NewFindAll(table string) *FindAll {
	return &FindAll{
		table: table,
	}
}

func (s *FindAll) WithDisableCache() *FindAll {
	s.disableCache = true
	return s
}

func (s *FindAll) Selects(columns ...string) *FindAll {
	s.selects = append(s.selects, columns...)
	return s
}

func (s *FindAll) Where(query string, args ...interface{}) *FindAll {
	s.where = append(s.where, &QueryWithArgs{Query: query, Args: args})
	return s
}

func (s *FindAll) Relation(relation ...string) *FindAll {
	s.relations = append(s.relations, relation...)
	return s
}

func (s *FindAll) Join(query string, args ...interface{}) *FindAll {
	s.joins = append(s.joins, &Join{Join: query, Args: args})
	return s
}

func (s *FindAll) Limit(limit int) *FindAll {
	s.limit = limit
	return s
}

func (s *FindAll) Offset(offset int) *FindAll {
	s.offset = offset
	return s
}

func (s *FindAll) OrderBy(orderBy string) *FindAll {
	s.orderBy = append(s.orderBy, orderBy)
	return s
}

func (s *FindAll) WhereAllWithDeleted() *FindAll {
	s.whereAllWithDeleted = true
	return s
}

func (s *FindAll) Exec(ctx context.Context) (r *sqlexecutor.BaseResponse, err error) {
	reqBytes, err := json.Marshal(&BaseRequest{Body: s})
	if err != nil {
		return
	}

	return sqlexecutor.NewSQLExecutorClient(client.conn).FindAll(ctx, &anypb.Any{Value: reqBytes})
}
