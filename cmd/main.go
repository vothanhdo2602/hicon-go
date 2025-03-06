package main

import (
	"context"
	"fmt"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/util/grpctil"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/external/util/wkrtil"
	"github.com/vothanhdo2602/hicon/hicon-sm/sqlexecutor"
	"github.com/vothanhdo2602/hicon/internal/grpcapi"
	"github.com/vothanhdo2602/hicon/internal/natsio"
	"github.com/vothanhdo2602/hicon/internal/orm"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
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

	sqlexecutor.RegisterSQLExecutorServer(srv, grpcapi.SQLExecutor{})

	//init workers
	wp := wkrtil.NewWorkerPool()
	defer wp.Stop()

	// init nats
	err := natsio.Init(ctx, nil)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer natsio.GracefulStop(ctx)

	go func() {
		UpsertConfiguration(ctx)

		FindByPrimaryKeys(ctx)
	}()

	l, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info(fmt.Sprintf("⚡️[grpc server]: listen on %s", addr))

	if err = srv.Serve(l); err != nil {
		logger.Fatal(err.Error())
	}

}

func UpsertConfiguration(ctx context.Context) {
	var (
		req = &sqlexecutor.UpsertConfiguration{
			DbConfiguration: &sqlexecutor.DBConfiguration{
				Type:     "postgres",
				Host:     "localhost",
				Port:     5432,
				Username: "hicon",
				Password: "hicon_private_pwd",
				Database: "hicon_database",
				MaxCons:  90,
			},
			Debug: true,
			TableConfigurations: []*sqlexecutor.TableConfiguration{
				{
					Name: "users",
					ColumnConfigs: []*sqlexecutor.ColumnConfig{
						{Name: "id", Type: "text", IsPrimaryKey: true},
						{Name: "type", Type: "string"},
					},
					RelationColumnConfigs: []*sqlexecutor.RelationColumnConfigs{
						{Name: "profile", Type: orm.BelongTo, RefTable: "profiles"},
					},
				},
				{
					Name: "profiles",
					ColumnConfigs: []*sqlexecutor.ColumnConfig{
						{Name: "id", Type: "text", IsPrimaryKey: true},
						{Name: "user_id", Type: "string"},
						{Name: "email", Type: "string"},
						{Name: "name", Type: "string"},
					},
				},
			},
		}
	)

	conn, err := grpctil.NewClient()
	if err != nil {
		return
	}
	defer conn.Close()

	_, err = sqlexecutor.NewSQLExecutorClient(conn).UpsertConfiguration(ctx, req)
	if err != nil {
		return
	}
}

func FindByPrimaryKeys(ctx context.Context) {
	var (
		req = &sqlexecutor.FindByPrimaryKeys{
			Table: "users",
			PrimaryKeys: map[string]*anypb.Any{
				"id": grpcapi.InterfaceToAnyPb("67c567cd8b606b2293af1519"),
			},
		}
	)

	for i := 0; i < 10; i++ {
		go func() {
			conn, err := grpctil.NewClient()
			if err != nil {
				return
			}
			defer conn.Close()

			_, err = sqlexecutor.NewSQLExecutorClient(conn).FindByPrimaryKeys(ctx, req)
			if err != nil {
				return
			}
		}()
	}
}
