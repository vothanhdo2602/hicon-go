package entity

import (
	"fmt"
)

const (
	SQLBucketKey = "sql_bucket"
)

type BaseEntity[T any] interface {
	GetTableName() string
	GetID() string
	MarshalBinary() ([]byte, error) // for redis
	GetCacheEntity(mapBucketKeyModels map[string]map[string]any) *T
}

func GetEntityBucketKey(database, tableName string) string {
	return fmt.Sprintf("%s:%s", database, tableName)
}

func GetSQLBucketKey(database string) string {
	return fmt.Sprintf("%s:%s", database, SQLBucketKey)
}
