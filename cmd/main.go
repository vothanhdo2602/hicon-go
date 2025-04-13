package main

import (
	"context"
	"fmt"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/hicon-sm/gosdk"
	"github.com/vothanhdo2602/hicon/hicon-sm/sqlexecutor"
	"github.com/vothanhdo2602/hicon/internal/grpcapi"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

func main() {
	log.Init()
	config.Init()

	var (
		ctx    = context.Background()
		logger = log.WithCtx(ctx)
		srv    = grpc.NewServer()
		addr   = config.GetAddr()
	)

	sqlexecutor.RegisterSQLExecutorServer(srv, &grpcapi.SQLExecutor{})

	go func() {
		var wg sync.WaitGroup
		gosdk.UpsertConfig(ctx)
		//

		now := time.Now()
		wg.Add(1)
		for i := 0; i < 1; i++ {
			//go gosdk.BulkInsert(ctx)
			//go gosdk.FindByPK(ctx)
			//go gosdk.FindOne(ctx)
			//go gosdk.FindAll(ctx)
			//go gosdk.UpdateByPK(ctx)
			//go gosdk.BulkUpdateByPK(ctx)
			//go gosdk.UpdateAll(ctx)
			//go gosdk.DeleteByPK(ctx)
			//go gosdk.BulkWriteWithTx(ctx)
		}

		wg.Wait()
		fmt.Printf("Done in %v", time.Since(now))
	}()

	l, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info(fmt.Sprintf("⚡️[grpc server]: listened on %s", addr))

	defer srv.GracefulStop()
	if err = srv.Serve(l); err != nil {
		logger.Fatal(err.Error())
	}
}
