package config

import (
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"reflect"
	"strings"
)

type ENV struct {
	DB     DB
	Common struct {
		Host string
		Port int
	}
}

type DB struct {
	DBConfiguration *DBConfiguration
	Redis           *Redis
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

func (s DB) GetDBConfiguration() *DBConfiguration {
	return GetENV().GetDB().DBConfiguration
}

func SetDBConfiguration(cfg *DBConfiguration) {
	env.DB.DBConfiguration = cfg
}

func SetRedisConfiguration(cfg *Redis) {
	env.DB.Redis = cfg
}

func ConfigurationUpdated() error {
	if env.DB.DBConfiguration != nil {
		return nil
	}
	return errors.New("no configuration")
}

func GetModelRegistry() *ModelRegistry {
	return GetENV().GetDB().DBConfiguration.ModelRegistry
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

func (s *DBConfiguration) GetDatabaseName() string {
	return GetENV().GetDB().GetDBConfiguration().Database
}
