// Package hicon provides a client SDK for optimizing database queries.
//
// MIT License - see LICENSE file for details.
package hicon

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/vothanhdo2602/hicon-go"
	"github.com/vothanhdo2602/hicon-go/hicon-sm/constant"
	"time"
)

func UpsertConfig(ctx context.Context) {
	c, _ := hicon.NewClient(ctx, "localhost:7979")
	r, err := c.NewUpsertConfig(
		hicon.WithDBConfig(&hicon.DBConfig{
			Type:     "postgres",
			Host:     "localhost",
			Port:     5432,
			Username: "hicon",
			Password: "hicon_private_pwd",
			Database: "hicon_database",
			MaxCons:  90,
		}),
		hicon.WithDebug(true),
		hicon.WithRedis(&hicon.Redis{
			Host:     "localhost",
			Port:     6379,
			Username: "hicon",
			Password: "hicon_private_pwd",
			PoolSize: 500,
		}),
		hicon.WithTable(
			&hicon.TableConfig{
				Name: "users",
				Columns: []*hicon.Column{
					{Name: "id", Type: "text", IsPrimaryKey: true},
					{Name: "type", Type: "string"},
					{Name: "created_at", Type: "time"},
					{Name: "deleted_at", Type: "time", SoftDelete: true},
				},
				RelationColumns: []*hicon.RelationColumn{
					{Name: "profile", Type: constant.HasOne, RefTable: "profiles", Join: "id=user_id"},
				},
			},
		),
		hicon.WithTable(
			&hicon.TableConfig{
				Name: "profiles",
				Columns: []*hicon.Column{
					{Name: "id", Type: "text", IsPrimaryKey: true},
					{Name: "user_id", Type: "string"},
					{Name: "email", Type: "string"},
					{Name: "name", Type: "string"},
					{Name: "deleted_at", Type: "time", SoftDelete: true},
				},
			},
		),
	).Exec(ctx)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("success: ", r)
}

func FindByPK(ctx context.Context) {
	c, _ := hicon.NewClient(ctx, "localhost:7979")

	var (
		query = c.NewFindByPK("users").
			Data(map[string]interface{}{
				"id": "67c567cd8b606b2293af1519",
			}).
			Selects("name")
	)

	resp, err := query.Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data, resp.Shared, resp.Message)
}

func FindOne(ctx context.Context) {
	c, _ := hicon.NewClient(ctx, "localhost:7979")

	var (
		query = c.NewFindOne("users").
			Relation("Profile")
	)

	resp, err := query.Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data, resp.Shared, resp.Message)
}

func FindAll(ctx context.Context) {
	c, _ := hicon.NewClient(ctx, "localhost:7979")

	var (
		query = c.NewFindAll("users").
			Relation("Profile").
			Limit(10).
			Offset(2)
	)

	resp, err := query.Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data, resp.Shared, resp.Message)
}

func Exec(ctx context.Context) {
	c, _ := hicon.NewClient(ctx, "localhost:7979")

	var (
		sql   = `SELECT "users"."id", "users"."type", "profile"."name" AS "profile__name", "profile"."id" AS "profile__id", "profile"."user_id" AS "profile__user_id", "profile"."email" AS "profile__email" FROM "users" LEFT JOIN "profiles" AS "profile" ON ("profile"."user_id" = "users"."id")`
		query = c.NewExec(sql).WithLockKey(sql)
	)

	resp, err := query.Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data, resp.Shared, resp.Message)
}

func BulkInsert(ctx context.Context) {
	c, _ := hicon.NewClient(ctx, "localhost:7979")

	var (
		userID = uuid.New()
		query  = c.NewBulkInsert("users").
			WithDisableCache().
			WithLockKey(fmt.Sprintf("create_user:%s", userID)).
			Data([]interface{}{
				map[string]interface{}{
					"id":   userID,
					"type": "system",
				},
			})
	)

	resp, err := query.Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data, resp.Shared, resp.Message)
}

func UpdateByPK(ctx context.Context) {
	c, _ := hicon.NewClient(ctx, "localhost:7979")

	var (
		userID = "67c567cd8b606b2293af1"
		query  = c.NewUpdateByPK("users").
			WithLockKey(fmt.Sprintf("update_user:%s", userID)).
			Data(map[string]interface{}{
				"id":   userID,     // auto set as condition if you set it as primary key in UpsertConfig func
				"type": "external", // update field
			}).
			Where("type = ?", "system") // add some condition
	)

	resp, err := query.Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data, resp.Shared, resp.Message)
}

func UpdateAll(ctx context.Context) {
	c, _ := hicon.NewClient(ctx, "localhost:7979")

	var (
		query = c.NewUpdateAll("users").
			WithLockKey(fmt.Sprintf("update_user_all_user_type_system")).
			Set("type = ?", "external").
			Where("type = ?", "system") // add some condition
	)

	resp, err := query.Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data, resp.Shared, resp.Message)
}

func BulkUpdateByPK(ctx context.Context) {
	c, _ := hicon.NewClient(ctx, "localhost:7979")

	var (
		query = c.NewBulkUpdateByPK("users").
			WithLockKey(fmt.Sprintf("update_user_all_user_type_system")).
			Set("type").
			Where("updated_at").
			Data([]interface{}{
				map[string]interface{}{
					"id":         "67d299ad4244a581108b7da4",
					"updated_at": time.Now(),
					"type":       "system",
					//"created_at": time.Now(),
				},
				map[string]interface{}{
					"id":         "67d299ad4244a581108b7da1",
					"updated_at": time.Now(),
					"type":       "system",
					//"created_at": time.Now(),
				},
			})
	)

	resp, err := query.Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data, resp.Shared, resp.Message)
}

func DeleteByPK(ctx context.Context) {
	c, _ := hicon.NewClient(ctx, "localhost:7979")

	var (
		userID = "67c567cd8b606b2293af1"
		query  = c.NewDeleteByPK("users").
			WithLockKey(fmt.Sprintf("delete_user:%s", userID)).
			Data(map[string]interface{}{
				"id": userID, // auto set as condition if you set it as primary key in UpsertConfig func
			})
	)

	resp, err := query.Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data, resp.Shared, resp.Message)
}

func BulkWriteWithTx(ctx context.Context) {
	c, _ := hicon.NewClient(ctx, "localhost:7979")

	var (
		bulkUpdateByPK = c.NewBulkUpdateByPK("users").
				WithLockKey(fmt.Sprintf("update_user_all_user_type_system")).
				Set("type").
				Where("updated_at").
				Data([]interface{}{
				map[string]interface{}{
					"id":         "67d299ad4244a581108b7da4",
					"updated_at": time.Now(),
					"type":       "system",
					//"created_at": time.Now(),
				},
				map[string]interface{}{
					"id":         "67d299ad4244a581108b7da1",
					"updated_at": time.Now(),
					"type":       "system",
					//"created_at": time.Now(),
				},
			})
		userID     = "67c567cd8b606b2293af1"
		updateByPK = c.NewUpdateByPK("users").
				WithLockKey(fmt.Sprintf("update_user:%s", userID)).
				Data(map[string]interface{}{
				"id":   userID,     // auto set as condition if you set it as primary key in UpsertConfig func
				"type": "external", // update field
			}).
			Where("type = ?", "system")
	)

	var (
		query = c.NewBulkWriteWithTx(bulkUpdateByPK.ToOperation(), updateByPK.ToOperation())
	)

	resp, err := query.Exec(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data, resp.Shared, resp.Message)
}
