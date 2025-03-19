package orm

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"time"
)

type QueryHook struct{}

func (h *QueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (h *QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	fmt.Println(time.Since(event.StartTime), event.Query)
}
