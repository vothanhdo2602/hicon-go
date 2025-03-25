package requestmodel

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/vothanhdo2602/hicon/external/constant"
)

type UpsertConfiguration struct {
	DBConfiguration     *DBConfiguration
	Redis               *Redis
	TableConfigurations []*TableConfiguration
	Debug               bool
	DisableCache        bool
}

type DBConfiguration struct {
	Type     string
	Host     string
	Port     int
	Username string
	Password string
	Database string
	MaxCons  int
	TLS      *TLS
}

type Redis struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       int
	PoolSize int
}

type TLS struct {
	CertPEM       string
	PrivateKeyPEM string
	RootCAPEM     string
}

type TableConfiguration struct {
	Name            string
	Columns         []*Column
	RelationColumns []*RelationColumn
}

type Column struct {
	Name         string
	Type         string
	Nullable     bool
	IsPrimaryKey bool
	SoftDelete   bool
}

type RelationColumn struct {
	Name     string
	RefTable string
	Type     string
	Join     string
}

func (m *UpsertConfiguration) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.DBConfiguration, validation.Required),
	)
}

func (m *DBConfiguration) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Type, validation.In(constant.DBPostgres, constant.DBMysql)),
		validation.Field(&m.Host, validation.Required),
		validation.Field(&m.Port, validation.Min(1), validation.Max(65535)),
		validation.Field(&m.MaxCons, validation.Min(1)),
	)
}

type FindByPK struct {
	Table               string
	DisableCache        bool
	Select              []string
	Data                interface{}
	WhereAllWithDeleted bool
}

func (m *FindByPK) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Table, validation.Required),
		validation.Field(&m.Data, validation.Required),
	)
}

type FindOne struct {
	Table               string
	DisableCache        bool
	Select              []string
	Where               []*QueryWithArgs
	Relations           []string
	Join                []*Join
	Offset              int
	OrderBy             []string
	WhereAllWithDeleted bool
}

func (m *FindOne) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Table, validation.Required),
	)
}

type FindAll struct {
	Table               string
	DisableCache        bool
	Select              []string
	Where               []*QueryWithArgs
	Relations           []string
	Joins               []*Join
	Limit               int
	Offset              int
	OrderBy             []string
	WhereAllWithDeleted bool
}

func (m *FindAll) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Table, validation.Required),
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

func (m *BulkInsert) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Table, validation.Required),
		validation.Field(&m.Data, validation.Required),
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

func (m *UpdateByPK) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Table, validation.Required),
		validation.Field(&m.Data, validation.Required),
	)
}

type UpdateAll struct {
	// Lock key for concurrent insert operations
	// The later task with the same lock key in the same time will not execute and get the result from the first task
	LockKey             string
	Table               string
	Data                interface{}
	Where               []*QueryWithArgs
	Set                 []*QueryWithArgs
	WhereAllWithDeleted bool
	DisableCache        bool
}

func (m *UpdateAll) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Table, validation.Required),
		validation.Field(&m.Data, validation.Required),
		validation.Field(&m.Where, validation.Required),
		validation.Field(&m.Set, validation.Required),
	)
}

type BulkUpdateByPK struct {
	// Lock key for concurrent insert operations
	// The later task with the same lock key in the same time will not execute and get the result from the first task
	LockKey      string      `json:"lock_key"`
	Table        string      `json:"table"`
	Set          []string    `json:"set"`
	Data         interface{} `json:"data"`
	DisableCache bool        `json:"disable_cache"`
}

func (m *BulkUpdateByPK) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Table, validation.Required),
		validation.Field(&m.Data, validation.Required),
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
	ForceDelete  bool
}

func (m *DeleteByPK) Validate() error {
	return validation.ValidateStruct(
		m,

		validation.Field(&m.Table, validation.Required),
		validation.Field(&m.Data, validation.Required),
	)
}

type BulkWriteWithTx struct {
	LockKey    string
	Operations []*Operation
}

func (m *BulkWriteWithTx) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Operations),
	)
}

type Operation struct {
	Name string
	Data interface{}
}

func (m *Operation) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Name, validation.In(constant.BWOperationBulkInsert, constant.BWOperationUpdateByPK, constant.BWOperationUpdateAll, constant.BWOperationBulkUpdateByPK, constant.BWOperationDeleteByPK)),
	)
}
