package config

import (
	"errors"
	"reflect"
	"time"
)

type ENV struct {
	DB struct {
		DBConfiguration *DBConfiguration
		Redis           struct {
			Host     string
			Port     int
			Username string
			Password string
			DB       string
		}
		Nats struct {
			URL            string
			User           string
			Password       string
			TLS            *TLS
			RequestTimeout time.Duration
			StreamName     string
		}
	}
	Common struct {
		Server struct {
			Host string
			Port int
		}
		Admin struct {
			Domain string
			Host   string
			Port   int
		}
	}
}

type TLS struct {
	CertPEM       string
	PrivateKeyPEM string
	RootCAPEM     string
}

type DBConfiguration struct {
	Type          string
	Host          string
	Port          int
	Username      string
	Password      string
	Database      string
	MaxCons       int
	TLS           *TLS
	Debug         bool
	DisableCache  bool
	ModelRegistry *ModelRegistry
}

type ModelRegistry struct {
	TableConfigurations map[string]*TableConfiguration
	Models              map[string][]reflect.StructField
}

func (s *ModelRegistry) GetModelBuilder(tbl string) []reflect.StructField {
	return s.Models[tbl]
}

type TableConfiguration struct {
	Name                  string
	PrimaryColumns        map[string]interface{}
	ColumnConfigs         map[string]*ColumnConfig
	RelationColumnConfigs map[string]*RelationColumnConfig
}

type RelationColumnConfig struct {
	Name     string
	RefTable string
	Type     string
}

type ColumnConfig struct {
	Type         string
	Nullable     bool
	IsPrimaryKey bool
}

var env ENV

func GetENV() ENV {
	return env
}

func SetDBConfiguration(cfg *DBConfiguration) {
	env.DB.DBConfiguration = cfg
}

func ConfigurationUpdated() error {
	if env.DB.DBConfiguration != nil {
		return nil
	}
	return errors.New("no configuration")
}

func GetModelRegistry() *ModelRegistry {
	return env.DB.DBConfiguration.ModelRegistry
}

func (s *ModelRegistry) GetNewModel(name string) interface{} {
	return reflect.New(reflect.StructOf(s.Models[name])).Interface()
}
