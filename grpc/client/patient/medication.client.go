package patientgrpcclient

import (
	"context"
	"gateway/proto/gcommon"
	"gateway/proto/patient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type MedicationGrpcClient struct {
	timeout time.Duration
	client  patient.MedicationServiceClient
}

func NewMedicationGrpcClient(props GrpcProps) *MedicationGrpcClient {
	return &MedicationGrpcClient{
		timeout: time.Duration(props.Config.GetInt("GRPC_TIMEOUT", 10)) * time.Second,
		client:  patient.NewMedicationServiceClient(props.Connection),
	}
}

func (c *MedicationGrpcClient) ListMedication(req *patient.ListMedicationRequest, md metadata.MD, opts ...grpc.CallOption) (*patient.ListMedicationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.client.ListMedication(ctx, req, opts...)
}

func (c *MedicationGrpcClient) UpsertMedication(req *patient.UpsertMedicationRequest, md metadata.MD, opts ...grpc.CallOption) (*gcommon.EmptyResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.client.UpsertMedication(ctx, req, opts...)
}
