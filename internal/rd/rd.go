package rd

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/constant"
	"github.com/vothanhdo2602/hicon/external/model/entity"
	"github.com/vothanhdo2602/hicon/external/util/commontil"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/external/util/pjson"
	"runtime"
	"sync"
)

var (
	client *redis.Client
)

func Init(ctx context.Context, wg *sync.WaitGroup) {
	var (
		logger = log.WithCtx(ctx)
		rdEnv  = config.GetENV().DB.Redis
		addr   = fmt.Sprintf("%s:%d", rdEnv.Host, rdEnv.Port)
	)

	if wg != nil {
		defer wg.Done()
	}

	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: rdEnv.Username,
		Password: rdEnv.Password,
		DB:       0,
		PoolSize: 30 * runtime.GOMAXPROCS(0),
	})

	if err := client.Ping(ctx).Err(); err != nil {
		logger.Error(err.Error())
	}
	logger.Info(fmt.Sprintf("⚡️[redis]: connected to %s", addr))
}

func HSet[T entity.BaseInterface[T]](ctx context.Context, database, table string, m T) {
	var (
		logger = log.WithCtx(ctx)
		pipe   = client.Pipeline()
		key    = entity.GetEntityBucketKey(database, table)
	)

	pipe.HSet(ctx, key, m.GetID(), m)
	pipe.HExpire(ctx, key, constant.Expiry10Minutes, m.GetID())

	if _, err := pipe.Exec(ctx); err != nil {
		logger.Error(err.Error())
	}
}

func HMSet[T entity.BaseInterface[T]](ctx context.Context, database, table string, data []T) {
	var (
		logger = log.WithCtx(ctx)
		pipe   = client.Pipeline()
	)

	if len(data) == 0 {
		return
	}

	key := entity.GetEntityBucketKey(database, table)
	for _, m := range data {
		pipe.HSet(ctx, key, m.GetID(), m)
		pipe.HExpire(ctx, key, constant.Expiry10Minutes, m.GetID())
	}

	if _, err := pipe.Exec(ctx); err != nil {
		logger.Error(err.Error())
	}
}

func HGet[T entity.BaseInterface[T]](ctx context.Context, database, table string, field string) *T {
	var (
		logger = log.WithCtx(ctx)
		output = commontil.InitStructFromGeneric[T]()
		key    = entity.GetEntityBucketKey(database, table)
	)
	r, err := client.HGet(ctx, key, field).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			logger.Error(err.Error())
		}
		return nil
	}

	pjson.Unmarshal(ctx, []byte(r), &output)
	return &output
}

func HMGet[T entity.BaseInterface[T]](ctx context.Context, database, table string, fields []string) []T {
	var (
		logger = log.WithCtx(ctx)
		key    = entity.GetEntityBucketKey(database, table)
	)
	cacheData, err := client.HMGet(ctx, key, fields...).Result()
	models := make([]T, len(cacheData))
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			logger.Error(err.Error())
		}
		return models
	}

	for i, v := range cacheData {
		pjson.Unmarshal(ctx, []byte(v.(string)), &models[i])
	}

	return models
}

func hGetAllWithSQL[T entity.BaseInterface[T]](ctx context.Context, database, table, sql string) (models []T, err error) {
	var (
		logger = log.WithCtx(ctx)
		key    = entity.GetSQLBucketKey(database)
	)
	r, err := client.HGet(ctx, key, sql).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			logger.Error(err.Error())
		}
		return models, err
	}

	pjson.Unmarshal(ctx, []byte(r), &models)

	return models, nil
}

func HSetAllWithSQL[T entity.BaseInterface[T]](ctx context.Context, database, table, sql string, data []T) {
	var (
		logger             = log.WithCtx(ctx)
		pipe               = client.Pipeline()
		mapBucketKeyModels = map[string]map[string]any{}
		refModels          []T
		sqlBucketKey       = entity.GetSQLBucketKey(database)
	)

	if len(data) == 0 {
		return
	}

	//for _, m := range data {
	//	refModels = append(refModels, m.GetCacheEntity(database, table))
	//}

	for bucketKey, models := range mapBucketKeyModels {
		var fields []string
		for field := range models {
			fields = append(fields, field)
		}

		pipe.HSet(ctx, bucketKey, models)
		pipe.HExpire(ctx, bucketKey, constant.Expiry10Minutes, fields...)
	}

	// Reference data
	pipe.HSet(ctx, sqlBucketKey, sql, refModels)
	pipe.HExpire(ctx, sqlBucketKey, constant.Expiry5Minutes, sql)

	if _, err := pipe.Exec(ctx); err != nil {
		logger.Error(err.Error())
	}
}

// func HGetAllWithSQL[T entity.BaseInterface[T]](ctx context.Context, sql string) (models []T, err error) {
// 	var (
// 		pipe = client.Pipeline()
// 	)
//
// 	r, err := hGetAllWithSQL[T](ctx, sql)
// 	if err != nil {
// 		return models, err
// 	}
//
// 	return
// }
