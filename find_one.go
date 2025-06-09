package hicon

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon-sm/constant"
	"github.com/vothanhdo2602/hicon-sm/sqlexecutor"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *Client) NewFindOne(table string) *FindOne {
	return &FindOne{
		Table: table,
	}
}

func (s *FindOne) Cache(isCache bool) *FindOne {
	s.DisableCache = !isCache
	return s
}

func (s *FindOne) WithSelects(columns ...string) *FindOne {
	s.Selects = append(s.Selects, columns...)
	return s
}

func (s *FindOne) WithWhere(query string, args ...interface{}) *FindOne {
	s.Where = append(s.Where, &QueryWithArgs{Query: query, Args: args})
	return s
}

func (s *FindOne) Relation(relation ...string) *FindOne {
	s.Relations = append(s.Relations, relation...)
	return s
}

func (s *FindOne) Join(query string, args ...interface{}) *FindOne {
	s.Joins = append(s.Joins, &Join{Join: query, Args: args})
	return s
}

func (s *FindOne) WithOffset(offset int) *FindOne {
	s.Offset = offset
	return s
}

func (s *FindOne) WithOrderBy(orderBy string) *FindOne {
	s.OrderBy = append(s.OrderBy, orderBy)
	return s
}

func (s *FindOne) WithWhereAllWithDeleted() *FindOne {
	s.WhereAllWithDeleted = true
	return s
}

func (s *FindOne) Exec(ctx context.Context, opts *ExecOptions) (r *BaseResponse, err error) {
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

	respBytes, err := sqlexecutor.NewSQLExecutorClient(client.conn).Exec(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		return
	}

	result := &BaseResponse{}
	err = json.Unmarshal(respBytes.Value, result)
	return result, err
}
