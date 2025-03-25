package main

import (
	"context"
	"fmt"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/util/grpctil"
	"github.com/vothanhdo2602/hicon/hicon-sm/sqlexecutor"
	"github.com/vothanhdo2602/hicon/internal/orm"
	"google.golang.org/protobuf/types/known/anypb"
)

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
			Table:  "users",
			Select: []string{"id"},
		}
	)

	data, _ := grpctil.ConvertInterfaceToPbAny(map[string]interface{}{
		"id": "67c567cd8b606b2293af1519",
	})
	req.Data = data

	conn, err := grpctil.NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).FindByPK(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data, resp.Shared)
}

func FindOne(ctx context.Context) {
	var (
		req = &sqlexecutor.FindOne{
			Table:        "users",
			DisableCache: false,
			Select:       []string{},
			Where:        []*sqlexecutor.QueryWithArgs{},
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
			Where:        []*sqlexecutor.QueryWithArgs{},
			Relations:    []string{"profile"},
			Offset:       0,
			OrderBy:      []string{},
			Limit:        10,
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

	data := []interface{}{
		map[string]interface{}{
			"id":   "67d299ad4244a581108b7ca5",
			"type": "system",
		},
	}

	dataConverted, _ := grpctil.ConvertSliceAnyToPbAnySlice(data)
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

	dataConverted, _ := grpctil.ConvertInterfaceToPbAny(data)
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

func UpdateAll(ctx context.Context) {
	id, _ := grpctil.ConvertInterfaceToPbAny("67c567cd8b606b2293af1")
	var (
		req = &sqlexecutor.UpdateAll{
			Table:        "users",
			DisableCache: false,
			Where: []*sqlexecutor.QueryWithArgs{
				{Query: "id = ?", Args: []*anypb.Any{id}},
			},
			Set: []*sqlexecutor.QueryWithArgs{
				{Query: "id = ?", Args: []*anypb.Any{id}},
			},
		}
	)

	data := map[string]interface{}{
		"id":   "67d299ad4244a581108b7da4",
		"type": "system",
		//"created_at": time.Now(),
	}

	dataConverted, _ := grpctil.ConvertInterfaceToPbAny(data)
	req.Data = dataConverted

	conn, err := grpctil.NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).UpdateAll(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data)
}

func BulkUpdateByPK(ctx context.Context) {
	var (
		req = &sqlexecutor.BulkUpdateByPK{
			Table:        "users",
			DisableCache: false,
			Set:          []string{"type"},
		}
	)

	data := []interface{}{
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
	}

	dataConverted, _ := grpctil.ConvertSliceAnyToPbAnySlice(data)
	req.Data = dataConverted

	conn, err := grpctil.NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).BulkUpdateByPK(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data)
}

func DeleteByPK(ctx context.Context) {
	var (
		req = &sqlexecutor.DeleteByPK{
			Table:        "users",
			DisableCache: false,
		}
	)

	data := map[string]interface{}{
		"id":   "67d299ad4244a581108b7da4",
		"type": "system",
		//"created_at": time.Now(),
	}

	dataConverted, _ := grpctil.ConvertInterfaceToPbAny(data)
	req.Data = dataConverted

	conn, err := grpctil.NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).DeleteByPK(ctx, req)
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

	bulkUpdateByPKBytes, _ := grpctil.ConvertInterfaceToPbAny(bulkUpdateByPK)
	updateByPKBytes, _ := grpctil.ConvertInterfaceToPbAny(updateByPK)

	var (
		req = &sqlexecutor.BulkWriteWithTx{
			Operations: []*sqlexecutor.Operation{
				{
					Name: "bulk_update_by_pk",
					Data: bulkUpdateByPKBytes,
				},
				{
					Name: "update_by_pk",
					Data: updateByPKBytes,
				},
			},
		}
	)

	conn, err := grpctil.NewClient()
	if err != nil {
		return
	}

	resp, err := sqlexecutor.NewSQLExecutorClient(conn).BulkWriteWithTx(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data, resp.Message)
}
