// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: hicon-sm/sql_executor.proto

package sqlexecutor

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_hicon_sm_sql_executor_proto protoreflect.FileDescriptor

var file_hicon_sm_sql_executor_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2d, 0x73, 0x6d, 0x2f, 0x73, 0x71, 0x6c, 0x5f, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x68,
	0x69, 0x63, 0x6f, 0x6e, 0x2d, 0x73, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xb7, 0x04, 0x0a, 0x0b, 0x53, 0x51, 0x4c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x12, 0x52, 0x0a, 0x13, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x68, 0x69, 0x63, 0x6f,
	0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x68, 0x69, 0x63,
	0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x1d, 0x2e, 0x68, 0x69,
	0x63, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x73, 0x1a, 0x18, 0x2e, 0x68, 0x69, 0x63,
	0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e,
	0x65, 0x12, 0x13, 0x2e, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x1a, 0x18, 0x2e, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x13, 0x2e,
	0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x1a, 0x18, 0x2e, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x04, 0x45, 0x78, 0x65, 0x63, 0x12, 0x10, 0x2e, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x1a, 0x18, 0x2e, 0x68, 0x69, 0x63, 0x6f, 0x6e,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0a, 0x42, 0x75, 0x6c, 0x6b, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x12, 0x16, 0x2e, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x42, 0x75, 0x6c, 0x6b, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x1a, 0x18, 0x2e, 0x68, 0x69, 0x63,
	0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x79, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x1f, 0x2e,
	0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x79, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x73, 0x1a, 0x18,
	0x2e, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0a, 0x42, 0x75,
	0x6c, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x68, 0x69, 0x63, 0x6f, 0x6e,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x1a, 0x18, 0x2e, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d,
	0x2e, 0x2f, 0x73, 0x71, 0x6c, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_hicon_sm_sql_executor_proto_goTypes = []any{
	(*UpsertConfiguration)(nil),     // 0: hicon.type.UpsertConfiguration
	(*FindByPrimaryKeys)(nil),       // 1: hicon.type.FindByPrimaryKeys
	(*FindOne)(nil),                 // 2: hicon.type.FindOne
	(*FindAll)(nil),                 // 3: hicon.type.FindAll
	(*Exec)(nil),                    // 4: hicon.type.Exec
	(*BulkInsert)(nil),              // 5: hicon.type.BulkInsert
	(*UpdateByPrimaryKeys)(nil),     // 6: hicon.type.UpdateByPrimaryKeys
	(*BulkUpdateByPrimaryKeys)(nil), // 7: hicon.type.BulkUpdateByPrimaryKeys
	(*BaseResponse)(nil),            // 8: hicon.type.BaseResponse
}
var file_hicon_sm_sql_executor_proto_depIdxs = []int32{
	0, // 0: SQLExecutor.UpsertConfiguration:input_type -> hicon.type.UpsertConfiguration
	1, // 1: SQLExecutor.FindByPrimaryKeys:input_type -> hicon.type.FindByPrimaryKeys
	2, // 2: SQLExecutor.FindOne:input_type -> hicon.type.FindOne
	3, // 3: SQLExecutor.FindAll:input_type -> hicon.type.FindAll
	4, // 4: SQLExecutor.Exec:input_type -> hicon.type.Exec
	5, // 5: SQLExecutor.BulkInsert:input_type -> hicon.type.BulkInsert
	6, // 6: SQLExecutor.UpdateByPrimaryKeys:input_type -> hicon.type.UpdateByPrimaryKeys
	7, // 7: SQLExecutor.BulkUpdateByPrimaryKeys:input_type -> hicon.type.BulkUpdateByPrimaryKeys
	8, // 8: SQLExecutor.UpsertConfiguration:output_type -> hicon.type.BaseResponse
	8, // 9: SQLExecutor.FindByPrimaryKeys:output_type -> hicon.type.BaseResponse
	8, // 10: SQLExecutor.FindOne:output_type -> hicon.type.BaseResponse
	8, // 11: SQLExecutor.FindAll:output_type -> hicon.type.BaseResponse
	8, // 12: SQLExecutor.Exec:output_type -> hicon.type.BaseResponse
	8, // 13: SQLExecutor.BulkInsert:output_type -> hicon.type.BaseResponse
	8, // 14: SQLExecutor.UpdateByPrimaryKeys:output_type -> hicon.type.BaseResponse
	8, // 15: SQLExecutor.BulkUpdateByPrimaryKeys:output_type -> hicon.type.BaseResponse
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hicon_sm_sql_executor_proto_init() }
func file_hicon_sm_sql_executor_proto_init() {
	if File_hicon_sm_sql_executor_proto != nil {
		return
	}
	file_hicon_sm_type_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_hicon_sm_sql_executor_proto_rawDesc), len(file_hicon_sm_sql_executor_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hicon_sm_sql_executor_proto_goTypes,
		DependencyIndexes: file_hicon_sm_sql_executor_proto_depIdxs,
	}.Build()
	File_hicon_sm_sql_executor_proto = out.File
	file_hicon_sm_sql_executor_proto_goTypes = nil
	file_hicon_sm_sql_executor_proto_depIdxs = nil
}
