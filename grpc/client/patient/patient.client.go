package patientgrpcclient

import (
	"context"
	"gateway/proto/gcommon"
	"gateway/proto/patient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type PatientGrpcClient struct {
	timeout time.Duration
	client  patient.PatientServiceClient
}

func NewPatientGrpcClient(props GrpcProps) *PatientGrpcClient {
	return &PatientGrpcClient{
		timeout: time.Duration(props.Config.GetInt("GRPC_TIMEOUT", 10)) * time.Second,
		client:  patient.NewPatientServiceClient(props.Connection),
	}
}

func (c *PatientGrpcClient) UpsertPatient(req *patient.UpsertPatientRequest, md metadata.MD, opts ...grpc.CallOption) (*gcommon.EmptyResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.client.UpsertPatient(ctx, req, opts...)
}

func (c *PatientGrpcClient) GetPatient(req *gcommon.IdRequest, md metadata.MD, opts ...grpc.CallOption) (*patient.GetPatientResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.client.GetPatient(ctx, req, opts...)
}

func (c *PatientGrpcClient) ListPatient(req *patient.ListPatientRequest, md metadata.MD, opts ...grpc.CallOption) (*patient.ListPatientResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.client.ListPatient(ctx, req, opts...)
}
