package patientgrpcclient

import (
	"gateway/config"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type GrpcProps struct {
	fx.In
	Config     config.Config
	Connection *grpc.ClientConn `name:"patientConn"`
}

var Module = fx.Provide(
	NewPatientGrpcClient,
	NewMedicationGrpcClient,
)
