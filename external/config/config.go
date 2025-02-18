package config

import (
	"time"
)

type ENV struct {
	DB struct {
		DBConfig *DBConfig
		Redis    struct {
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

type DBConfig struct {
	Type         string
	Host         string
	Port         int
	Username     string
	Password     string
	Database     string
	MaxCons      int
	TLS          *TLS
	Debug        bool
	DisableCache bool
	TableConfigs []*TableConfig
	TableStructs map[string]interface{}
}

type TableConfig struct {
	Name                  string
	PrimaryColumns        []string
	ColumnConfigs         map[string]*ColumnConfig
	RelationColumnConfigs map[string]*RelationColumnConfigs
}

type RelationColumnConfigs struct {
	RefTable string
	Type     string
}

type ColumnConfig struct {
	Type         string
	Nullable     bool
	IsPrimaryKey bool
}

var env *ENV

func GetENV() *ENV {
	return env
}

func SetDBConfig(cfg *DBConfig) {
	env.DB.DBConfig = cfg
}
