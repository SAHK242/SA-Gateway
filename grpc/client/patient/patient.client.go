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

func (c *PatientGrpcClient) UpsertMedicalRecord(req *patient.UpsertMedicalRecordRequest, md metadata.MD, opts ...grpc.CallOption) (*gcommon.EmptyResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.client.UpsertMedicalRecord(ctx, req, opts...)
}

func (c *PatientGrpcClient) UpsertMedicalTreatment(req *patient.UpsertMedicalTreatmentRequest, md metadata.MD, opts ...grpc.CallOption) (*gcommon.EmptyResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.client.UpsertMedicalTreatment(ctx, req, opts...)
}

func (c *PatientGrpcClient) UpsertMedicalSurgery(req *patient.UpsertMedicalSurgeryRequest, md metadata.MD, opts ...grpc.CallOption) (*gcommon.EmptyResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.client.UpsertMedicalSurgery(ctx, req, opts...)
}

func (c *PatientGrpcClient) UpsertMedicalPrescription(req *patient.UpsertMedicalPrescriptionRequest, md metadata.MD, opts ...grpc.CallOption) (*gcommon.EmptyResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.client.UpsertMedicalPrescription(ctx, req, opts...)
}

func (c *PatientGrpcClient) GetMedicalHistory(req *patient.GetMedicalHistoryRequest, md metadata.MD, opts ...grpc.CallOption) (*patient.GetMedicalHistoryResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.client.GetMedicalHistory(ctx, req, opts...)
}

func (c *PatientGrpcClient) GetMedicalHistoryDetail(req *gcommon.IdRequest, md metadata.MD, opts ...grpc.CallOption) (*patient.GetMedicalHistoryDetailResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, md)

	return c.client.GetMedicalHistoryDetail(ctx, req, opts...)
}
