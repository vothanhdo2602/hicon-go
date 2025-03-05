package natstil

type BaseResponse[T any] struct {
	Data    T
	Status  int
	Message string
	Shared  bool
	Success bool
}
