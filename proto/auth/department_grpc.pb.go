// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: auth/department.proto

package auth

import (
	context "context"
	gcommon "gateway/proto/gcommon"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	DepartmentService_CreateDepartment_FullMethodName = "/auth.DepartmentService/CreateDepartment"
	DepartmentService_ListDepartment_FullMethodName   = "/auth.DepartmentService/ListDepartment"
)

// DepartmentServiceClient is the client API for DepartmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DepartmentServiceClient interface {
	CreateDepartment(ctx context.Context, in *CreateDepartmentRequest, opts ...grpc.CallOption) (*gcommon.EmptyResponse, error)
	ListDepartment(ctx context.Context, in *ListDepartmentRequest, opts ...grpc.CallOption) (*ListDepartmentResponse, error)
}

type departmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDepartmentServiceClient(cc grpc.ClientConnInterface) DepartmentServiceClient {
	return &departmentServiceClient{cc}
}

func (c *departmentServiceClient) CreateDepartment(ctx context.Context, in *CreateDepartmentRequest, opts ...grpc.CallOption) (*gcommon.EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(gcommon.EmptyResponse)
	err := c.cc.Invoke(ctx, DepartmentService_CreateDepartment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) ListDepartment(ctx context.Context, in *ListDepartmentRequest, opts ...grpc.CallOption) (*ListDepartmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDepartmentResponse)
	err := c.cc.Invoke(ctx, DepartmentService_ListDepartment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DepartmentServiceServer is the server API for DepartmentService service.
// All implementations must embed UnimplementedDepartmentServiceServer
// for forward compatibility
type DepartmentServiceServer interface {
	CreateDepartment(context.Context, *CreateDepartmentRequest) (*gcommon.EmptyResponse, error)
	ListDepartment(context.Context, *ListDepartmentRequest) (*ListDepartmentResponse, error)
	mustEmbedUnimplementedDepartmentServiceServer()
}

// UnimplementedDepartmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDepartmentServiceServer struct {
}

func (UnimplementedDepartmentServiceServer) CreateDepartment(context.Context, *CreateDepartmentRequest) (*gcommon.EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDepartment not implemented")
}
func (UnimplementedDepartmentServiceServer) ListDepartment(context.Context, *ListDepartmentRequest) (*ListDepartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDepartment not implemented")
}
func (UnimplementedDepartmentServiceServer) mustEmbedUnimplementedDepartmentServiceServer() {}

// UnsafeDepartmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DepartmentServiceServer will
// result in compilation errors.
type UnsafeDepartmentServiceServer interface {
	mustEmbedUnimplementedDepartmentServiceServer()
}

func RegisterDepartmentServiceServer(s grpc.ServiceRegistrar, srv DepartmentServiceServer) {
	s.RegisterService(&DepartmentService_ServiceDesc, srv)
}

func _DepartmentService_CreateDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).CreateDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentService_CreateDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).CreateDepartment(ctx, req.(*CreateDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_ListDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).ListDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentService_ListDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).ListDepartment(ctx, req.(*ListDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DepartmentService_ServiceDesc is the grpc.ServiceDesc for DepartmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DepartmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.DepartmentService",
	HandlerType: (*DepartmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDepartment",
			Handler:    _DepartmentService_CreateDepartment_Handler,
		},
		{
			MethodName: "ListDepartment",
			Handler:    _DepartmentService_ListDepartment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/department.proto",
}
