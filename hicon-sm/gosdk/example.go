package gosdk

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon/hicon-sm/model/requestmodel"
	"github.com/vothanhdo2602/hicon/hicon-sm/sqlexecutor"
	"github.com/vothanhdo2602/hicon/internal/orm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
)

var (
	conn *grpc.ClientConn
)

func NewClient() (*grpc.ClientConn, error) {
	if conn == nil {
		newConn, err := grpc.NewClient("localhost:7979", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}
		conn = newConn
	}

	return conn, nil
}

func UpsertConfig(ctx context.Context) {
	var (
		req = &requestmodel.UpsertConfig{
			DBConfig: &requestmodel.DBConfig{
				Type:     "postgres",
				Host:     "localhost",
				Port:     5432,
				Username: "hicon",
				Password: "hicon_private_pwd",
				Database: "hicon_database",
				MaxCons:  90,
			},
			Debug: true,
			TableConfigs: []*requestmodel.TableConfig{
				{
					Name: "users",
					Columns: []*requestmodel.Column{
						{Name: "id", Type: "text", IsPrimaryKey: true},
						{Name: "type", Type: "string"},
						{Name: "created_at", Type: "time"},
						{Name: "deleted_at", Type: "time", SoftDelete: true},
					},
					RelationColumns: []*requestmodel.RelationColumn{
						{Name: "profile", Type: orm.HasOne, RefTable: "profiles", Join: "id=user_id"},
					},
				},
				{
					Name: "profiles",
					Columns: []*requestmodel.Column{
						{Name: "id", Type: "text", IsPrimaryKey: true},
						{Name: "user_id", Type: "string"},
						{Name: "email", Type: "string"},
						{Name: "name", Type: "string"},
						{Name: "deleted_at", Type: "time", SoftDelete: true},
					},
				},
			},
			Redis: &requestmodel.Redis{
				Host:     "localhost",
				Port:     6379,
				Username: "hicon",
				Password: "hicon_private_pwd",
				PoolSize: 500,
			},
		}
	)

	reqBytes, err := json.Marshal(&requestmodel.BaseRequest{Body: req})
	if err != nil {
		return
	}

	conn, err = NewClient()
	if err != nil {
		return
	}

	_, err = sqlexecutor.NewSQLExecutorClient(conn).UpsertConfig(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		return
	}
}

func FindByPK(ctx context.Context) {
	var (
		req = &requestmodel.FindByPK{
			Table:        "users",
			Select:       []string{},
			DisableCache: false,
			Data: map[string]interface{}{
				"id": "67c567cd8b606b2293af1519",
			},
		}
	)

	reqBytes, err := json.Marshal(&requestmodel.BaseRequest{Body: req})
	if err != nil {
		return
	}

	conn, err = NewClient()
	if err != nil {
		return
	}

	_, err = sqlexecutor.NewSQLExecutorClient(conn).FindByPK(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//fmt.Println("resp", resp.Data, resp.Shared, resp.Message)
}

func FindOne(ctx context.Context) {
	var (
		req = &requestmodel.FindOne{
			Table:        "users",
			DisableCache: false,
			Select:       []string{},
			Where:        []*requestmodel.QueryWithArgs{},
			Relations:    []string{"Profile"},
			Offset:       0,
			OrderBy:      []string{},
		}
	)

	reqBytes, err := json.Marshal(&requestmodel.BaseRequest{Body: req})
	if err != nil {
		return
	}

	conn, err = NewClient()
	if err != nil {
		return
	}

	_, err = sqlexecutor.NewSQLExecutorClient(conn).FindOne(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//fmt.Println("resp", resp.Data)
}

func FindAll(ctx context.Context) {
	var (
		req = &requestmodel.FindAll{
			Table:        "users",
			DisableCache: false,
			Select:       []string{},
			Where:        []*requestmodel.QueryWithArgs{},
			Relations:    []string{"profile"},
			Offset:       0,
			OrderBy:      []string{},
			Limit:        10,
		}
	)

	reqBytes, err := json.Marshal(&requestmodel.BaseRequest{Body: req})
	if err != nil {
		return
	}

	conn, err = NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).FindAll(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp)
}

func Exec(ctx context.Context) {
	var (
		req = &requestmodel.Exec{
			SQL: `SELECT "users"."id", "users"."type", "profile"."name" AS "profile__name", "profile"."id" AS "profile__id", "profile"."user_id" AS "profile__user_id", "profile"."email" AS "profile__email" FROM "users" LEFT JOIN "profiles" AS "profile" ON ("profile"."user_id" = "users"."id")`,
		}
	)

	reqBytes, err := json.Marshal(&requestmodel.BaseRequest{Body: req})
	if err != nil {
		return
	}

	conn, err = NewClient()
	if err != nil {
		return
	}

	_, err = sqlexecutor.NewSQLExecutorClient(conn).Exec(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//fmt.Println("resp", resp.Data)
}

func BulkInsert(ctx context.Context) {
	var (
		req = &requestmodel.BulkInsert{
			Table:        "users",
			DisableCache: true,
			Data: []interface{}{
				map[string]interface{}{
					"id":   "67d299ad4244a581108b7ca5",
					"type": "system",
				},
			},
		}
	)

	reqBytes, err := json.Marshal(&requestmodel.BaseRequest{Body: req})
	if err != nil {
		return
	}

	conn, err = NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).BulkInsert(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data)
}

func UpdateByPK(ctx context.Context) {
	var (
		req = &requestmodel.UpdateByPK{
			Table:        "users",
			DisableCache: false,
			Data: map[string]interface{}{
				"id":   "67d299ad4244a581108b7da4",
				"type": "system",
				//"created_at": time.Now(),
			},
		}
	)

	reqBytes, err := json.Marshal(&requestmodel.BaseRequest{Body: req})
	if err != nil {
		return
	}

	conn, err = NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).UpdateByPK(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data)
}

func UpdateAll(ctx context.Context) {
	var (
		req = &requestmodel.UpdateAll{
			Table:        "users",
			DisableCache: false,
			Where: []*requestmodel.QueryWithArgs{
				{Query: "id = ?", Args: []interface{}{"67c567cd8b606b2293af1"}},
			},
			Set: []*requestmodel.QueryWithArgs{
				{Query: "id = ?", Args: []interface{}{"67c567cd8b606b2293af1"}},
			},
			Data: map[string]interface{}{
				"id":   "67d299ad4244a581108b7da4",
				"type": "system",
				//"created_at": time.Now(),
			},
		}
	)

	reqBytes, err := json.Marshal(&requestmodel.BaseRequest{Body: req})
	if err != nil {
		return
	}

	conn, err = NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).UpdateAll(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data)
}

func BulkUpdateByPK(ctx context.Context) {
	var (
		req = &requestmodel.BulkUpdateByPK{
			Table:        "users",
			DisableCache: false,
			Set:          []string{"type"},
			Data: []interface{}{
				map[string]interface{}{
					"id":   "67d299ad4244a581108b7da4",
					"type": "system",
					//"created_at": time.Now(),
				},
				map[string]interface{}{
					"id":   "67d299ad4244a581108b7da4",
					"type": "system",
					//"created_at": time.Now(),
				},
			},
		}
	)

	reqBytes, err := json.Marshal(&requestmodel.BaseRequest{Body: req})
	if err != nil {
		return
	}

	conn, err = NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).BulkUpdateByPK(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data)
}

func DeleteByPK(ctx context.Context) {
	var (
		req = &requestmodel.DeleteByPK{
			Table:        "users",
			DisableCache: false,
			Data: map[string]interface{}{
				"id":   "67d299ad4244a581108b7da4",
				"type": "system",
				//"created_at": time.Now(),
			},
		}
	)

	reqBytes, err := json.Marshal(&requestmodel.BaseRequest{Body: req})
	if err != nil {
		return
	}

	conn, err = NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).DeleteByPK(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data)
}

func BulkWriteWithTx(ctx context.Context) {
	var (
		bulkUpdateByPK = &requestmodel.BulkUpdateByPK{
			Table:        "users",
			DisableCache: false,
			Data: []interface{}{
				map[string]interface{}{
					"id":   "67d299ad4244a581108b7da4",
					"type": "system",
					//"created_at": time.Now(),
				},
				map[string]interface{}{
					"id":   "67d299ad4244a581108b7da4",
					"type": "system",
					//"created_at": time.Now(),
				},
			},
		}
		updateByPK = &requestmodel.UpdateByPK{
			Table:        "users",
			DisableCache: false,
			Data: map[string]interface{}{
				"id":   "67d299ad4244a581108b7da4",
				"type": "system",
				//"created_at": time.Now(),
			},
		}
	)

	var (
		req = &requestmodel.BulkWriteWithTx{
			Operations: []*requestmodel.Operation{
				{
					Name: "bulk_update_by_pk",
					Data: bulkUpdateByPK,
				},
				{
					Name: "update_by_pk",
					Data: updateByPK,
				},
			},
		}
	)

	reqBytes, err := json.Marshal(&requestmodel.BaseRequest{Body: req})
	if err != nil {
		return
	}

	conn, err = NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).BulkWriteWithTx(ctx, &anypb.Any{Value: reqBytes})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data, resp.Message)
}
