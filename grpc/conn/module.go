package grpcconn

import (
	"context"
	"gateway/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	conn "gateway/gprccon"
)

var transport = insecure.NewCredentials()

var Module = fx.Provide(
	fx.Annotate(
		NewAuthGrpcConn,
		fx.ResultTags(`name:"authConn"`),
	),
	fx.Annotate(
		NewPatientGrpcConn,
		fx.ResultTags(`name:"patientConn"`),
	),
)

func newGrpcConn(cfg config.Config, logger *zap.SugaredLogger, serverName string) *grpc.ClientConn {
	server := cfg.MustGetString(serverName)

	connection := conn.NewConnection(server, transport,
		conn.WithMaxHeaderSize(cfg.MustGetInt32("GRPC_MAX_HEADER_SIZE")),
		conn.WithInterceptors(
			CustomClientLoggingInterceptor(logger.Desugar()),
		),
	)

	return connection
}

func NewAuthGrpcConn(cfg config.Config, logger *zap.SugaredLogger) *grpc.ClientConn {
	return newGrpcConn(cfg, logger, "AUTH_GRPC_SERVER")
}

func NewPatientGrpcConn(cfg config.Config, logger *zap.SugaredLogger) *grpc.ClientConn {
	return newGrpcConn(cfg, logger, "PATIENT_GRPC_SERVER")
}

func CustomClientLoggingInterceptor(logger *zap.Logger) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		logger.Info("gRPC request", zap.String("method", method))

		// Invoke the actual RPC call
		err := invoker(ctx, method, req, reply, cc, opts...)

		if err != nil {
			logger.Error("gRPC call failed", zap.String("method", method), zap.Error(err))
		} else {
			logger.Info("gRPC response received", zap.String("method", method))
		}

		return err
	}
}
