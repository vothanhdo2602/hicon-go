package hicon

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/vothanhdo2602/hicon-go"
	"github.com/vothanhdo2602/hicon-sm/constant"
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
			//Password: "hicon_private_pwd",
			PoolSize: 500,
		}),
		hicon.WithTable(&hicon.TableConfig{
			Name: "users",
			Columns: []*hicon.Column{
				{Name: "id", Type: constant.ColumnTypeString, IsPrimaryKey: true},
				{Name: "username", Type: constant.ColumnTypeString},
				{Name: "email", Type: constant.ColumnTypeString},
				{Name: "first_name", Type: constant.ColumnTypeString},
				{Name: "last_name", Type: constant.ColumnTypeString},
				{Name: "created_at", Type: constant.ColumnTypeTimestamp},
				{Name: "is_active", Type: constant.ColumnTypeBoolean},
			},
			RelationColumns: []*hicon.RelationColumn{
				{
					Name:     "roles",
					RefTable: "user_roles",
					Type:     constant.HasManyToMany,
					Join:     "user=role",
				},
			},
		}),
		hicon.WithTable(&hicon.TableConfig{
			Name: "roles",
			Columns: []*hicon.Column{
				{Name: "id", Type: constant.ColumnTypeString, IsPrimaryKey: true},
				{Name: "role_name", Type: constant.ColumnTypeString},
				{Name: "description", Type: constant.ColumnTypeString},
				{Name: "created_at", Type: constant.ColumnTypeString},
				{Name: "is_active", Type: constant.ColumnTypeBoolean},
			},
			RelationColumns: []*hicon.RelationColumn{
				{
					Name:     "permissions",
					RefTable: "permissions",
					Type:     constant.HasManyToMany,
					Join:     "id=permission_id",
				},
			},
		}),
		hicon.WithTable(&hicon.TableConfig{
			Name: "permissions",
			Columns: []*hicon.Column{
				{Name: "id", Type: constant.ColumnTypeString, IsPrimaryKey: true},
				{Name: "permission_name", Type: constant.ColumnTypeString},
				{Name: "description", Type: constant.ColumnTypeString},
				{Name: "resource", Type: constant.ColumnTypeString},
				{Name: "action", Type: constant.ColumnTypeTimestamp},
				{Name: "created_at", Type: constant.ColumnTypeTimestamp},
			},
		}),
		hicon.WithTable(&hicon.TableConfig{
			Name: "user_roles",
			Columns: []*hicon.Column{
				{Name: "user_id", Type: constant.ColumnTypeString, IsPrimaryKey: true},
				{Name: "role_id", Type: constant.ColumnTypeString, IsPrimaryKey: true},
				{Name: "assigned_at", Type: constant.ColumnTypeTimestamp},
				{Name: "assigned_by", Type: constant.ColumnTypeString},
			},
			RelationColumns: []*hicon.RelationColumn{
				{
					Name:     "user",
					RefTable: "users",
					Type:     constant.BelongsTo,
					Join:     "user_id=id",
				}, {
					Name:     "role",
					RefTable: "roles",
					Type:     constant.BelongsTo,
					Join:     "role_id=id",
				},
			},
		}),
		hicon.WithTable(&hicon.TableConfig{
			Name: "role_permissions",
			Columns: []*hicon.Column{
				{Name: "role_id", Type: constant.ColumnTypeString, IsPrimaryKey: true},
				{Name: "permission_id", Type: constant.ColumnTypeString, IsPrimaryKey: true},
				{Name: "granted_at", Type: constant.ColumnTypeTimestamp},
			},
			RelationColumns: []*hicon.RelationColumn{
				{
					Name:     "permission",
					RefTable: "permissions",
					Type:     constant.BelongsTo,
					Join:     "permission_id=id",
				}, {
					Name:     "role",
					RefTable: "roles",
					Type:     constant.BelongsTo,
					Join:     "role_id=id",
				},
			},
		}),
	).Exec(ctx, nil)
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
			WithData(map[string]interface{}{
				"id": "67c567cd8b606b2293af1519",
			}).
			WithSelects("name")
	)

	resp, err := query.Exec(ctx, nil)
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

	resp, err := query.Exec(ctx, nil)
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
			WithLimit(10).
			WithOffset(2)
	)

	resp, err := query.Exec(ctx, nil)
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

	resp, err := query.Exec(ctx, nil)
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
			Cache(true).
			WithLockKey(fmt.Sprintf("create_user:%s", userID)).
			WithData([]interface{}{
				map[string]interface{}{
					"id":   userID,
					"type": "system",
				},
			})
	)

	resp, err := query.Exec(ctx, nil)
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
			WithData(map[string]interface{}{
				"id":   userID,     // auto set as condition if you set it as primary key in UpsertConfig func
				"type": "external", // update field
			}).
			WithWhere("type = ?", "system") // add some condition
	)

	resp, err := query.Exec(ctx, nil)
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
			WithSet("type = ?", "external").
			WithWhere("type = ?", "system") // add some condition
	)

	resp, err := query.Exec(ctx, nil)
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
			WithSet("type").
			WithWhere("updated_at").
			WithData([]interface{}{
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

	resp, err := query.Exec(ctx, nil)
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
			WithData(map[string]interface{}{
				"id": userID, // auto set as condition if you set it as primary key in UpsertConfig func
			})
	)

	resp, err := query.Exec(ctx, nil)
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
				WithSet("type").
				WithWhere("updated_at").
				WithData([]interface{}{
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
				WithData(map[string]interface{}{
				"id":   userID,     // auto set as condition if you set it as primary key in UpsertConfig func
				"type": "external", // update field
			}).
			WithWhere("type = ?", "system")
	)

	var (
		query = c.NewBulkWriteWithTx(bulkUpdateByPK.ToOperation(), updateByPK.ToOperation())
	)

	resp, err := query.Exec(ctx, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("resp", resp.Data, resp.Shared, resp.Message)
}
