package grpctil

import (
	"github.com/goccy/go-json"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/model/responsemodel"
	"github.com/vothanhdo2602/hicon/hicon-sm/sqlexecutor"
	"google.golang.org/protobuf/proto"
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

func ConvertInterfaceToPbAny(v interface{}) (*anypb.Any, error) {
	anyValue := &anypb.Any{}
	bytes, _ := json.Marshal(v)
	bytesValue := &wrappers.BytesValue{
		Value: bytes,
	}
	err := anypb.MarshalFrom(anyValue, bytesValue, proto.MarshalOptions{})
	return anyValue, err
}

func ConvertSliceAnyToPbAnySlice(input []interface{}) ([]*anypb.Any, error) {
	result := make([]*anypb.Any, 0, len(input))

	for _, item := range input {
		d, err := ConvertInterfaceToPbAny(item)
		if err != nil {
			return nil, err
		}
		result = append(result, d)
	}

	return result, nil
}

func ConvertPbAnyToInterface(anyValue *anypb.Any) (interface{}, error) {
	var value interface{}
	bytesValue := &wrappers.BytesValue{}
	err := anypb.UnmarshalTo(anyValue, bytesValue, proto.UnmarshalOptions{})
	if err != nil {
		return value, err
	}
	uErr := json.Unmarshal(bytesValue.Value, &value)
	if uErr != nil {
		return value, uErr
	}
	return value, nil
}

func ConvertSlicePbAnyToSliceInterface(anyValue []*anypb.Any) ([]interface{}, error) {
	var values []interface{}
	for _, v := range anyValue {
		r, err := ConvertPbAnyToInterface(v)
		if err != nil {
			return values, err
		}
		values = append(values, r)
	}

	return values, nil
}

func ConvertQueryWithArgsProtoToRequestModel(protoWhere *sqlexecutor.QueryWithArgs) (*requestmodel.QueryWithArgs, error) {
	if protoWhere == nil {
		return nil, nil
	}

	whereGo := &requestmodel.QueryWithArgs{
		Query: protoWhere.Query,
		Args:  make([]interface{}, 0, len(protoWhere.Args)),
	}

	for _, arg := range protoWhere.Args {
		val, _ := ConvertPbAnyToInterface(arg)
		whereGo.Args = append(whereGo.Args, val)
	}

	return whereGo, nil
}

func ConvertSliceQueryWithArgsProtoToRequestModel(protoWhere []*sqlexecutor.QueryWithArgs) ([]*requestmodel.QueryWithArgs, error) {
	var r []*requestmodel.QueryWithArgs
	for _, w := range protoWhere {
		newWhere, err := ConvertQueryWithArgsProtoToRequestModel(w)
		if err != nil {
			return nil, err
		}

		r = append(r, newWhere)
	}

	return r, nil
}

func ConvertPbAnyToInterfaceWithType[T any](anyValue *anypb.Any) (T, error) {
	var value T
	bytesValue := &wrappers.BytesValue{}
	err := anypb.UnmarshalTo(anyValue, bytesValue, proto.UnmarshalOptions{})
	if err != nil {
		return value, err
	}
	uErr := json.Unmarshal(bytesValue.Value, &value)
	if uErr != nil {
		return value, uErr
	}
	return value, nil
}
