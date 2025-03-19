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
	"time"
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
}

func BuildEntity(tableConfig *config.TableConfiguration) ([]reflect.StructField, []reflect.StructField, []reflect.StructField, error) {
	var (
		// Embedded BaseModel with table name tag
		embeddedBaseModel = reflect.StructField{
			Name:      "BaseModel",
			Type:      reflect.TypeOf(bun.BaseModel{}),
			Tag:       reflect.StructTag(fmt.Sprintf(`bun:"table:%s"`, tableConfig.Name)),
			Anonymous: true,
		}
		fields = []reflect.StructField{
			embeddedBaseModel,
		}
		ptrFields = []reflect.StructField{
			embeddedBaseModel,
		}
		refFields = []reflect.StructField{
			embeddedBaseModel,
		}
	)

	// Add fields based on column configurations
	for colName, col := range tableConfig.ColumnConfigs {
		fieldType, ptrFieldType := getGoType(col.Type, col.Nullable)

		// Build bun tag
		tags := []string{fmt.Sprintf("column:%s", colName)}
		if col.IsPrimaryKey {
			tags = append(tags, "pk")
		}

		if col.SoftDelete {
			tags = append(tags, "soft_delete,nullzero")
		} else if col.Nullable {
			tags = append(tags, "nullzero")
		}

		// Add field with both json and bun tags
		var (
			name = cases.Title(language.English).String(colName)
			tag  = reflect.StructTag(fmt.Sprintf(`json:"%s,omitempty" redis:"%s,omitempty" bun:"%s"`, colName, colName, strings.Join(tags, ",")))
		)

		field := reflect.StructField{
			Name: name,
			Type: reflect.TypeOf(fieldType),
			Tag:  tag,
		}
		fields = append(fields, field)

		if col.IsPrimaryKey {
			refFields = append(refFields, field)
		}

		ptrField := reflect.StructField{
			Name: name,
			Type: reflect.TypeOf(ptrFieldType),
			Tag:  tag,
		}
		ptrFields = append(ptrFields, ptrField)
	}

	// Create instance
	return fields, ptrFields, refFields, nil
}

func MapRelationToEntity(tableConfig *config.TableConfiguration, entities map[string][]reflect.StructField, refModels map[string][]reflect.StructField) error {
	for colName, col := range tableConfig.RelationColumns {
		if _, ok := entities[col.RefTable]; !ok {
			return errors.New(fmt.Sprintf("Table %s not found", col.RefTable))
		}

		var (
			fieldType    reflect.Type
			refFieldType reflect.Type
			modelType    = reflect.StructOf(entities[col.RefTable])
			refModelType = reflect.StructOf(refModels[col.RefTable])
			name         = cases.Title(language.English).String(colName)
			tag          = reflect.StructTag(fmt.Sprintf(`json:"%s,omitempty" bun:"rel:%s,join:%s"`, colName, col.Type, col.Join))
		)

		switch col.Type {
		case HasOne, BelongsTo:
			fieldType = reflect.PointerTo(modelType)
			refFieldType = reflect.PointerTo(refModelType)
		case HasMany, HasManyToMany:
			fieldType = reflect.SliceOf(modelType)
			refFieldType = reflect.SliceOf(refModelType)
		default:
			return errors.New(fmt.Sprintf("unsupported relation type: %s, just in 'array' or 'object'", col.Type))
		}

		// Add field with both json and bun tags
		field := reflect.StructField{
			Name: name,
			Type: fieldType,
			Tag:  tag,
		}
		entities[tableConfig.Name] = append(entities[tableConfig.Name], field)

		refField := reflect.StructField{
			Name: name,
			Type: refFieldType,
			Tag:  tag,
		}
		refModels[tableConfig.Name] = append(refModels[tableConfig.Name], refField)
	}

	return nil
}

func getGoType(dbType string, nullable bool) (interface{}, interface{}) {
	var (
		fieldType    interface{}
		ptrFieldType interface{}
	)

	switch strings.ToLower(dbType) {
	case "string", "text", "varchar", "char":
		//fieldType = ""
		ptrFieldType = (*string)(nil)

		//if nullable {
		fieldType = (*string)(nil)
		//}
	case "time", "timestamp":
		//fieldType = time.Time{}
		ptrFieldType = (*time.Time)(nil)

		//if nullable {
		fieldType = (*time.Time)(nil)
		//}
	case "int", "integer", "bigint":
		//fieldType = 0
		ptrFieldType = (*int)(nil)

		//if nullable {
		fieldType = (*int)(nil)
		//}
	case "float", "double", "decimal":
		//fieldType = float64(0)
		ptrFieldType = (*float64)(nil)

		//if nullable {
		fieldType = (*float64)(nil)
		//}
	case "bool", "boolean":
		//fieldType = false
		ptrFieldType = (*bool)(nil)

		//if nullable {
		fieldType = (*bool)(nil)
		//}
	default:
		fieldType = interface{}(nil)
		ptrFieldType = interface{}(nil)
	}

	return fieldType, ptrFieldType
}
