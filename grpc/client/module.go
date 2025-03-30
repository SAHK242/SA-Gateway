package grpcclient

import (
	auth "gateway/grpc/client/auth"
	"go.uber.org/fx"
)

var Module = fx.Module("grpc_client",
	auth.Module,
)
