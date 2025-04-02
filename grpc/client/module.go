package grpcclient

import (
	auth "gateway/grpc/client/auth"
	patient "gateway/grpc/client/patient"
	"go.uber.org/fx"
)

var Module = fx.Module("grpc_client",
	auth.Module,
	patient.Module,
)
