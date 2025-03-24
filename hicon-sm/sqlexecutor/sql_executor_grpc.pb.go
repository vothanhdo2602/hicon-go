// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: hicon-sm/sql_executor.proto

package sqlexecutor

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SQLExecutor_UpsertConfiguration_FullMethodName = "/SQLExecutor/UpsertConfiguration"
	SQLExecutor_FindByPK_FullMethodName            = "/SQLExecutor/FindByPK"
	SQLExecutor_FindOne_FullMethodName             = "/SQLExecutor/FindOne"
	SQLExecutor_FindAll_FullMethodName             = "/SQLExecutor/FindAll"
	SQLExecutor_Exec_FullMethodName                = "/SQLExecutor/Exec"
	SQLExecutor_BulkInsert_FullMethodName          = "/SQLExecutor/BulkInsert"
	SQLExecutor_UpdateByPK_FullMethodName          = "/SQLExecutor/UpdateByPK"
	SQLExecutor_UpdateAll_FullMethodName           = "/SQLExecutor/UpdateAll"
	SQLExecutor_BulkUpdateByPK_FullMethodName      = "/SQLExecutor/BulkUpdateByPK"
	SQLExecutor_DeleteByPK_FullMethodName          = "/SQLExecutor/DeleteByPK"
	SQLExecutor_BulkWriteWithTx_FullMethodName     = "/SQLExecutor/BulkWriteWithTx"
)

// SQLExecutorClient is the client API for SQLExecutor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SQLExecutorClient interface {
	UpsertConfiguration(ctx context.Context, in *UpsertConfiguration, opts ...grpc.CallOption) (*BaseResponse, error)
	FindByPK(ctx context.Context, in *FindByPK, opts ...grpc.CallOption) (*BaseResponse, error)
	FindOne(ctx context.Context, in *FindOne, opts ...grpc.CallOption) (*BaseResponse, error)
	FindAll(ctx context.Context, in *FindAll, opts ...grpc.CallOption) (*BaseResponse, error)
	Exec(ctx context.Context, in *Exec, opts ...grpc.CallOption) (*BaseResponse, error)
	BulkInsert(ctx context.Context, in *BulkInsert, opts ...grpc.CallOption) (*BaseResponse, error)
	UpdateByPK(ctx context.Context, in *UpdateByPK, opts ...grpc.CallOption) (*BaseResponse, error)
	UpdateAll(ctx context.Context, in *UpdateAll, opts ...grpc.CallOption) (*BaseResponse, error)
	BulkUpdateByPK(ctx context.Context, in *BulkUpdateByPK, opts ...grpc.CallOption) (*BaseResponse, error)
	DeleteByPK(ctx context.Context, in *DeleteByPK, opts ...grpc.CallOption) (*BaseResponse, error)
	BulkWriteWithTx(ctx context.Context, in *BulkWriteWithTx, opts ...grpc.CallOption) (*BaseResponse, error)
}

type sQLExecutorClient struct {
	cc grpc.ClientConnInterface
}

func NewSQLExecutorClient(cc grpc.ClientConnInterface) SQLExecutorClient {
	return &sQLExecutorClient{cc}
}

func (c *sQLExecutorClient) UpsertConfiguration(ctx context.Context, in *UpsertConfiguration, opts ...grpc.CallOption) (*BaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, SQLExecutor_UpsertConfiguration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sQLExecutorClient) FindByPK(ctx context.Context, in *FindByPK, opts ...grpc.CallOption) (*BaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, SQLExecutor_FindByPK_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sQLExecutorClient) FindOne(ctx context.Context, in *FindOne, opts ...grpc.CallOption) (*BaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, SQLExecutor_FindOne_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sQLExecutorClient) FindAll(ctx context.Context, in *FindAll, opts ...grpc.CallOption) (*BaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, SQLExecutor_FindAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sQLExecutorClient) Exec(ctx context.Context, in *Exec, opts ...grpc.CallOption) (*BaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, SQLExecutor_Exec_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sQLExecutorClient) BulkInsert(ctx context.Context, in *BulkInsert, opts ...grpc.CallOption) (*BaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, SQLExecutor_BulkInsert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sQLExecutorClient) UpdateByPK(ctx context.Context, in *UpdateByPK, opts ...grpc.CallOption) (*BaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, SQLExecutor_UpdateByPK_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sQLExecutorClient) UpdateAll(ctx context.Context, in *UpdateAll, opts ...grpc.CallOption) (*BaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, SQLExecutor_UpdateAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sQLExecutorClient) BulkUpdateByPK(ctx context.Context, in *BulkUpdateByPK, opts ...grpc.CallOption) (*BaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, SQLExecutor_BulkUpdateByPK_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sQLExecutorClient) DeleteByPK(ctx context.Context, in *DeleteByPK, opts ...grpc.CallOption) (*BaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, SQLExecutor_DeleteByPK_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sQLExecutorClient) BulkWriteWithTx(ctx context.Context, in *BulkWriteWithTx, opts ...grpc.CallOption) (*BaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, SQLExecutor_BulkWriteWithTx_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SQLExecutorServer is the server API for SQLExecutor service.
// All implementations must embed UnimplementedSQLExecutorServer
// for forward compatibility.
type SQLExecutorServer interface {
	UpsertConfiguration(context.Context, *UpsertConfiguration) (*BaseResponse, error)
	FindByPK(context.Context, *FindByPK) (*BaseResponse, error)
	FindOne(context.Context, *FindOne) (*BaseResponse, error)
	FindAll(context.Context, *FindAll) (*BaseResponse, error)
	Exec(context.Context, *Exec) (*BaseResponse, error)
	BulkInsert(context.Context, *BulkInsert) (*BaseResponse, error)
	UpdateByPK(context.Context, *UpdateByPK) (*BaseResponse, error)
	UpdateAll(context.Context, *UpdateAll) (*BaseResponse, error)
	BulkUpdateByPK(context.Context, *BulkUpdateByPK) (*BaseResponse, error)
	DeleteByPK(context.Context, *DeleteByPK) (*BaseResponse, error)
	BulkWriteWithTx(context.Context, *BulkWriteWithTx) (*BaseResponse, error)
	mustEmbedUnimplementedSQLExecutorServer()
}

// UnimplementedSQLExecutorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSQLExecutorServer struct{}

func (UnimplementedSQLExecutorServer) UpsertConfiguration(context.Context, *UpsertConfiguration) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertConfiguration not implemented")
}
func (UnimplementedSQLExecutorServer) FindByPK(context.Context, *FindByPK) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByPK not implemented")
}
func (UnimplementedSQLExecutorServer) FindOne(context.Context, *FindOne) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (UnimplementedSQLExecutorServer) FindAll(context.Context, *FindAll) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedSQLExecutorServer) Exec(context.Context, *Exec) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}
func (UnimplementedSQLExecutorServer) BulkInsert(context.Context, *BulkInsert) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkInsert not implemented")
}
func (UnimplementedSQLExecutorServer) UpdateByPK(context.Context, *UpdateByPK) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateByPK not implemented")
}
func (UnimplementedSQLExecutorServer) UpdateAll(context.Context, *UpdateAll) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAll not implemented")
}
func (UnimplementedSQLExecutorServer) BulkUpdateByPK(context.Context, *BulkUpdateByPK) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkUpdateByPK not implemented")
}
func (UnimplementedSQLExecutorServer) DeleteByPK(context.Context, *DeleteByPK) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByPK not implemented")
}
func (UnimplementedSQLExecutorServer) BulkWriteWithTx(context.Context, *BulkWriteWithTx) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkWriteWithTx not implemented")
}
func (UnimplementedSQLExecutorServer) mustEmbedUnimplementedSQLExecutorServer() {}
func (UnimplementedSQLExecutorServer) testEmbeddedByValue()                     {}

// UnsafeSQLExecutorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SQLExecutorServer will
// result in compilation errors.
type UnsafeSQLExecutorServer interface {
	mustEmbedUnimplementedSQLExecutorServer()
}

func RegisterSQLExecutorServer(s grpc.ServiceRegistrar, srv SQLExecutorServer) {
	// If the following call pancis, it indicates UnimplementedSQLExecutorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SQLExecutor_ServiceDesc, srv)
}

func _SQLExecutor_UpsertConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertConfiguration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLExecutorServer).UpsertConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SQLExecutor_UpsertConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLExecutorServer).UpsertConfiguration(ctx, req.(*UpsertConfiguration))
	}
	return interceptor(ctx, in, info, handler)
}

func _SQLExecutor_FindByPK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByPK)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLExecutorServer).FindByPK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SQLExecutor_FindByPK_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLExecutorServer).FindByPK(ctx, req.(*FindByPK))
	}
	return interceptor(ctx, in, info, handler)
}

func _SQLExecutor_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOne)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLExecutorServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SQLExecutor_FindOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLExecutorServer).FindOne(ctx, req.(*FindOne))
	}
	return interceptor(ctx, in, info, handler)
}

func _SQLExecutor_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAll)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLExecutorServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SQLExecutor_FindAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLExecutorServer).FindAll(ctx, req.(*FindAll))
	}
	return interceptor(ctx, in, info, handler)
}

func _SQLExecutor_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Exec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLExecutorServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SQLExecutor_Exec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLExecutorServer).Exec(ctx, req.(*Exec))
	}
	return interceptor(ctx, in, info, handler)
}

func _SQLExecutor_BulkInsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkInsert)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLExecutorServer).BulkInsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SQLExecutor_BulkInsert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLExecutorServer).BulkInsert(ctx, req.(*BulkInsert))
	}
	return interceptor(ctx, in, info, handler)
}

func _SQLExecutor_UpdateByPK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateByPK)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLExecutorServer).UpdateByPK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SQLExecutor_UpdateByPK_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLExecutorServer).UpdateByPK(ctx, req.(*UpdateByPK))
	}
	return interceptor(ctx, in, info, handler)
}

func _SQLExecutor_UpdateAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAll)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLExecutorServer).UpdateAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SQLExecutor_UpdateAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLExecutorServer).UpdateAll(ctx, req.(*UpdateAll))
	}
	return interceptor(ctx, in, info, handler)
}

func _SQLExecutor_BulkUpdateByPK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkUpdateByPK)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLExecutorServer).BulkUpdateByPK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SQLExecutor_BulkUpdateByPK_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLExecutorServer).BulkUpdateByPK(ctx, req.(*BulkUpdateByPK))
	}
	return interceptor(ctx, in, info, handler)
}

func _SQLExecutor_DeleteByPK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteByPK)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLExecutorServer).DeleteByPK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SQLExecutor_DeleteByPK_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLExecutorServer).DeleteByPK(ctx, req.(*DeleteByPK))
	}
	return interceptor(ctx, in, info, handler)
}

func _SQLExecutor_BulkWriteWithTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkWriteWithTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLExecutorServer).BulkWriteWithTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SQLExecutor_BulkWriteWithTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLExecutorServer).BulkWriteWithTx(ctx, req.(*BulkWriteWithTx))
	}
	return interceptor(ctx, in, info, handler)
}

// SQLExecutor_ServiceDesc is the grpc.ServiceDesc for SQLExecutor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SQLExecutor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SQLExecutor",
	HandlerType: (*SQLExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertConfiguration",
			Handler:    _SQLExecutor_UpsertConfiguration_Handler,
		},
		{
			MethodName: "FindByPK",
			Handler:    _SQLExecutor_FindByPK_Handler,
		},
		{
			MethodName: "FindOne",
			Handler:    _SQLExecutor_FindOne_Handler,
		},
		{
			MethodName: "FindAll",
			Handler:    _SQLExecutor_FindAll_Handler,
		},
		{
			MethodName: "Exec",
			Handler:    _SQLExecutor_Exec_Handler,
		},
		{
			MethodName: "BulkInsert",
			Handler:    _SQLExecutor_BulkInsert_Handler,
		},
		{
			MethodName: "UpdateByPK",
			Handler:    _SQLExecutor_UpdateByPK_Handler,
		},
		{
			MethodName: "UpdateAll",
			Handler:    _SQLExecutor_UpdateAll_Handler,
		},
		{
			MethodName: "BulkUpdateByPK",
			Handler:    _SQLExecutor_BulkUpdateByPK_Handler,
		},
		{
			MethodName: "DeleteByPK",
			Handler:    _SQLExecutor_DeleteByPK_Handler,
		},
		{
			MethodName: "BulkWriteWithTx",
			Handler:    _SQLExecutor_BulkWriteWithTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hicon-sm/sql_executor.proto",
}
