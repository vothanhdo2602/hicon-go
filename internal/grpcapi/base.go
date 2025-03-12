package grpcapi

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
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

func ProtoValueToInterface(protoValue *sqlexecutor.Value) interface{} {
	if protoValue == nil {
		return nil
	}

	switch k := protoValue.Kind.(type) {
	case *sqlexecutor.Value_StringValue:
		return k.StringValue
	case *sqlexecutor.Value_IntValue:
		return k.IntValue
	case *sqlexecutor.Value_FloatValue:
		return k.FloatValue
	case *sqlexecutor.Value_BoolValue:
		return k.BoolValue
	case *sqlexecutor.Value_BytesValue:
		return k.BytesValue
	default:
		return nil
	}
}

func InterfaceToProtoValue(val interface{}) (*sqlexecutor.Value, error) {
	if val == nil {
		return nil, nil
	}

	protoValue := &sqlexecutor.Value{}

	switch v := val.(type) {
	case string:
		protoValue.Kind = &sqlexecutor.Value_StringValue{StringValue: v}
	case int:
		protoValue.Kind = &sqlexecutor.Value_IntValue{IntValue: int64(v)}
	case int32:
		protoValue.Kind = &sqlexecutor.Value_IntValue{IntValue: int64(v)}
	case int64:
		protoValue.Kind = &sqlexecutor.Value_IntValue{IntValue: v}
	case float32:
		protoValue.Kind = &sqlexecutor.Value_FloatValue{FloatValue: float64(v)}
	case float64:
		protoValue.Kind = &sqlexecutor.Value_FloatValue{FloatValue: v}
	case bool:
		protoValue.Kind = &sqlexecutor.Value_BoolValue{BoolValue: v}
	case []byte:
		protoValue.Kind = &sqlexecutor.Value_BytesValue{BytesValue: v}
	default:
		return nil, fmt.Errorf("unsupported type: %T", val)
	}

	return protoValue, nil
}

func ConvertWhereProtoToGo(protoWhere *sqlexecutor.Where) (*requestmodel.Where, error) {
	if protoWhere == nil {
		return nil, nil
	}

	whereGo := &requestmodel.Where{
		Condition: protoWhere.Condition,
		Args:      make([]interface{}, 0, len(protoWhere.Args)),
	}

	for _, arg := range protoWhere.Args {
		val := ProtoValueToInterface(arg)
		whereGo.Args = append(whereGo.Args, val)
	}

	return whereGo, nil
}
