package dao

type SQLExecutorInterface[T any] interface {
	BaseInterface[T]
}
type sqlExecutorImpl[T any] struct {
	baseImpl[T]
}

func SQLExecutor[T any]() SQLExecutorInterface[T] {
	return &sqlExecutorImpl[T]{}
}
