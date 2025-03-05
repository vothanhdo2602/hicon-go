package requestmodel

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type FindByPrimaryKeys struct {
	Table        string
	DisableCache bool
	PrimaryKeys  map[string]interface{}
}

func (m *FindByPrimaryKeys) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Table, validation.Required),
		validation.Field(&m.PrimaryKeys, validation.Required),
	)
}

type FindOne struct {
	Table        string
	DisableCache bool
	Select       []string
	Wheres       []*Where
	Relations    []*Relation
	Offset       int
	OrderBy      []string
}

type Where struct {
	Condition string
	Args      []interface{}
}

type Relation struct {
	Type string
}

func (m *FindOne) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Table, validation.Required),
	)
}
