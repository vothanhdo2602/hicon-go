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
	"go.uber.org/zap"
	"sync"
)

var (
	db *bun.DB
)

func Init(ctx context.Context, wg *sync.WaitGroup) error {
	var (
		logger  = log.WithCtx(ctx)
		sslMode = "disable"
	)

	if wg != nil {
		defer wg.Done()
	}

	c := config.GetENV().DB.DBConfig
	addr := fmt.Sprintf("%s:%v", c.Host, c.Port)

	if c.TLS != nil {
		sslMode = "verify-full"
	}

	switch c.Type {
	case constant.DBMysql:
		sqlDB, err := sql.Open(constant.DBMysql, fmt.Sprintf("%s:%s@tcp(%s)/%s", c.Username, c.Password, addr, c.Database))
		if err != nil {
			panic(err)
		}

		db = bun.NewDB(sqlDB, mysqldialect.New())
	case constant.DBPostgres:
		cfg, err := pgxpool.ParseConfig(fmt.Sprintf("%s://%s:%d@%s/%s?sslmode=%s", constant.DBPostgres, c.Host, c.Port, addr, c.Database, sslMode))
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		cfg.MinConns = int32(c.MaxCons)
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
		db = bun.NewDB(sqlDB, pgdialect.New(), bun.WithDiscardUnknownColumns())
	default:
		return errors.New("database type not supported")
	}

	if err := db.DB.Ping(); err != nil {
		logger.Error(err.Error(), zap.String("postgres", fmt.Sprintf("⚡️[postgres]: addr: %s", addr)))
		return err
	}

	logger.Info(fmt.Sprintf("⚡️[postgres]: connected to %s", addr))

	db.AddQueryHook(
		bundebug.NewQueryHook(bundebug.WithVerbose(c.Debug)),
	)

	// open telemetry tracing
	//db.AddQueryHook(bunotel.NewQueryHook(bunotel.WithDBName(c.Database)))

	return nil
}

func GetDB() bun.IDB {
	return db
}
