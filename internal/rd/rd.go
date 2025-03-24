package rd

import (
	"context"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/constant"
	"github.com/vothanhdo2602/hicon/external/model/entity"
	"github.com/vothanhdo2602/hicon/external/util/commontil"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"github.com/vothanhdo2602/hicon/external/util/pjson"
	"github.com/vothanhdo2602/hicon/external/util/pstring"
	"github.com/vothanhdo2602/hicon/internal/orm"
	"reflect"
	"sync"
	"time"
)

var (
	client *redis.Client
)

func Init(ctx context.Context, wg *sync.WaitGroup, rdEnv *config.Redis) error {
	var (
		logger = log.WithCtx(ctx)
		addr   = fmt.Sprintf("%s:%d", rdEnv.Host, rdEnv.Port)
	)

	if wg != nil {
		defer wg.Done()
	}

	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: rdEnv.Username,
		//Password: rdEnv.Password,
		DB:           0,
		PoolSize:     rdEnv.PoolSize,
		MaxIdleConns: rdEnv.PoolSize,
		MinIdleConns: rdEnv.PoolSize,
		//MaxActiveConns:  3,
		ConnMaxIdleTime: time.Duration(300) * time.Second,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info(fmt.Sprintf("⚡️[redis]: connected to %s", addr))

	return nil
}

func GetRedis() *redis.Client {
	return client
}

func HSet(ctx context.Context, mp *constant.ModelParams, m interface{}) {
	var (
		logger = log.WithCtx(ctx)
		pipe   = client.Pipeline()
		key    = entity.GetEntityBucketKey(mp.Database, mp.Table)
		pk     = entity.GetPK(mp.Table, m)
	)

	v, _ := json.Marshal(m)
	pipe.HSet(ctx, key, pk, v)
	pipe.HExpire(ctx, key, constant.Expiry10Minutes, pk)

	if _, err := pipe.Exec(ctx); err != nil {
		logger.Error(err.Error())
	}
}

func HGet(ctx context.Context, mp *constant.ModelParams, m interface{}) interface{} {
	var (
		logger = log.WithCtx(ctx)
		key    = entity.GetEntityBucketKey(mp.Database, mp.Table)
		pk     = entity.GetPK(mp.Table, m)
	)

	r, err := client.HGet(ctx, key, pk).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			logger.Error(err.Error())
		}
		return nil
	}

	go client.HExpire(commontil.CopyContext(ctx), key, constant.Expiry10Minutes, pk)

	_ = pjson.Unmarshal(ctx, []byte(r), &m)

	return m
}

func HDel(ctx context.Context, mp *constant.ModelParams, m interface{}) {
	if client == nil {
		return
	}

	var (
		logger = log.WithCtx(ctx)
		key    = entity.GetEntityBucketKey(mp.Database, mp.Table)
		pk     = entity.GetPK(mp.Table, m)
	)

	if _, err := client.HDel(ctx, key, pk).Result(); err != nil {
		logger.Error(err.Error())
	}
}

func HSetSQL(ctx context.Context, mp *constant.ModelParams, sql string, m interface{}) {
	var (
		logger       = log.WithCtx(ctx)
		pipe         = client.Pipeline()
		sqlBucketKey = entity.GetSQLBucketKey(mp.Database)
	)

	// Reference data
	refModel, err := config.TransformModel(mp.Table, nil, m, config.RefModelType)
	if err != nil {
		return
	}

	refModelBytes, _ := json.Marshal(refModel)
	pipe.HSet(ctx, sqlBucketKey, sql, refModelBytes)
	pipe.HExpire(ctx, sqlBucketKey, constant.Expiry5Minutes, sql)

	CacheNestedModel(ctx, mp, m, pipe)

	_, err = pipe.Exec(ctx)
	if err != nil {
		logger.Error(err.Error())
	}
}

func HGetSQL(ctx context.Context, mp *constant.ModelParams, sql string, m interface{}, findByPKFn func(context.Context, interface{}, string, *constant.ModelParams) (interface{}, error, bool)) interface{} {
	var (
		logger       = log.WithCtx(ctx)
		sqlBucketKey = entity.GetSQLBucketKey(mp.Database)
	)

	r, err := client.HGet(ctx, sqlBucketKey, sql).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			logger.Error(err.Error())
		}
		return nil
	}

	go client.HExpire(commontil.CopyContext(ctx), sqlBucketKey, constant.Expiry5Minutes, sql)

	_ = pjson.Unmarshal(ctx, []byte(r), &m)

	m = FulfillNestedModel(ctx, mp, m, findByPKFn)

	return m
}

func HSetAllSQL(ctx context.Context, mp *constant.ModelParams, sql string, models interface{}) {
	var (
		logger       = log.WithCtx(ctx)
		pipe         = client.Pipeline()
		sqlBucketKey = entity.GetSQLBucketKey(mp.Database)
	)

	// Reference data
	refModel, err := config.TransformModels(mp.Table, nil, models, config.RefModelType)
	if err != nil {
		return
	}

	refModelBytes, _ := json.Marshal(refModel)
	pipe.HSet(ctx, sqlBucketKey, sql, refModelBytes)
	pipe.HExpire(ctx, sqlBucketKey, constant.Expiry5Minutes, sql)

	if asserted, ok := models.(*[]interface{}); ok {
		newModels := *asserted
		for i := range newModels {
			CacheNestedModel(ctx, mp, newModels[i], pipe)
		}
	}

	_, err = pipe.Exec(ctx)
	if err != nil {
		logger.Error(err.Error())
	}
}

func HGetAllSQL(ctx context.Context, mp *constant.ModelParams, sql string, models interface{}, findByPKFn func(context.Context, interface{}, string, *constant.ModelParams) (interface{}, error, bool)) interface{} {
	var (
		logger       = log.WithCtx(ctx)
		sqlBucketKey = entity.GetSQLBucketKey(mp.Database)
		wg           sync.WaitGroup
	)

	r, err := client.HGet(ctx, sqlBucketKey, sql).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			logger.Error(err.Error())
		}
		return nil
	}

	go client.HExpire(commontil.CopyContext(ctx), sqlBucketKey, constant.Expiry5Minutes, sql)

	_ = pjson.Unmarshal(ctx, []byte(r), &models)

	if asserted, ok := models.(*[]interface{}); ok {
		newModels := *asserted
		wg.Add(len(newModels))
		for i := range newModels {
			go func(i int) {
				defer wg.Done()
				newModels[i] = FulfillNestedModel(ctx, mp, newModels[i], findByPKFn)
			}(i)
		}
		wg.Wait()
	}

	return models
}

func CacheNestedModel(ctx context.Context, mp *constant.ModelParams, m interface{}, pipe redis.Pipeliner) {
	var (
		relCols = config.GetModelRegistry().GetTableConfiguration(mp.Table).RelationColumns
		pk      = entity.GetPK(mp.Table, m)
		key     = entity.GetEntityBucketKey(mp.Database, mp.Table)
	)

	newModel, err := config.TransformModel(mp.Table, nil, m, config.PtrModelType)
	if err != nil {
		return
	}

	newModelBytes, _ := json.Marshal(newModel)
	pipe.HSet(ctx, key, pk, newModelBytes)
	pipe.HExpire(ctx, key, constant.Expiry10Minutes, pk)

	val := entity.GetReflectValue(m)
	for _, c := range relCols {
		fields := val.FieldByName(pstring.Title(c.Name))
		if !fields.IsValid() || fields.IsZero() {
			continue
		}

		newMP := &constant.ModelParams{
			Table:    c.RefTable,
			Database: mp.Database,
			ModeType: config.PtrModelType,
		}

		switch c.Type {
		case orm.HasOne, orm.BelongsTo:
			CacheNestedModel(ctx, newMP, fields.Interface(), pipe)
		case orm.HasMany, orm.HasManyToMany:
			for _, field := range fields.Interface().([]interface{}) {
				CacheNestedModel(ctx, newMP, field, pipe)
			}
		}
	}
}

func FulfillNestedModel(ctx context.Context, mp *constant.ModelParams, m interface{}, findByPKFn func(context.Context, interface{}, string, *constant.ModelParams) (interface{}, error, bool)) interface{} {
	var (
		relCols = config.GetModelRegistry().GetTableConfiguration(mp.Table).RelationColumns
		pk      = entity.GetPK(mp.Table, m)
		wg      sync.WaitGroup
	)

	mp.ModeType = config.DefaultModelType
	newModel, _, _ := findByPKFn(ctx, m, pk, mp)
	if newModel == nil {
		return nil
	}

	val := entity.GetReflectValue(m)
	newVal := entity.GetReflectValue(newModel)
	for _, c := range relCols {
		var (
			name      = pstring.Title(c.Name)
			fields    = val.FieldByName(name)
			newFields = newVal.FieldByName(name)
			newMP     = &constant.ModelParams{
				Table:    c.RefTable,
				Database: mp.Database,
				ModeType: mp.ModeType,
			}
			fieldsInterface = fields.Interface()
		)

		if !fields.IsValid() || fields.IsZero() {
			continue
		}

		switch c.Type {
		case orm.HasOne, orm.BelongsTo:
			wg.Add(1)
			go func(fieldsInterface interface{}) {
				defer wg.Done()
				newFields.Set(reflect.ValueOf(FulfillNestedModel(ctx, newMP, fieldsInterface, findByPKFn)))
			}(fieldsInterface)
		case orm.HasMany, orm.HasManyToMany:
			var (
				newArr = make([]interface{}, len(fieldsInterface.([]interface{})))
			)
			for i, field := range fieldsInterface.([]interface{}) {
				wg.Add(1)
				go func(i int, fieldsInterface interface{}) {
					defer wg.Done()
					newArr[i] = FulfillNestedModel(ctx, newMP, field, findByPKFn)
				}(i, field)
			}

			newFields.Set(reflect.ValueOf(newArr))
		}
	}
	wg.Wait()

	//b, _ := json.MarshalIndent(newModel, "", " ")
	//fmt.Println("MarshalIndent FulfillNestedModel", string(b))

	return newModel
}

func HMSet(ctx context.Context, mp *constant.ModelParams, models interface{}) {
	var (
		logger = log.WithCtx(ctx)
		pipe   = client.Pipeline()
		key    = entity.GetEntityBucketKey(mp.Database, mp.Table)
	)
	if asserted, ok := models.(*[]interface{}); ok {
		newModels := *asserted
		for i := range newModels {
			var (
				pk = entity.GetPK(mp.Table, newModels[i])
			)
			newModelBytes, _ := json.Marshal(newModels[i])
			pipe.HSet(ctx, key, pk, newModelBytes)
			pipe.HExpire(ctx, key, constant.Expiry10Minutes, pk)
			CacheNestedModel(ctx, mp, newModels[i], pipe)
		}
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		logger.Error(err.Error())
	}
}

func HMDel(ctx context.Context, mp *constant.ModelParams, models interface{}) {
	if client == nil {
		return
	}

	var (
		logger = log.WithCtx(ctx)
		pipe   = client.Pipeline()
		key    = entity.GetEntityBucketKey(mp.Database, mp.Table)
		PKs    []string
	)
	if asserted, ok := models.(*[]interface{}); ok {
		newModels := *asserted
		for i := range newModels {
			PKs = append(PKs, entity.GetPK(mp.Table, newModels[i]))
		}
	}

	pipe.HDel(ctx, key, PKs...)
	_, err := pipe.HDel(ctx, key, PKs...).Result()
	if err != nil {
		logger.Error(err.Error())
	}
}
