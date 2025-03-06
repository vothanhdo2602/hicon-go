package responsemodel

import (
	"net/http"
)

type BaseResponse struct {
	Data    interface{}
	Status  int
	Message string
	Shared  bool
	Success bool
}

func R200(data any, shared bool) *BaseResponse {
	return &BaseResponse{
		Data:    data,
		Status:  http.StatusOK,
		Shared:  shared,
		Success: true,
	}
}

func R400(msg string) *BaseResponse {
	return &BaseResponse{
		Message: msg,
		Status:  http.StatusBadRequest,
		Success: false,
	}
}
