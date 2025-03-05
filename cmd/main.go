package main

import (
	"context"
	"fmt"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/external/util/wkrtil"
	"github.com/vothanhdo2602/hicon/internal/natsio"
	"github.com/vothanhdo2602/hicon/internal/natsio/reqrep"
	"github.com/vothanhdo2602/hicon/internal/orm"
)

func main() {
	log.Init()

	var (
		ctx = context.Background()
	)

	//init workers
	wp := wkrtil.NewWorkerPool()
	defer wp.Stop()

	// init nats
	err := natsio.Init(ctx, nil)
	if err != nil {
		return
	}
	defer natsio.GracefulStop(ctx)

	UpdateConfig(ctx)

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println("@@@@@@ send index: ", i)
			FindByPrimaryKeys(ctx, i)
			fmt.Println("@@@@@@ done index: ", i)
		}()
	}

	<-ctx.Done()
}

func UpdateConfig(ctx context.Context) {
	req := &requestmodel.UpdateConfig{
		DBConfiguration: &requestmodel.DBConfiguration{
			Type:     "postgres",
			Host:     "localhost",
			Port:     5432,
			Username: "hicon",
			Password: "hicon_private_pwd",
			Database: "hicon_database",
			MaxCons:  100,
		},
		//Debug: true,
		TableConfigurations: []*requestmodel.TableConfiguration{
			{
				Name: "users",
				ColumnConfigs: []*requestmodel.ColumnConfig{
					{Name: "id", Type: "text", IsPrimaryKey: true},
					{Name: "type", Type: "string"},
				},
				RelationColumnConfigs: []*requestmodel.RelationColumnConfigs{
					{Name: "profile", Type: orm.BelongTo, RefTable: "profiles"},
				},
			},
			{
				Name: "profiles",
				ColumnConfigs: []*requestmodel.ColumnConfig{
					{Name: "id", Type: "text", IsPrimaryKey: true},
					{Name: "user_id", Type: "string"},
					{Name: "email", Type: "string"},
					{Name: "name", Type: "string"},
				},
			},
		},
	}

	if err := reqrep.UpdateConfig(ctx, req); err != nil {
		return
	}
}

func FindByPrimaryKeys(ctx context.Context, i int) {
	var (
		req = &requestmodel.FindByPrimaryKeys{
			Table: "users",
			PrimaryKeys: map[string]interface{}{
				"id":    "67c567cd8b606b2293af1519",
				"index": i,
			},
		}
	)

	if _, err := reqrep.FindByPrimaryKeys(ctx, req); err != nil {
	}
}
