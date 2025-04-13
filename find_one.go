package hicon

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon-go/hicon-sm/sqlexecutor"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *Client) NewFindOne(table string) *FindOne {
	return &FindOne{
		table: table,
	}
}

func (s *FindOne) WithDisableCache() *FindOne {
	s.disableCache = true
	return s
}

func (s *FindOne) Selects(columns ...string) *FindOne {
	s.selects = append(s.selects, columns...)
	return s
}

func (s *FindOne) Where(query string, args ...interface{}) *FindOne {
	s.where = append(s.where, &QueryWithArgs{Query: query, Args: args})
	return s
}

func (s *FindOne) Relation(relation ...string) *FindOne {
	s.relations = append(s.relations, relation...)
	return s
}

func (s *FindOne) Join(query string, args ...interface{}) *FindOne {
	s.joins = append(s.joins, &Join{Join: query, Args: args})
	return s
}

func (s *FindOne) Offset(offset int) *FindOne {
	s.offset = offset
	return s
}

func (s *FindOne) OrderBy(orderBy string) *FindOne {
	s.orderBy = append(s.orderBy, orderBy)
	return s
}

func (s *FindOne) WhereAllWithDeleted() *FindOne {
	s.whereAllWithDeleted = true
	return s
}

func (s *FindOne) Exec(ctx context.Context) (r *sqlexecutor.BaseResponse, err error) {
	reqBytes, err := json.Marshal(&BaseRequest{Body: s})
	if err != nil {
		return
	}

	return sqlexecutor.NewSQLExecutorClient(client.conn).Exec(ctx, &anypb.Any{Value: reqBytes})
}
