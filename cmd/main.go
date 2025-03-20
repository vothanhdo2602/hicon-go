package main

import (
	"context"
	"fmt"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/util/grpctil"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/hicon-sm/sqlexecutor"
	"github.com/vothanhdo2602/hicon/internal/grpcapi"
	"github.com/vothanhdo2602/hicon/internal/natsio/reqrep"
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
		srv    = grpc.NewServer(grpc.MaxConcurrentStreams(0))
		addr   = config.GetAddr()
	)

	sqlexecutor.RegisterSQLExecutorServer(srv, grpcapi.SQLExecutor{})

	//init workers
	//wp := wkrtil.NewWorkerPool()
	//defer wp.Stop()

	// init nats
	//err := natsio.Init(ctx, nil)
	//if err != nil {
	//	logger.Fatal(err.Error())
	//}
	//defer natsio.GracefulStop(ctx)

	go func() {
		UpsertConfiguration(ctx)

		for i := 0; i < 10; i++ {
			go FindAll(ctx)
			//go BulkInsert(ctx)
			//go FindByPK(ctx)
			//go FindOne(ctx)
			//go UpdateByPK(ctx)

			//go FindByPKReqrep(ctx)
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
					Columns: []*sqlexecutor.Column{
						{Name: "id", Type: "text", IsPrimaryKey: true},
						{Name: "type", Type: "string"},
						{Name: "created_at", Type: "time"},
						{Name: "deleted_at", Type: "time", SoftDelete: true},
					},
					RelationColumns: []*sqlexecutor.RelationColumn{
						{Name: "profile", Type: orm.HasOne, RefTable: "profiles", Join: "id=user_id"},
					},
				},
				{
					Name: "profiles",
					Columns: []*sqlexecutor.Column{
						{Name: "id", Type: "text", IsPrimaryKey: true},
						{Name: "user_id", Type: "string"},
						{Name: "email", Type: "string"},
						{Name: "name", Type: "string"},
						{Name: "deleted_at", Type: "time", SoftDelete: true},
					},
				},
			},
			Redis: &sqlexecutor.Redis{
				Host:     "localhost",
				Port:     6379,
				Username: "hicon",
				Password: "hicon_private_pwd",
				PoolSize: 500,
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

func FindByPK(ctx context.Context) {
	var (
		req = &sqlexecutor.FindByPK{
			Table: "users",
			Data: map[string]*anypb.Any{
				"id": grpcapi.InterfaceToAnyPb("67c567cd8b606b2293af1519"),
			},
			Select: []string{"id"},
		}
	)

	conn, err := grpctil.NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).FindByPK(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data)
}

func FindOne(ctx context.Context) {
	var (
		req = &sqlexecutor.FindOne{
			Table:        "users",
			DisableCache: false,
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
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data)
}

func FindAll(ctx context.Context) {
	var (
		req = &sqlexecutor.FindAll{
			Table:        "users",
			DisableCache: false,
			Select:       []string{},
			Where:        []*sqlexecutor.Where{},
			Relations:    []string{"profile"},
			Offset:       0,
			OrderBy:      []string{},
		}
	)

	conn, err := grpctil.NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).FindAll(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data)
}

func Exec(ctx context.Context) {
	var (
		req = &sqlexecutor.Exec{
			Sql: `SELECT "users"."id", "users"."type", "profile"."name" AS "profile__name", "profile"."id" AS "profile__id", "profile"."user_id" AS "profile__user_id", "profile"."email" AS "profile__email" FROM "users" LEFT JOIN "profiles" AS "profile" ON ("profile"."user_id" = "users"."id")`,
		}
	)

	conn, err := grpctil.NewClient()
	if err != nil {
		return
	}

	_, err = sqlexecutor.NewSQLExecutorClient(conn).Exec(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//fmt.Println("resp", resp.Data)
}

func BulkInsert(ctx context.Context) {
	var (
		req = &sqlexecutor.BulkInsert{
			Table:        "users",
			DisableCache: true,
			Data:         []*anypb.Any{},
		}
	)

	data := []map[string]interface{}{
		{
			"id":   "67d299ad4244a581108b7ca5",
			"type": "system",
		},
	}

	dataConverted, _ := grpcapi.ConvertSliceAnyToPbAnySlice(data)
	req.Data = dataConverted

	conn, err := grpctil.NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).BulkInsert(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data)
}

func UpdateByPK(ctx context.Context) {
	var (
		req = &sqlexecutor.UpdateByPK{
			Table:        "users",
			DisableCache: false,
		}
	)

	data := map[string]interface{}{
		"id":   "67d299ad4244a581108b7da4",
		"type": "system",
		//"created_at": time.Now(),
	}

	dataConverted, _ := grpcapi.ConvertInterfaceToAny(data)
	req.Data = dataConverted

	conn, err := grpctil.NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).UpdateByPK(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data)
}

func FindByPKReqrep(ctx context.Context) {
	var (
		req = &requestmodel.FindByPK{
			Table: "users",
			Data: map[string]interface{}{
				"id": "67c567cd8b606b2293af1519",
			},
		}
	)

	resp, err := reqrep.FindByPK(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if resp.Data == nil {
		fmt.Println("resp", resp.Data)
	}
}
