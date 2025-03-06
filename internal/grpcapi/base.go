package grpcapi

import (
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon/external/model/responsemodel"
	"github.com/vothanhdo2602/hicon/hicon-sm/sqlexecutor"
	"google.golang.org/protobuf/types/known/anypb"
)

func NewResponse(r *responsemodel.BaseResponse) *sqlexecutor.BaseResponse {
	data, _ := json.Marshal(r.Data)
	return &sqlexecutor.BaseResponse{
		Shared:  r.Shared,
		Status:  int32(r.Status),
		Success: r.Success,
		Message: r.Message,
		Data: &anypb.Any{
			Value: data,
		},
	}
}

func InterfaceToAnyPb(data interface{}) *anypb.Any {
	bytes, _ := json.Marshal(data)
	return &anypb.Any{
		Value: bytes,
	}
}

func AnyMapToInterfaceMap(anyMap map[string]*anypb.Any) map[string]interface{} {
	result := map[string]interface{}{}

	for key, anyValue := range anyMap {
		var value interface{}
		_ = json.Unmarshal(anyValue.GetValue(), &value)
		result[key] = value
	}

	return result
}
