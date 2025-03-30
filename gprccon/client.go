package grpcconn

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewConnection(addr string, transport credentials.TransportCredentials, opts ...Option) *grpc.ClientConn {
	o := eval(opts...)

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(transport),
		grpc.WithTransportCredentials(transport),
		grpc.WithMaxHeaderListSize(o.maxHeaderSize),
		grpc.WithChainUnaryInterceptor(o.interceptors...),
	)

	if err != nil {
		panic(err)
	}

	return conn
}
