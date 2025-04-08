package authgrpcclient

import (
	"context"
	"gateway/proto/auth"
	"gateway/proto/gcommon"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type DepartmentGrpcClient struct {
	timeout time.Duration
	client  auth.DepartmentServiceClient
}

func NewDepartmentGrpcClient(props GrpcProps) *DepartmentGrpcClient {
	return &DepartmentGrpcClient{
		timeout: time.Duration(props.Config.GetInt("GRPC_TIMEOUT", 10)) * time.Second,
		client:  auth.NewDepartmentServiceClient(props.Connection),
	}
}

func (c *DepartmentGrpcClient) CreateDepartment(req *auth.CreateDepartmentRequest, md metadata.MD, opts ...grpc.CallOption) (*gcommon.EmptyResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.client.CreateDepartment(ctx, req, opts...)
}

func (c *DepartmentGrpcClient) ListDepartment(req *auth.ListDepartmentRequest, md metadata.MD, opts ...grpc.CallOption) (*auth.ListDepartmentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.client.ListDepartment(ctx, req, opts...)
}
