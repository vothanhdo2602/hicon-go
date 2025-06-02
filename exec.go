package hicon

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon-sm/sqlexecutor"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *Client) NewExec(table string, args ...interface{}) *Exec {
	return &Exec{
		sql:  table,
		args: args,
	}
}

func (s *Exec) WithLockKey(lockKey string) *Exec {
	s.lockKey = lockKey
	return s
}

func (s *Exec) Exec(ctx context.Context, opts ExecOptions) (r *sqlexecutor.BaseResponse, err error) {
	headers := map[string]string{}
	if opts.RequestID != "" {
		headers["X-Request-ID"] = opts.RequestID
	}

	reqBytes, err := json.Marshal(&BaseRequest{Body: s, Headers: headers})
	if err != nil {
		return
	}

	return sqlexecutor.NewSQLExecutorClient(client.conn).Exec(ctx, &anypb.Any{Value: reqBytes})
}
