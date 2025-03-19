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
	ColumnConfigs   []*ColumnConfig
	RelationColumns []*RelationColumns
}

type ColumnConfig struct {
	Name         string
	Type         string
	Nullable     bool
	IsPrimaryKey bool
	SoftDelete   bool
}

type RelationColumns struct {
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
	Table        string
	DisableCache bool
	Select       []string
	Data         map[string]interface{}
}

func (m *FindByPK) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Table, validation.Required),
		validation.Field(&m.Data, validation.Required),
	)
}

type FindOne struct {
	Table        string
	DisableCache bool
	Select       []string
	Where        []*Where
	Relations    []string
	Join         []*Join
	Offset       int
	OrderBy      []string
}

func (m *FindOne) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Table, validation.Required),
	)
}

type FindAll struct {
	Table        string
	DisableCache bool
	Select       []string
	Where        []*Where
	Relations    []string
	Joins        []*Join
	Limit        int
	Offset       int
	OrderBy      []string
}

func (m *FindAll) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Table, validation.Required),
	)
}

type Where struct {
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
	//The later task will not execute and get the result from the first task with the same lock key in the same time
	LockKey      string
	Table        string
	Data         interface{}
	DisableCache bool
}

func (m *UpdateByPK) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Table, validation.Required),
		validation.Field(&m.Data, validation.Required),
	)
}

type BulkUpdateByPK struct {
	// Lock key for concurrent insert operations
	//The later task will not execute and get the result from the first task with the same lock key in the same time
	LockKey      string
	Table        string
	Data         []interface{}
	DisableCache bool
}
