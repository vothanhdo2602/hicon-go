package requestmodel

type BaseRequest struct {
	Headers map[string]string `json:"headers"`
	Body    interface{}
}

type BaseRequestWithType[T any] struct {
	Headers map[string]string `json:"headers"`
	Body    *T
}
