package main

import (
	"context"
	"fmt"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/hicon-sm/sqlexecutor"
	"github.com/vothanhdo2602/hicon/internal/grpcapi"
	"google.golang.org/grpc"
	"net"
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
		UpsertConfiguration(ctx)

		for i := 0; i < 1; i++ {
			FindAll(ctx)
			//go BulkInsert(ctx)
			//go FindByPK(ctx)
			//go FindOne(ctx)
			//go UpdateByPK(ctx)
			//go BulkUpdateByPK(ctx)
			//go UpdateAll(ctx)
			//go DeleteByPK(ctx)
			//go BulkWriteWithTx(ctx)
		}
	}()

	l, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info(fmt.Sprintf("⚡️[grpc server]: listen on %s", addr))

	defer srv.GracefulStop()
	if err = srv.Serve(l); err != nil {
		logger.Fatal(err.Error())
	}
}
