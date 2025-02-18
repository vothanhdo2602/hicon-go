package commontil

import (
	"context"
	"time"
)

type ContextWithoutDeadline struct {
	ctx context.Context
}

func (*ContextWithoutDeadline) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (*ContextWithoutDeadline) Done() <-chan struct{} {
	return nil
}

func (*ContextWithoutDeadline) Err() error {
	return nil
}

func (s *ContextWithoutDeadline) Value(key interface{}) interface{} {
	return s.ctx.Value(key)
}

func CopyContext(ctx context.Context) context.Context {
	return &ContextWithoutDeadline{ctx}
}
