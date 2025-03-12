package requestmodel

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/vothanhdo2602/hicon/external/constant"
)

type UpsertConfiguration struct {
	DBConfiguration     *DBConfiguration
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

type TLS struct {
	CertPEM       string
	PrivateKeyPEM string
	RootCAPEM     string
}

type TableConfiguration struct {
	Name                  string
	ColumnConfigs         []*ColumnConfig
	RelationColumnConfigs []*RelationColumnConfigs
}

type ColumnConfig struct {
	Name         string
	Type         string
	Nullable     bool
	IsPrimaryKey bool
}

type RelationColumnConfigs struct {
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

type FindByPrimaryKeys struct {
	Table        string
	DisableCache bool
	Select       []string
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
	Where        []*Where
	Relations    []string
	Offset       int
	OrderBy      []string
}

type FindAll struct {
	Table        string
	DisableCache bool
	Select       []string
	Where        []*Where
	Relations    []string
	Limit        int
	Offset       int
	OrderBy      []string
}

type Where struct {
	Condition string
	Args      []interface{}
}

type Relation struct {
	Name string
}

func (m *FindOne) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Table, validation.Required),
	)
}
