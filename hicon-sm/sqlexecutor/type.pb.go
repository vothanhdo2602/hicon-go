// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: hicon-sm/type.proto

package sqlexecutor

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BaseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Shared        bool                   `protobuf:"varint,1,opt,name=shared,proto3" json:"shared,omitempty"`
	Status        int32                  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Success       bool                   `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Data          *anypb.Any             `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	mi := &file_hicon_sm_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hicon_sm_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponse.ProtoReflect.Descriptor instead.
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return file_hicon_sm_type_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResponse) GetShared() bool {
	if x != nil {
		return x.Shared
	}
	return false
}

func (x *BaseResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *BaseResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BaseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BaseResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

// TLS Configuration for Secure Connections
type TLS struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CertPem       string                 `protobuf:"bytes,1,opt,name=cert_pem,json=certPem,proto3" json:"cert_pem,omitempty"`
	PrivateKeyPem string                 `protobuf:"bytes,2,opt,name=private_key_pem,json=privateKeyPem,proto3" json:"private_key_pem,omitempty"`
	RootCaPem     string                 `protobuf:"bytes,3,opt,name=root_ca_pem,json=rootCaPem,proto3" json:"root_ca_pem,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TLS) Reset() {
	*x = TLS{}
	mi := &file_hicon_sm_type_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TLS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLS) ProtoMessage() {}

func (x *TLS) ProtoReflect() protoreflect.Message {
	mi := &file_hicon_sm_type_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLS.ProtoReflect.Descriptor instead.
func (*TLS) Descriptor() ([]byte, []int) {
	return file_hicon_sm_type_proto_rawDescGZIP(), []int{1}
}

func (x *TLS) GetCertPem() string {
	if x != nil {
		return x.CertPem
	}
	return ""
}

func (x *TLS) GetPrivateKeyPem() string {
	if x != nil {
		return x.PrivateKeyPem
	}
	return ""
}

func (x *TLS) GetRootCaPem() string {
	if x != nil {
		return x.RootCaPem
	}
	return ""
}

// Database Configuration
type DBConfiguration struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Host          string                 `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port          int32                  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Username      string                 `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Database      string                 `protobuf:"bytes,6,opt,name=database,proto3" json:"database,omitempty"`
	MaxCons       int32                  `protobuf:"varint,7,opt,name=max_cons,json=maxCons,proto3" json:"max_cons,omitempty"`
	TLS           *TLS                   `protobuf:"bytes,8,opt,name=TLS,proto3" json:"TLS,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DBConfiguration) Reset() {
	*x = DBConfiguration{}
	mi := &file_hicon_sm_type_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DBConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBConfiguration) ProtoMessage() {}

func (x *DBConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_hicon_sm_type_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBConfiguration.ProtoReflect.Descriptor instead.
func (*DBConfiguration) Descriptor() ([]byte, []int) {
	return file_hicon_sm_type_proto_rawDescGZIP(), []int{2}
}

func (x *DBConfiguration) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DBConfiguration) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *DBConfiguration) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *DBConfiguration) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DBConfiguration) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *DBConfiguration) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *DBConfiguration) GetMaxCons() int32 {
	if x != nil {
		return x.MaxCons
	}
	return 0
}

func (x *DBConfiguration) GetTLS() *TLS {
	if x != nil {
		return x.TLS
	}
	return nil
}

// Column Configuration for Table
type ColumnConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Nullable      bool                   `protobuf:"varint,3,opt,name=nullable,proto3" json:"nullable,omitempty"`
	IsPrimaryKey  bool                   `protobuf:"varint,4,opt,name=is_primary_key,json=isPrimaryKey,proto3" json:"is_primary_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ColumnConfig) Reset() {
	*x = ColumnConfig{}
	mi := &file_hicon_sm_type_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ColumnConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnConfig) ProtoMessage() {}

func (x *ColumnConfig) ProtoReflect() protoreflect.Message {
	mi := &file_hicon_sm_type_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnConfig.ProtoReflect.Descriptor instead.
func (*ColumnConfig) Descriptor() ([]byte, []int) {
	return file_hicon_sm_type_proto_rawDescGZIP(), []int{3}
}

func (x *ColumnConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ColumnConfig) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ColumnConfig) GetNullable() bool {
	if x != nil {
		return x.Nullable
	}
	return false
}

func (x *ColumnConfig) GetIsPrimaryKey() bool {
	if x != nil {
		return x.IsPrimaryKey
	}
	return false
}

// Relation Column Configuration
type RelationColumnConfigs struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RefTable      string                 `protobuf:"bytes,2,opt,name=ref_table,json=refTable,proto3" json:"ref_table,omitempty"`
	Type          string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RelationColumnConfigs) Reset() {
	*x = RelationColumnConfigs{}
	mi := &file_hicon_sm_type_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RelationColumnConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationColumnConfigs) ProtoMessage() {}

func (x *RelationColumnConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_hicon_sm_type_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationColumnConfigs.ProtoReflect.Descriptor instead.
func (*RelationColumnConfigs) Descriptor() ([]byte, []int) {
	return file_hicon_sm_type_proto_rawDescGZIP(), []int{4}
}

func (x *RelationColumnConfigs) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RelationColumnConfigs) GetRefTable() string {
	if x != nil {
		return x.RefTable
	}
	return ""
}

func (x *RelationColumnConfigs) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// Table Configuration
type TableConfiguration struct {
	state                 protoimpl.MessageState   `protogen:"open.v1"`
	Name                  string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ColumnConfigs         []*ColumnConfig          `protobuf:"bytes,2,rep,name=column_configs,json=columnConfigs,proto3" json:"column_configs,omitempty"`
	RelationColumnConfigs []*RelationColumnConfigs `protobuf:"bytes,3,rep,name=relation_column_configs,json=relationColumnConfigs,proto3" json:"relation_column_configs,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *TableConfiguration) Reset() {
	*x = TableConfiguration{}
	mi := &file_hicon_sm_type_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TableConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableConfiguration) ProtoMessage() {}

func (x *TableConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_hicon_sm_type_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableConfiguration.ProtoReflect.Descriptor instead.
func (*TableConfiguration) Descriptor() ([]byte, []int) {
	return file_hicon_sm_type_proto_rawDescGZIP(), []int{5}
}

func (x *TableConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TableConfiguration) GetColumnConfigs() []*ColumnConfig {
	if x != nil {
		return x.ColumnConfigs
	}
	return nil
}

func (x *TableConfiguration) GetRelationColumnConfigs() []*RelationColumnConfigs {
	if x != nil {
		return x.RelationColumnConfigs
	}
	return nil
}

// Main Upsert Configuration
type UpsertConfiguration struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	DbConfiguration     *DBConfiguration       `protobuf:"bytes,1,opt,name=db_configuration,json=dbConfiguration,proto3" json:"db_configuration,omitempty"`
	TableConfigurations []*TableConfiguration  `protobuf:"bytes,2,rep,name=table_configurations,json=tableConfigurations,proto3" json:"table_configurations,omitempty"`
	Debug               bool                   `protobuf:"varint,3,opt,name=debug,proto3" json:"debug,omitempty"`
	DisableCache        bool                   `protobuf:"varint,4,opt,name=disable_cache,json=disableCache,proto3" json:"disable_cache,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *UpsertConfiguration) Reset() {
	*x = UpsertConfiguration{}
	mi := &file_hicon_sm_type_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertConfiguration) ProtoMessage() {}

func (x *UpsertConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_hicon_sm_type_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertConfiguration.ProtoReflect.Descriptor instead.
func (*UpsertConfiguration) Descriptor() ([]byte, []int) {
	return file_hicon_sm_type_proto_rawDescGZIP(), []int{6}
}

func (x *UpsertConfiguration) GetDbConfiguration() *DBConfiguration {
	if x != nil {
		return x.DbConfiguration
	}
	return nil
}

func (x *UpsertConfiguration) GetTableConfigurations() []*TableConfiguration {
	if x != nil {
		return x.TableConfigurations
	}
	return nil
}

func (x *UpsertConfiguration) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *UpsertConfiguration) GetDisableCache() bool {
	if x != nil {
		return x.DisableCache
	}
	return false
}

type FindByPrimaryKeys struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Table         string                 `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	DisableCache  bool                   `protobuf:"varint,2,opt,name=disable_cache,json=disableCache,proto3" json:"disable_cache,omitempty"`
	PrimaryKeys   map[string]*anypb.Any  `protobuf:"bytes,3,rep,name=primary_keys,json=primaryKeys,proto3" json:"primary_keys,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindByPrimaryKeys) Reset() {
	*x = FindByPrimaryKeys{}
	mi := &file_hicon_sm_type_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindByPrimaryKeys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByPrimaryKeys) ProtoMessage() {}

func (x *FindByPrimaryKeys) ProtoReflect() protoreflect.Message {
	mi := &file_hicon_sm_type_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByPrimaryKeys.ProtoReflect.Descriptor instead.
func (*FindByPrimaryKeys) Descriptor() ([]byte, []int) {
	return file_hicon_sm_type_proto_rawDescGZIP(), []int{7}
}

func (x *FindByPrimaryKeys) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *FindByPrimaryKeys) GetDisableCache() bool {
	if x != nil {
		return x.DisableCache
	}
	return false
}

func (x *FindByPrimaryKeys) GetPrimaryKeys() map[string]*anypb.Any {
	if x != nil {
		return x.PrimaryKeys
	}
	return nil
}

var File_hicon_sm_type_proto protoreflect.FileDescriptor

var file_hicon_sm_type_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2d, 0x73, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a,
	0x0c, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x68, 0x0a, 0x03, 0x54,
	0x4c, 0x53, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x70, 0x65, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x65, 0x72, 0x74, 0x50, 0x65, 0x6d, 0x12, 0x26, 0x0a,
	0x0f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x65, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x50, 0x65, 0x6d, 0x12, 0x1e, 0x0a, 0x0b, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x63, 0x61,
	0x5f, 0x70, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x6f, 0x74,
	0x43, 0x61, 0x50, 0x65, 0x6d, 0x22, 0xdf, 0x01, 0x0a, 0x0f, 0x44, 0x42, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78,
	0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78,
	0x43, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x03, 0x54, 0x4c, 0x53, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x54,
	0x4c, 0x53, 0x52, 0x03, 0x54, 0x4c, 0x53, 0x22, 0x78, 0x0a, 0x0c, 0x43, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x6e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x69,
	0x73, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65,
	0x79, 0x22, 0x5c, 0x0a, 0x15, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x65, 0x66, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x66, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0xc4, 0x01, 0x0a, 0x12, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x59, 0x0a, 0x17, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68,
	0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52,
	0x15, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0xeb, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46,
	0x0a, 0x10, 0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x68, 0x69, 0x63, 0x6f, 0x6e,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x42, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x64, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x14, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62,
	0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x22, 0xf7, 0x01, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x68, 0x69,
	0x63, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x73, 0x1a, 0x54, 0x0a, 0x10, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2f,
	0x5a, 0x0d, 0x2e, 0x2f, 0x73, 0x71, 0x6c, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0xaa,
	0x02, 0x1d, 0x4d, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_hicon_sm_type_proto_rawDescOnce sync.Once
	file_hicon_sm_type_proto_rawDescData []byte
)

func file_hicon_sm_type_proto_rawDescGZIP() []byte {
	file_hicon_sm_type_proto_rawDescOnce.Do(func() {
		file_hicon_sm_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_hicon_sm_type_proto_rawDesc), len(file_hicon_sm_type_proto_rawDesc)))
	})
	return file_hicon_sm_type_proto_rawDescData
}

var file_hicon_sm_type_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_hicon_sm_type_proto_goTypes = []any{
	(*BaseResponse)(nil),          // 0: hicon.type.BaseResponse
	(*TLS)(nil),                   // 1: hicon.type.TLS
	(*DBConfiguration)(nil),       // 2: hicon.type.DBConfiguration
	(*ColumnConfig)(nil),          // 3: hicon.type.ColumnConfig
	(*RelationColumnConfigs)(nil), // 4: hicon.type.RelationColumnConfigs
	(*TableConfiguration)(nil),    // 5: hicon.type.TableConfiguration
	(*UpsertConfiguration)(nil),   // 6: hicon.type.UpsertConfiguration
	(*FindByPrimaryKeys)(nil),     // 7: hicon.type.FindByPrimaryKeys
	nil,                           // 8: hicon.type.FindByPrimaryKeys.PrimaryKeysEntry
	(*anypb.Any)(nil),             // 9: google.protobuf.Any
}
var file_hicon_sm_type_proto_depIdxs = []int32{
	9, // 0: hicon.type.BaseResponse.data:type_name -> google.protobuf.Any
	1, // 1: hicon.type.DBConfiguration.TLS:type_name -> hicon.type.TLS
	3, // 2: hicon.type.TableConfiguration.column_configs:type_name -> hicon.type.ColumnConfig
	4, // 3: hicon.type.TableConfiguration.relation_column_configs:type_name -> hicon.type.RelationColumnConfigs
	2, // 4: hicon.type.UpsertConfiguration.db_configuration:type_name -> hicon.type.DBConfiguration
	5, // 5: hicon.type.UpsertConfiguration.table_configurations:type_name -> hicon.type.TableConfiguration
	8, // 6: hicon.type.FindByPrimaryKeys.primary_keys:type_name -> hicon.type.FindByPrimaryKeys.PrimaryKeysEntry
	9, // 7: hicon.type.FindByPrimaryKeys.PrimaryKeysEntry.value:type_name -> google.protobuf.Any
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_hicon_sm_type_proto_init() }
func file_hicon_sm_type_proto_init() {
	if File_hicon_sm_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_hicon_sm_type_proto_rawDesc), len(file_hicon_sm_type_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hicon_sm_type_proto_goTypes,
		DependencyIndexes: file_hicon_sm_type_proto_depIdxs,
		MessageInfos:      file_hicon_sm_type_proto_msgTypes,
	}.Build()
	File_hicon_sm_type_proto = out.File
	file_hicon_sm_type_proto_goTypes = nil
	file_hicon_sm_type_proto_depIdxs = nil
}
