package config

import (
	"errors"
	"fmt"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"reflect"
)

type DBConfig struct {
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
	TableConfigs map[string]*TableConfig
	Models       map[string][]reflect.StructField // model with full columns and relations
	PtrModels    map[string][]reflect.StructField // model for cache and update action
	RefModels    map[string][]reflect.StructField // reference model for cache
}

const (
	DefaultModelType = "default_model_type"
	PtrModelType     = "ptr_model_type"
	RefModelType     = "ref_model_type"
)

func (s *ModelRegistry) GetModelBuilder(tbl, modelType string) []reflect.StructField {
	switch modelType {
	case DefaultModelType:
		return s.Models[tbl]
	case PtrModelType:
		return s.PtrModels[tbl]
	case RefModelType:
		return s.RefModels[tbl]
	}
	return []reflect.StructField{}
}

type TableConfig struct {
	Name              string
	PrimaryColumns    map[string]interface{}
	Columns           map[string]*Column
	RelationColumns   map[string]*RelationColumn
	SoftDeleteColumns map[string]string
}

func (s *ModelRegistry) GetTableConfig(tbl string) *TableConfig {
	return s.TableConfigs[tbl]
}

type RelationColumn struct {
	Name     string
	RefTable string
	Type     string
	Join     string
}

type Column struct {
	Type         string
	Nullable     bool
	IsPrimaryKey bool
	SoftDelete   bool
}

func NewDBConfig(req *requestmodel.UpsertConfig) (*DBConfig, error) {
	dbCfg := &DBConfig{
		Type:         req.DBConfig.Type,
		Host:         req.DBConfig.Host,
		Port:         req.DBConfig.Port,
		Username:     req.DBConfig.Username,
		Password:     req.DBConfig.Password,
		Database:     req.DBConfig.Database,
		MaxCons:      req.DBConfig.MaxCons,
		DisableCache: req.DisableCache,
		Debug:        req.Debug,
		ModelRegistry: &ModelRegistry{
			TableConfigs: map[string]*TableConfig{},
			Models:       map[string][]reflect.StructField{},
			PtrModels:    map[string][]reflect.StructField{},
			RefModels:    map[string][]reflect.StructField{},
		},
	}

	if req.DBConfig.TLS != nil {
		dbCfg.TLS = &TLS{
			RootCAPEM: req.DBConfig.TLS.RootCAPEM,
		}
	}

	for _, t := range req.TableConfigs {
		tblCfg := &TableConfig{
			Name:              t.Name,
			Columns:           map[string]*Column{},
			PrimaryColumns:    map[string]interface{}{},
			RelationColumns:   map[string]*RelationColumn{},
			SoftDeleteColumns: map[string]string{},
		}

		for _, col := range t.Columns {
			tblCfg.Columns[col.Name] = &Column{
				Type:         col.Type,
				Nullable:     col.Nullable,
				IsPrimaryKey: col.IsPrimaryKey,
				SoftDelete:   col.SoftDelete,
			}

			if col.IsPrimaryKey {
				tblCfg.PrimaryColumns[col.Name] = col.Name
			}

			if col.SoftDelete {
				tblCfg.SoftDeleteColumns[col.Name] = col.Name
			}
		}

		if len(tblCfg.PrimaryColumns) == 0 {
			return nil, errors.New(fmt.Sprintf("Table %s has no primary columns", tblCfg.Name))
		}

		for _, col := range t.RelationColumns {
			tblCfg.RelationColumns[col.Name] = &RelationColumn{
				Name:     col.Name,
				RefTable: col.RefTable,
				Type:     col.Type,
				Join:     col.Join,
			}
		}

		dbCfg.ModelRegistry.TableConfigs[t.Name] = tblCfg
	}

	return dbCfg, nil
}

func NewRedisConfiguration(req *requestmodel.UpsertConfig) *Redis {
	return &Redis{
		Host:     req.Redis.Host,
		Port:     req.Redis.Port,
		Username: req.Redis.Username,
		Password: req.Redis.Password,
		DB:       req.Redis.DB,
		PoolSize: req.Redis.PoolSize,
	}
}
