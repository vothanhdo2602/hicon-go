package hicon

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon-sm/sqlexecutor"
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

func (s *FindByPK) Exec(ctx context.Context, opts ExecOptions) (r *sqlexecutor.BaseResponse, err error) {
	headers := map[string]string{}
	if opts.RequestID != "" {
		headers["X-Request-ID"] = opts.RequestID
	}

	reqBytes, err := json.Marshal(&BaseRequest{Body: s, Headers: headers})
	if err != nil {
		return
	}

	return sqlexecutor.NewSQLExecutorClient(client.conn).FindByPK(ctx, &anypb.Any{Value: reqBytes})
}
