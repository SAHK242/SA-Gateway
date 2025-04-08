package authgrpcclient

import (
	"context"
	"gateway/proto/auth"
	"gateway/proto/gcommon"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type AuthGrpcClient struct {
	timeout time.Duration
	client  auth.AuthServiceClient
}

func NewAuthGrpcClient(props GrpcProps) *AuthGrpcClient {
	return &AuthGrpcClient{
		timeout: time.Duration(props.Config.GetInt("GRPC_TIMEOUT", 10)) * time.Second,
		client:  auth.NewAuthServiceClient(props.Connection),
	}
}

func (c *AuthGrpcClient) Login(req *auth.LoginRequest, md metadata.MD, opts ...grpc.CallOption) (*auth.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.client.Login(ctx, req, opts...)
}

func (c *AuthGrpcClient) ChangePassword(req *auth.ChangePasswordRequest, md metadata.MD, opts ...grpc.CallOption) (*gcommon.EmptyResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.client.ChangePassword(ctx, req, opts...)
}
