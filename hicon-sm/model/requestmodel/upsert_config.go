package requestmodel

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/vothanhdo2602/hicon/external/constant"
)

// Option is a functional configuration option type
type Option func(config *UpsertConfig)

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

// NewUpsertConfig New creates a new NewUpsertConfig with default settings
func (s *HiconClient) NewUpsertConfig() *UpsertConfig {
	return &UpsertConfig{}
}

// WithDebug enables or disables debug mode
func WithDebug(debug bool) Option {
	return func(c *UpsertConfig) {
		c.Debug = debug
	}
}

// WithDisableCache toggles cache functionality
func WithDisableCache(disable bool) Option {
	return func(c *UpsertConfig) {
		c.DisableCache = disable
	}
}

// WithDBConfig WithDB Database configuration methods
func WithDBConfig(dbConfig *DBConfig) Option {
	return func(c *UpsertConfig) {
		c.DBConfig = dbConfig
	}
}

// WithRedis Redis configuration methods
func WithRedis(redis *Redis) Option {
	return func(c *UpsertConfig) {
		c.Redis = redis
	}
}

// WithTable adds a table configuration
func WithTable(name string, columns []*Column, relations []*RelationColumn) Option {
	return func(c *UpsertConfig) {
		c.TableConfigs = append(c.TableConfigs, &TableConfig{
			Name:            name,
			Columns:         columns,
			RelationColumns: relations,
		})
	}
}

// Build finalizes the configuration
func (s *UpsertConfig) Build(opts ...Option) *UpsertConfig {
	for _, opt := range opts {
		opt(s)
	}
	return s
}
