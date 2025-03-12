package config

import (
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"reflect"
	"strings"
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
		Host string
		Port int
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
	Entities            map[string]interface{}
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
	Join     string
}

type ColumnConfig struct {
	Type         string
	Nullable     bool
	IsPrimaryKey bool
}

var env ENV

func Init() {
	env.Common.Host = "0.0.0.0"
	env.Common.Port = 7979
}

func GetAddr() string {
	return fmt.Sprintf("%s:%d", env.Common.Host, env.Common.Port)
}

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

func ModelWithSelectFields(table string, fields []string) interface{} {
	var (
		reflectFields   []reflect.StructField
		registeredModel = env.DB.DBConfiguration.ModelRegistry.Models[table]
	)

	for _, field := range fields {
		for _, rmField := range registeredModel {
			if strings.ToLower(rmField.Name) == strings.ToLower(field) {
				reflectFields = append(reflectFields, rmField)
				break
			}
		}
	}

	return reflect.New(reflect.StructOf(reflectFields)).Interface()
}

func TransformModel(table string, fields []string, data interface{}) (interface{}, error) {
	var (
		newModel = ModelWithSelectFields(table, fields)
	)

	bytesModel, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytesModel, &newModel)
	if err != nil {
		return nil, err
	}

	return newModel, nil
}

func (s *DBConfiguration) GetDatabaseName() string {
	return env.DB.DBConfiguration.Database
}
