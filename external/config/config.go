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
	Credential *Credential
	DB         DB
	Common     struct {
		Host string
		Port int
	}
}

type Credential struct {
	AccessKey string
	SecretKey string
	ExpireAt  *time.Time
	IsValid   bool
}

type DB struct {
	DBConfig *DBConfig
	Redis    *Redis
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

func (s ENV) GetDB() DB {
	return s.DB
}

func (s DB) GetDBConfig() *DBConfig {
	return GetENV().GetDB().DBConfig
}

func SetDBConfig(cfg *DBConfig) {
	env.DB.DBConfig = cfg
}

func SetRedisConfiguration(cfg *Redis) {
	env.DB.Redis = cfg
}

func ConfigurationUpdated() error {
	//if env.Credential == nil || !env.Credential.IsValid {
	//	return errors.New("invalid credentials, please check your credentials and reconnect or contact to maintainer, gmail: illusionless10@gmail.com")
	//}

	if env.DB.DBConfig != nil {
		return nil
	}
	return errors.New("not found configuration, please send your configuration with func UpsertConfig")
}

func GetModelRegistry() *ModelRegistry {
	return GetENV().GetDB().DBConfig.ModelRegistry
}

func (s *ModelRegistry) GetNewModel(name string) interface{} {
	return reflect.New(reflect.StructOf(s.Models[name])).Interface()
}

func (s *ModelRegistry) GetNewSliceModel(name string) interface{} {
	return reflect.New(reflect.SliceOf(reflect.StructOf(s.Models[name]))).Interface()
}

func ModelWithSelectFields(table string, fields []string, modelType string) interface{} {
	var (
		reflectFields   []reflect.StructField
		registeredModel = GetModelRegistry().GetModelBuilder(table, modelType)
	)

	if len(fields) == 0 {
		return reflect.New(reflect.StructOf(registeredModel)).Interface()
	}

	for _, field := range fields {
		for _, registeredField := range registeredModel {
			if strings.ToLower(registeredField.Name) == strings.ToLower(field) {
				reflectFields = append(reflectFields, registeredField)
				break
			}
		}
	}

	return reflect.New(reflect.StructOf(reflectFields)).Interface()
}

func TransformModel(table string, fields []string, data interface{}, modelType string) (interface{}, error) {
	var (
		newModel = ModelWithSelectFields(table, fields, modelType)
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

func ModelsWithSelectFields(table string, fields []string, modelType string) interface{} {
	var (
		reflectFields   []reflect.StructField
		registeredModel = GetModelRegistry().GetModelBuilder(table, modelType)
	)

	if len(fields) == 0 {
		return reflect.New(reflect.SliceOf(reflect.StructOf(registeredModel))).Interface()
	}

	for _, field := range fields {
		for _, registeredField := range registeredModel {
			if strings.ToLower(registeredField.Name) == strings.ToLower(field) {
				reflectFields = append(reflectFields, registeredField)
				break
			}
		}
	}

	return reflect.New(reflect.SliceOf(reflect.StructOf(reflectFields))).Interface()
}

func TransformModels(table string, fields []string, data interface{}, modelType string) (interface{}, error) {
	var (
		newModels = ModelsWithSelectFields(table, fields, modelType)
	)

	bytesModel, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytesModel, &newModels)
	if err != nil {
		return nil, err
	}

	return newModels, nil
}

func (s *DBConfig) GetDatabaseName() string {
	return GetENV().GetDB().GetDBConfig().Database
}
