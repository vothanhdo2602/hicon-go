package orm

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/vothanhdo2602/hicon/external/util/log"
)

func BeginTx(ctx context.Context) (bun.Tx, error) {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.WithCtx(ctx).Error(err.Error())
	}
	return tx, err
}

func HandleTxErr(ctx context.Context, tx bun.Tx, err error) {
	var (
		logger = log.WithCtx(ctx)
	)

	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			logger.Error(rollbackErr.Error())
		}
		return
	}

	if commitErr := tx.Commit(); commitErr != nil {
		logger.Error(commitErr.Error())
	}
}
