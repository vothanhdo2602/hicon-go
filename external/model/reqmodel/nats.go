package reqmodel

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/vothanhdo2602/hicon/external/constant"
	"github.com/vothanhdo2602/hicon/internal/orm"
)

type UpdateConfig struct {
	DBConfig struct {
		Type     string
		Host     string
		Port     int
		Username string
		Password string
		Database string
		MaxCons  int
		TLS      *TLS
	}
	TableConfigs []*TableConfig
	Debug        bool
	DisableCache bool
}

type TLS struct {
	CertPEM       string
	PrivateKeyPEM string
	RootCAPEM     string
}

type TableConfig struct {
	Name                  string
	ColumnConfigs         map[string]*ColumnConfig
	RelationColumnConfigs map[string]*RelationColumnConfigs
}

type ColumnConfig struct {
	Name         string
	Type         string
	Nullable     bool
	IsPrimaryKey bool
}

type RelationColumnConfigs struct {
	RefTable string
	Type     orm.RelationType
}

func (m UpdateConfig) Validate() error {
	return validation.ValidateStruct(
		&m,
		validation.Field(&m.DBConfig.Type, validation.In(constant.DBPostgres, constant.DBMysql)),
		validation.Field(&m.DBConfig.Port, validation.Min(1), validation.Max(65535)),
		validation.Field(&m.DBConfig.MaxCons, validation.Min(1)),
	)
}
