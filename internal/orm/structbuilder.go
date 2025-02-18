package orm

import (
	"errors"
	"fmt"
	dynamicstruct "github.com/ompluscator/dynamic-struct"
	"github.com/uptrace/bun"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/util/ptr"
	"reflect"
	"strings"
)

// RelationType defines the type of relationship
const (
	HasOne        = "has-one"
	BelongTo      = "belongs-to"
	HasMany       = "has-many"
	HasManyToMany = "has-many-to-many"
)

func BuildEntity(tableConfig *config.TableConfig) (dynamicstruct.Builder, error) {
	// Add bun.BaseModel as embedded field
	def := dynamicstruct.NewStruct().AddField("BaseModel", bun.BaseModel{}, fmt.Sprintf(`bun:"table:"%s"`, tableConfig.Name))

	// Add fields based on column configurations
	for colName, col := range tableConfig.ColumnConfigs {
		fieldType := getGoType(col.Type, col.Nullable)
		if fieldType == nil {
			return nil, fmt.Errorf("unsupported field type: %s", col.Type)
		}

		// Build bun tag
		tags := []string{fmt.Sprintf("column:%s", colName)}
		if col.IsPrimaryKey {
			tags = append(tags, "pk")
		}
		if col.Nullable {
			tags = append(tags, "nullzero")
		}

		// Add field with both json and bun tags
		def.AddField(
			strings.Title(colName), // Convert to exported field name
			fieldType,
			fmt.Sprintf(`json:"%s,omitempty" redis:"%s,omitempty" bun:"%s"`, colName, colName, strings.Join(tags, ",")),
		)
	}

	// Create instance
	instance := def
	return instance, nil
}

func MapRelationToEntity(tableConfig *config.TableConfig, entities map[string]dynamicstruct.Builder) error {
	for colName, col := range tableConfig.RelationColumnConfigs {
		if _, ok := entities[col.RefTable]; !ok {
			return errors.New(fmt.Sprintf("Table %s not found", col.RefTable))
		}

		var (
			fieldType interface{}
			modelType = reflect.TypeOf(entities[col.RefTable].Build())
		)
		switch col.Type {
		case HasOne, BelongTo:
			fieldType = ptr.ToPtr(modelType)
		case HasMany, HasManyToMany:
			fieldType = reflect.SliceOf(modelType)
		default:
			return errors.New(fmt.Sprintf("unsupported relation type: %s, just in 'array' or 'object'", col.Type))
		}

		// Add field with both json and bun tags
		title := strings.Title(colName)
		entities[colName].AddField(
			title, // Convert to exported field name
			fieldType,
			fmt.Sprintf(`json:"%s" bun:"%s"`, strings.ToLower(title), col.Type),
		)
	}

	return nil
}

func getGoType(dbType string, nullable bool) interface{} {
	var fieldType interface{}

	switch strings.ToLower(dbType) {
	case "string", "text", "varchar", "char":
		if nullable {
			fieldType = (*string)(nil)
		} else {
			fieldType = ""
		}
	case "int", "integer", "bigint":
		if nullable {
			fieldType = (*int64)(nil)
		} else {
			fieldType = int64(0)
		}
	case "float", "double", "decimal":
		if nullable {
			fieldType = (*float64)(nil)
		} else {
			fieldType = float64(0)
		}
	case "bool", "boolean":
		if nullable {
			fieldType = (*bool)(nil)
		} else {
			fieldType = false
		}
	case "time", "timestamp":
		fieldType = (*string)(nil) // Using string for time fields for simplicity
	default:
		fieldType = interface{}(nil)
	}

	return fieldType
}
