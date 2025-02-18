package main

import (
	"context"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/external/util/wkrtil"
)

func main() {
	log.Init()

	var (
		ctx = context.Background()
	)

	// workers
	wkrPool := wkrtil.NewWorkerPool()
	defer wkrPool.Stop()

	<-ctx.Done()
}
