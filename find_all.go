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
