package dao

var (
	sqlExecutor *sqlExecutorImpl
)

type SQLExecutorInterface interface {
	BaseInterface
}
type sqlExecutorImpl struct {
	baseImpl
}

func SQLExecutor() SQLExecutorInterface {
	if sqlExecutor == nil {
		sqlExecutor = &sqlExecutorImpl{}
	}
	return sqlExecutor
}
