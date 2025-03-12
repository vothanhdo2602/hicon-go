package main

import (
	"context"
	"fmt"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/util/grpctil"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/external/util/wkrtil"
	"github.com/vothanhdo2602/hicon/hicon-sm/sqlexecutor"
	"github.com/vothanhdo2602/hicon/internal/grpcapi"
	"github.com/vothanhdo2602/hicon/internal/natsio"
	"github.com/vothanhdo2602/hicon/internal/natsio/reqrep"
	"github.com/vothanhdo2602/hicon/internal/orm"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
	"net"
	"time"
)

func main() {
	log.Init()
	config.Init()

	var (
		ctx    = context.Background()
		logger = log.WithCtx(ctx)
		srv    = grpc.NewServer(grpc.MaxConcurrentStreams(0))
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

		time.Sleep(2 * time.Second)

		//for i := 0; i < 100000; i++ {
		go FindOne(ctx)
		//go FindByPrimaryKeys(ctx)
		//go FindByPrimaryKeysReqrep(ctx)
		//}
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
						{Name: "profile", Type: orm.HasOne, RefTable: "profiles", Join: "id=user_id"},
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
			//Select: []string{"id"},
		}
	)

	conn, err := grpctil.NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).FindByPrimaryKeys(ctx, req)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}

	fmt.Println("resp", resp.Data)
}

func FindOne(ctx context.Context) {
	var (
		req = &sqlexecutor.FindOne{
			Table:        "users",
			DisableCache: true,
			Select:       []string{},
			Where:        []*sqlexecutor.Where{},
			Relations:    []string{"Profile"},
			Offset:       0,
			OrderBy:      []string{},
		}
	)

	conn, err := grpctil.NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).FindOne(ctx, req)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}

	fmt.Println("resp", resp.Data)
}

func FindByPrimaryKeysReqrep(ctx context.Context) {
	var (
		req = &requestmodel.FindByPrimaryKeys{
			Table: "users",
			PrimaryKeys: map[string]interface{}{
				"id": "67c567cd8b606b2293af1519",
			},
		}
	)

	resp, err := reqrep.FindByPrimaryKeys(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if resp.Data == nil {
		fmt.Println("resp", resp.Data)
	}
}
