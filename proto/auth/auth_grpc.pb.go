// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: auth/auth.proto

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
	AuthService_Login_FullMethodName            = "/auth.AuthService/Login"
	AuthService_ChangePassword_FullMethodName   = "/auth.AuthService/ChangePassword"
	AuthService_CreateEmployee_FullMethodName   = "/auth.AuthService/CreateEmployee"
	AuthService_CreateDepartment_FullMethodName = "/auth.AuthService/CreateDepartment"
	AuthService_ListEmployee_FullMethodName     = "/auth.AuthService/ListEmployee"
	AuthService_ListDepartment_FullMethodName   = "/auth.AuthService/ListDepartment"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*gcommon.EmptyResponse, error)
	CreateEmployee(ctx context.Context, in *CreateEmployeeRequest, opts ...grpc.CallOption) (*gcommon.EmptyResponse, error)
	CreateDepartment(ctx context.Context, in *CreateDepartmentRequest, opts ...grpc.CallOption) (*gcommon.EmptyResponse, error)
	ListEmployee(ctx context.Context, in *ListEmployeeRequest, opts ...grpc.CallOption) (*ListEmployeeResponse, error)
	ListDepartment(ctx context.Context, in *ListDepartmentRequest, opts ...grpc.CallOption) (*ListDepartmentResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AuthService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*gcommon.EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(gcommon.EmptyResponse)
	err := c.cc.Invoke(ctx, AuthService_ChangePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CreateEmployee(ctx context.Context, in *CreateEmployeeRequest, opts ...grpc.CallOption) (*gcommon.EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(gcommon.EmptyResponse)
	err := c.cc.Invoke(ctx, AuthService_CreateEmployee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CreateDepartment(ctx context.Context, in *CreateDepartmentRequest, opts ...grpc.CallOption) (*gcommon.EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(gcommon.EmptyResponse)
	err := c.cc.Invoke(ctx, AuthService_CreateDepartment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ListEmployee(ctx context.Context, in *ListEmployeeRequest, opts ...grpc.CallOption) (*ListEmployeeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListEmployeeResponse)
	err := c.cc.Invoke(ctx, AuthService_ListEmployee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ListDepartment(ctx context.Context, in *ListDepartmentRequest, opts ...grpc.CallOption) (*ListDepartmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDepartmentResponse)
	err := c.cc.Invoke(ctx, AuthService_ListDepartment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*gcommon.EmptyResponse, error)
	CreateEmployee(context.Context, *CreateEmployeeRequest) (*gcommon.EmptyResponse, error)
	CreateDepartment(context.Context, *CreateDepartmentRequest) (*gcommon.EmptyResponse, error)
	ListEmployee(context.Context, *ListEmployeeRequest) (*ListEmployeeResponse, error)
	ListDepartment(context.Context, *ListDepartmentRequest) (*ListDepartmentResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*gcommon.EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedAuthServiceServer) CreateEmployee(context.Context, *CreateEmployeeRequest) (*gcommon.EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmployee not implemented")
}
func (UnimplementedAuthServiceServer) CreateDepartment(context.Context, *CreateDepartmentRequest) (*gcommon.EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDepartment not implemented")
}
func (UnimplementedAuthServiceServer) ListEmployee(context.Context, *ListEmployeeRequest) (*ListEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEmployee not implemented")
}
func (UnimplementedAuthServiceServer) ListDepartment(context.Context, *ListDepartmentRequest) (*ListDepartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDepartment not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CreateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CreateEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateEmployee(ctx, req.(*CreateEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CreateDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CreateDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateDepartment(ctx, req.(*CreateDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ListEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ListEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ListEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ListEmployee(ctx, req.(*ListEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ListDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ListDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ListDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ListDepartment(ctx, req.(*ListDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _AuthService_ChangePassword_Handler,
		},
		{
			MethodName: "CreateEmployee",
			Handler:    _AuthService_CreateEmployee_Handler,
		},
		{
			MethodName: "CreateDepartment",
			Handler:    _AuthService_CreateDepartment_Handler,
		},
		{
			MethodName: "ListEmployee",
			Handler:    _AuthService_ListEmployee_Handler,
		},
		{
			MethodName: "ListDepartment",
			Handler:    _AuthService_ListDepartment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}
