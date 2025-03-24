package config

import (
	"errors"
	"fmt"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"reflect"
)

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
	Models              map[string][]reflect.StructField // model with full columns and relations
	PtrModels           map[string][]reflect.StructField // model for cache and update action
	RefModels           map[string][]reflect.StructField // reference model for cache
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

type TableConfiguration struct {
	Name              string
	PrimaryColumns    map[string]interface{}
	ColumnConfigs     map[string]*ColumnConfig
	RelationColumns   map[string]*RelationColumn
	SoftDeleteColumns map[string]string
}

func (s *ModelRegistry) GetTableConfiguration(tbl string) *TableConfiguration {
	return s.TableConfigurations[tbl]
}

type RelationColumn struct {
	Name     string
	RefTable string
	Type     string
	Join     string
}

type ColumnConfig struct {
	Type         string
	Nullable     bool
	IsPrimaryKey bool
	SoftDelete   bool
}

func NewDBConfiguration(req *requestmodel.UpsertConfiguration) (*DBConfiguration, error) {
	dbCfg := &DBConfiguration{
		Type:         req.DBConfiguration.Type,
		Host:         req.DBConfiguration.Host,
		Port:         req.DBConfiguration.Port,
		Username:     req.DBConfiguration.Username,
		Password:     req.DBConfiguration.Password,
		Database:     req.DBConfiguration.Database,
		MaxCons:      req.DBConfiguration.MaxCons,
		DisableCache: req.DisableCache,
		Debug:        req.Debug,
		ModelRegistry: &ModelRegistry{
			TableConfigurations: map[string]*TableConfiguration{},
			Models:              map[string][]reflect.StructField{},
			PtrModels:           map[string][]reflect.StructField{},
			RefModels:           map[string][]reflect.StructField{},
		},
	}

	if req.DBConfiguration.TLS != nil {
		dbCfg.TLS = &TLS{
			RootCAPEM: req.DBConfiguration.TLS.RootCAPEM,
		}
	}

	for _, t := range req.TableConfigurations {
		tblCfg := &TableConfiguration{
			Name:              t.Name,
			ColumnConfigs:     map[string]*ColumnConfig{},
			PrimaryColumns:    map[string]interface{}{},
			RelationColumns:   map[string]*RelationColumn{},
			SoftDeleteColumns: map[string]string{},
		}

		for _, col := range t.ColumnConfigs {
			tblCfg.ColumnConfigs[col.Name] = &ColumnConfig{
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

		dbCfg.ModelRegistry.TableConfigurations[t.Name] = tblCfg
	}

	return dbCfg, nil
}

func NewRedisConfiguration(req *requestmodel.UpsertConfiguration) *Redis {
	return &Redis{
		Host:     req.Redis.Host,
		Port:     req.Redis.Port,
		Username: req.Redis.Username,
		Password: req.Redis.Password,
		DB:       req.Redis.DB,
		PoolSize: req.Redis.PoolSize,
	}
}
