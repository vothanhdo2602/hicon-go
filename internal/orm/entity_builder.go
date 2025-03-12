package orm

import (
	"errors"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/vothanhdo2602/hicon/external/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect"
	"strings"
)

// RelationType defines the type of relationship
const (
	HasOne        = "has-one"
	BelongsTo     = "belongs-to"
	HasMany       = "has-many"
	HasManyToMany = "has-many-to-many"
)

type CustomBaseModel struct {
	bun.BaseModel
	tableName string
}

// TableName implements bun.TableNamer interface
func (m *CustomBaseModel) TableName() string {
	return m.tableName
}

func BuildEntity(tableConfig *config.TableConfiguration) ([]reflect.StructField, error) {
	var (
		fields = []reflect.StructField{
			// Embedded BaseModel with table name tag
			{
				Name:      "BaseModel",
				Type:      reflect.TypeOf(bun.BaseModel{}),
				Tag:       reflect.StructTag(fmt.Sprintf(`bun:"table:%s"`, tableConfig.Name)),
				Anonymous: true,
			},
		}
	)

	// Add fields based on column configurations
	for colName, col := range tableConfig.ColumnConfigs {
		fieldType := getGoType(col.Type, col.Nullable)

		// Build bun tag
		tags := []string{fmt.Sprintf("column:%s", colName)}
		if col.IsPrimaryKey {
			tags = append(tags, "pk")
		}
		if col.Nullable {
			tags = append(tags, "nullzero")
		}

		// Add field with both json and bun tags
		field := reflect.StructField{
			Name: cases.Title(language.English).String(colName),
			Type: reflect.TypeOf(fieldType),
			Tag:  reflect.StructTag(fmt.Sprintf(`json:"%s,omitempty" redis:"%s,omitempty" bun:"%s"`, colName, colName, strings.Join(tags, ","))),
		}
		fields = append(fields, field)
	}

	// Create instance
	return fields, nil
}

func MapRelationToEntity(tableConfig *config.TableConfiguration, entities map[string][]reflect.StructField) error {
	for colName, col := range tableConfig.RelationColumnConfigs {
		if _, ok := entities[col.RefTable]; !ok {
			return errors.New(fmt.Sprintf("Table %s not found", col.RefTable))
		}

		var (
			fieldType reflect.Type
			modelType = reflect.StructOf(entities[col.RefTable])
		)
		switch col.Type {
		case HasOne, BelongsTo:
			fieldType = reflect.PointerTo(modelType)
		case HasMany, HasManyToMany:
			fieldType = reflect.SliceOf(modelType)
		default:
			return errors.New(fmt.Sprintf("unsupported relation type: %s, just in 'array' or 'object'", col.Type))
		}

		// Add field with both json and bun tags
		field := reflect.StructField{
			Name: cases.Title(language.English).String(colName),
			Type: fieldType,
			Tag:  reflect.StructTag(fmt.Sprintf(`json:"%s,omitempty" bun:"rel:%s,join:%s"`, colName, col.Type, col.Join)),
		}
		entities[tableConfig.Name] = append(entities[tableConfig.Name], field)
	}

	return nil
}

func getGoType(dbType string, nullable bool) interface{} {
	var (
		fieldType interface{}
	)

	switch strings.ToLower(dbType) {
	case "string", "text", "varchar", "char":
		fieldType = ""
		if nullable {
			fieldType = (*string)(nil)
		}
	case "int", "integer", "bigint":
		fieldType = 0
		if nullable {
			fieldType = (*int)(nil)
		}
	case "float", "double", "decimal":
		fieldType = float64(0)
		if nullable {
			fieldType = (*float64)(nil)
		}
	case "bool", "boolean":
		fieldType = false
		if nullable {
			fieldType = (*bool)(nil)
		}
	case "time", "timestamp":
		fieldType = ""
		if nullable {
			fieldType = (*string)(nil)
		}
	default:
		fieldType = interface{}(nil)
	}

	return fieldType
}
