package hicon

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon-sm/constant"
	"github.com/vothanhdo2602/hicon-sm/sqlexecutor"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *Client) NewFindAll(table string) *FindAll {
	return &FindAll{
		Table: table,
	}
}

func (s *FindAll) Cache(isCache bool) *FindAll {
	s.DisableCache = !isCache
	return s
}

func (s *FindAll) WithSelects(columns ...string) *FindAll {
	s.Selects = append(s.Selects, columns...)
	return s
}

func (s *FindAll) WithWhere(query string, args ...interface{}) *FindAll {
	s.Where = append(s.Where, &QueryWithArgs{Query: query, Args: args})
	return s
}

func (s *FindAll) Relation(relation ...string) *FindAll {
	s.Relations = append(s.Relations, relation...)
	return s
}

func (s *FindAll) Join(query string, args ...interface{}) *FindAll {
	s.Joins = append(s.Joins, &Join{Join: query, Args: args})
	return s
}

func (s *FindAll) WithLimit(limit int) *FindAll {
	s.Limit = limit
	return s
}

func (s *FindAll) WithOffset(offset int) *FindAll {
	s.Offset = offset
	return s
}

func (s *FindAll) WithOrderBy(orderBy string) *FindAll {
	s.OrderBy = append(s.OrderBy, orderBy)
	return s
}

func (s *FindAll) WithWhereAllWithDeleted() *FindAll {
	s.WhereAllWithDeleted = true
	return s
}

func (s *FindAll) Exec(ctx context.Context, opts *ExecOptions) (r *BaseResponse, err error) {
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

	respBytes, err := sqlexecutor.NewSQLExecutorClient(client.conn).FindAll(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		return
	}

	result := &BaseResponse{}
	err = json.Unmarshal(respBytes.Value, result)
	return result, err
}
