package entity

import (
	"fmt"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/util/pstring"
	"reflect"
	"strings"
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

func GetPK(table string, m interface{}) string {
	var (
		pk   = config.GetModelRegistry().TableConfigurations[table].PrimaryColumns
		keys []string
	)

	val := GetReflectValue(m)
	for k := range pk {
		keys = append(keys, GetValueByNameAsString(val, k))
	}

	return strings.Join(keys, ";")
}

func GetReflectValue(m interface{}) reflect.Value {
	val := reflect.ValueOf(m)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func GetSQLBucketKey(database string) string {
	return fmt.Sprintf("%s:%s", SQLBucketKey, database)
}

func GetValueByNameAsString(val reflect.Value, fieldName string) string {
	field := val.FieldByName(pstring.Title(fieldName))
	if !field.IsValid() {
		return ""
	}

	if field.Kind() == reflect.Ptr {
		field = field.Elem()
	}

	switch field.Kind() {
	case reflect.String:
		return field.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", field.Int())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%g", field.Float())
	case reflect.Bool:
		return fmt.Sprintf("%t", field.Bool())
	default:
		return ""
	}
}

func IsZeroValueField(v interface{}, fieldName string) bool {
	val := GetReflectValue(v)
	if val.IsZero() {
		return true
	}

	field := val.FieldByName(pstring.Title(fieldName))
	if field.IsZero() {
		return true
	}

	return false
}
