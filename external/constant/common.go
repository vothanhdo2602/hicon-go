package constant

import (
	"fmt"
	"github.com/uptrace/bun"
	"time"
)

const (
	MaintainerMail = "illusionless10@gmail.com"
)

func ErrorContactMaintainer(err error) error {
	return fmt.Errorf(`
		invalid data please contact to maintainer, mail: %s
		error: %v
	`, MaintainerMail, err.Error())
}

const (
	Expiry10Minutes = 10 * time.Minute
	Expiry5Minutes  = 5 * time.Minute
)

type ModelParams struct {
	Database            string
	Table               string
	DisableCache        bool
	LockKey             string
	ModeType            string
	CacheKey            string
	WhereAllWithDeleted bool
	Tx                  bun.IDB
}

const (
	BWOperationExec           = "exec"
	BWOperationBulkInsert     = "bulk_insert"
	BWOperationUpdateByPK     = "update_by_pk"
	BWOperationUpdateAll      = "update_all"
	BWOperationBulkUpdateByPK = "bulk_update_by_pk"
	BWOperationDeleteByPK     = "delete_by_pk"
)
