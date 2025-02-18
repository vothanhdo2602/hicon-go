package natstil

type IResponse[T any] struct {
	Data    T
	Status  int
	Message string
	Success bool
}
