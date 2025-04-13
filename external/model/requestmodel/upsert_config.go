package requestmodel

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/vothanhdo2602/hicon/hicon-sm/constant"
)

type Credential struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

func (s *Credential) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.AccessKey, validation.Required),
		validation.Field(&s.SecretKey, validation.Required),
	)
}

type UpsertConfig struct {
	DBConfig     *DBConfig
	Redis        *Redis
	TableConfigs []*TableConfig
	Debug        bool
	DisableCache bool
}

type DBConfig struct {
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

type TableConfig struct {
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

func (s *UpsertConfig) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.DBConfig, validation.Required),
	)
}

func (s *DBConfig) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Type, validation.In(constant.DBPostgres, constant.DBMysql)),
		validation.Field(&s.Host, validation.Required),
		validation.Field(&s.Port, validation.Min(1), validation.Max(65535)),
		validation.Field(&s.MaxCons, validation.Min(1)),
	)
}
