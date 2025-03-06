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
