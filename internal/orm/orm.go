package orm

import (
	"context"
	"crypto/tls"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/vothanhdo2602/hicon/external/config"
	"github.com/vothanhdo2602/hicon/external/constant"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"sync"
)

var (
	db *bun.DB
)

func Init(ctx context.Context, wg *sync.WaitGroup, c *config.DBConfiguration) error {
	var (
		logger  = log.WithCtx(ctx)
		sslMode = "disable"
		newDB   *bun.DB
	)

	if wg != nil {
		defer wg.Done()
	}

	addr := fmt.Sprintf("%s:%v", c.Host, c.Port)

	if c.TLS != nil {
		sslMode = "verify-full"
	}

	switch c.Type {
	case constant.DBMysql:
		sqlDB, err := sql.Open(constant.DBMysql, fmt.Sprintf("%s:%s@tcp(%s)/%s", c.Username, c.Password, addr, c.Database))
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		sqlDB.SetMaxIdleConns(c.MaxCons)
		sqlDB.SetMaxOpenConns(c.MaxCons)

		newDB = bun.NewDB(sqlDB, mysqldialect.New())
	case constant.DBPostgres:
		cfg, err := pgxpool.ParseConfig(fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=%s", c.Type, c.Username, c.Password, addr, c.Database, sslMode))
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		cfg.MinConns = int32(c.MaxCons)
		cfg.MaxConns = int32(c.MaxCons)
		cfg.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

		pool, err := pgxpool.NewWithConfig(ctx, cfg)
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		if c.TLS != nil {
			cfg.ConnConfig.TLSConfig.InsecureSkipVerify = false
			cfg.ConnConfig.TLSConfig.MinVersion = tls.VersionTLS12
			cfg.ConnConfig.TLSConfig.RootCAs.AppendCertsFromPEM([]byte(c.TLS.RootCAPEM))
		}
		sqlDB := stdlib.OpenDBFromPool(pool)
		newDB = bun.NewDB(sqlDB, pgdialect.New(), bun.WithDiscardUnknownColumns())
	default:
		return errors.New("database type not supported")
	}

	if err := newDB.DB.Ping(); err != nil {
		db = nil
		return err
	}

	logger.Info(fmt.Sprintf("⚡️[%s]: connected to %s", c.Type, addr))

	newDB.AddQueryHook(
		bundebug.NewQueryHook(bundebug.WithVerbose(c.Debug)),
	)

	// open telemetry tracing
	//newDB.AddQueryHook(bunotel.NewQueryHook(bunotel.WithDBName(c.Database)))

	//newDB.AddQueryHook(&QueryHook{})

	// reassign when done setting
	// for high availability
	db = newDB

	return nil
}

func GetDB() bun.IDB {
	return db
}
