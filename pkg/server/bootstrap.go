package server

import (
	"context"
	"github.com/vothanhdo2602/hicon/internal/orm"
	"github.com/vothanhdo2602/hicon/internal/rd"
	"sync"
)

func Bootstrap(ctx context.Context) {
	var (
		wg sync.WaitGroup
	)

	modules := []func(context.Context, *sync.WaitGroup){
		rd.Init,
	}
	for _, f := range modules {
		wg.Add(1)
		go f(ctx, &wg)
	}
	wg.Wait()
}
