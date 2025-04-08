package authgrpcclient

import (
	"context"
	"gateway/proto/auth"
	"gateway/proto/gcommon"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type EmployeeGrpcClient struct {
	timeout time.Duration
	client  auth.EmployeeServiceClient
}

func NewEmployeeGrpcClient(props GrpcProps) *EmployeeGrpcClient {
	return &EmployeeGrpcClient{
		timeout: time.Duration(props.Config.GetInt("GRPC_TIMEOUT", 10)) * time.Second,
		client:  auth.NewEmployeeServiceClient(props.Connection),
	}
}

func (c *EmployeeGrpcClient) CreateEmployee(req *auth.CreateEmployeeRequest, md metadata.MD, opts ...grpc.CallOption) (*gcommon.EmptyResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.client.CreateEmployee(ctx, req, opts...)
}

func (c *EmployeeGrpcClient) ListEmployee(req *auth.ListEmployeeRequest, md metadata.MD, opts ...grpc.CallOption) (*auth.ListEmployeeResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.client.ListEmployee(ctx, req, opts...)
}
