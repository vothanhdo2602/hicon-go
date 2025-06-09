package hicon

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon-sm/constant"
	"github.com/vothanhdo2602/hicon-sm/sqlexecutor"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *Client) NewFindByPK(table string) *FindByPK {
	return &FindByPK{
		Table: table,
	}
}

func (s *FindByPK) Cache(isCache bool) *FindByPK {
	s.DisableCache = !isCache
	return s
}

func (s *FindByPK) WithSelects(columns ...string) *FindByPK {
	s.Selects = append(s.Selects, columns...)
	return s
}

func (s *FindByPK) WithData(data interface{}) *FindByPK {
	s.Data = data
	return s
}

func (s *FindByPK) WithWhereAllWithDeleted() *FindByPK {
	s.WhereAllWithDeleted = true
	return s
}

func (s *FindByPK) Exec(ctx context.Context, opts *ExecOptions) (r *BaseResponse, err error) {
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

	respBytes, err := sqlexecutor.NewSQLExecutorClient(client.conn).FindByPK(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		return
	}

	result := &BaseResponse{}
	err = json.Unmarshal(respBytes.Value, result)
	return result, err
}
