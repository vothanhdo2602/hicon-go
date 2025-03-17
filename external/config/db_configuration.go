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
	Models              map[string][]reflect.StructField
	PtrModels           map[string][]reflect.StructField
}

func (s *ModelRegistry) GetModelBuilder(tbl string, ptrModel bool) []reflect.StructField {
	if ptrModel {
		return s.PtrModels[tbl]
	}
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
		},
	}

	if req.DBConfiguration.TLS != nil {
		dbCfg.TLS = &TLS{
			RootCAPEM: req.DBConfiguration.TLS.RootCAPEM,
		}
	}

	for _, t := range req.TableConfigurations {
		tblCfg := &TableConfiguration{
			Name:                  t.Name,
			ColumnConfigs:         map[string]*ColumnConfig{},
			PrimaryColumns:        map[string]interface{}{},
			RelationColumnConfigs: map[string]*RelationColumnConfig{},
		}

		for _, col := range t.ColumnConfigs {
			tblCfg.ColumnConfigs[col.Name] = &ColumnConfig{
				Type:         col.Type,
				Nullable:     col.Nullable,
				IsPrimaryKey: col.IsPrimaryKey,
			}

			if col.IsPrimaryKey {
				tblCfg.PrimaryColumns[col.Name] = col.Name
			}
		}

		if len(tblCfg.PrimaryColumns) == 0 {
			return nil, errors.New(fmt.Sprintf("Table %s has no primary columns", tblCfg.Name))
		}

		for _, col := range t.RelationColumnConfigs {
			tblCfg.RelationColumnConfigs[col.Name] = &RelationColumnConfig{
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
