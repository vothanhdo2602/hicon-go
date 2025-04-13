package requestmodel

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/vothanhdo2602/hicon/hicon-sm/constant"
)

type FindByPK struct {
	Table               string
	DisableCache        bool
	Selects             []string
	Data                interface{}
	WhereAllWithDeleted bool
}

func (s *FindByPK) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Table, validation.Required),
		validation.Field(&s.Data, validation.Required),
	)
}

type FindOne struct {
	Table               string
	DisableCache        bool
	Selects             []string
	Where               []*QueryWithArgs
	Relations           []string
	Joins               []*Join
	Offset              int
	OrderBy             []string
	WhereAllWithDeleted bool
}

func (s *FindOne) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Table, validation.Required),
	)
}

type FindAll struct {
	Table               string
	DisableCache        bool
	Selects             []string
	Where               []*QueryWithArgs
	Relations           []string
	Joins               []*Join
	Limit               int
	Offset              int
	OrderBy             []string
	WhereAllWithDeleted bool
}

func (s *FindAll) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Table, validation.Required),
	)
}

type QueryWithArgs struct {
	Query string
	Args  []interface{}
}

type Join struct {
	Join string
	Args []interface{}
}

type Exec struct {
	LockKey string
	SQL     string
	Args    []interface{}
}

type BulkInsert struct {
	LockKey      string
	Table        string
	Data         []interface{}
	DisableCache bool
}

func (s *BulkInsert) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Table, validation.Required),
		validation.Field(&s.Data, validation.Required),
	)
}

type UpdateByPK struct {
	// Lock key for concurrent insert operations
	// The later task with the same lock key in the same time will not execute and get the result from the first task
	LockKey      string
	Table        string
	Data         interface{}
	Where        []*QueryWithArgs
	DisableCache bool
}

func (s *UpdateByPK) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Table, validation.Required),
		validation.Field(&s.Data, validation.Required),
	)
}

type UpdateAll struct {
	// Lock key for concurrent insert operations
	// The later task with the same lock key in the same time will not execute and get the result from the first task
	LockKey             string
	Table               string
	Where               []*QueryWithArgs
	Set                 []*QueryWithArgs
	WhereAllWithDeleted bool
	DisableCache        bool
}

func (s *UpdateAll) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Table, validation.Required),
		validation.Field(&s.Where, validation.Required),
		validation.Field(&s.Set, validation.Required),
	)
}

type BulkUpdateByPK struct {
	// Lock key for concurrent insert operations
	// The later task with the same lock key in the same time will not execute and get the result from the first task
	LockKey      string      `json:"lock_key"`
	Table        string      `json:"table"`
	Set          []string    `json:"set"`
	Where        []string    `json:"where"`
	Data         interface{} `json:"data"`
	DisableCache bool        `json:"disable_cache"`
}

func (s *BulkUpdateByPK) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Table, validation.Required),
		validation.Field(&s.Data, validation.Required),
	)
}

type DeleteByPK struct {
	// Lock key for concurrent insert operations
	// The later task with the same lock key in the same time will not execute and get the result from the first task
	LockKey      string
	Table        string
	Data         interface{}
	Where        []*QueryWithArgs
	DisableCache bool
	ForceDelete  bool // if enable soft delete in table
}

func (s *DeleteByPK) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Table, validation.Required),
		validation.Field(&s.Data, validation.Required),
	)
}

type BulkWriteWithTx struct {
	LockKey    string
	Operations []*Operation
}

func (s *BulkWriteWithTx) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Operations),
	)
}

type Operation struct {
	Name string
	Data interface{}
}

func (s *Operation) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(
			&s.Name,
			validation.In(
				constant.BWOperationExec,
				constant.BWOperationBulkInsert,
				constant.BWOperationUpdateByPK,
				constant.BWOperationUpdateAll,
				constant.BWOperationBulkUpdateByPK,
				constant.BWOperationDeleteByPK,
			),
		),
		validation.Field(
			&s.Data, validation.Required,
		),
	)
}
