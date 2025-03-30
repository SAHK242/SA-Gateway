package grpc

import (
	"go.uber.org/fx"

	grpcclient "gateway/grpc/client"
	grpcconn "gateway/grpc/conn"
	grpchelper "gateway/grpc/helper"
)

var Module = fx.Module("grpc",
	grpcconn.Module,
	grpcclient.Module,
	grpchelper.Module,
)
