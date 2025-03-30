package grpcconn

import (
	"math"

	"google.golang.org/grpc"
)

const defaultMaxHeaderSize = uint32(8192)

type (
	options struct {
		interceptors  []grpc.UnaryClientInterceptor
		maxHeaderSize uint32
	}

	Option func(*options)
)

func WithInterceptors(interceptors ...grpc.UnaryClientInterceptor) Option {
	return func(o *options) {
		o.interceptors = interceptors
	}
}

func WithMaxHeaderSize(maxHeaderSize int32) Option {
	return func(o *options) {
		o.maxHeaderSize = uint32(math.Max(float64(defaultMaxHeaderSize), float64(maxHeaderSize)))
	}
}

func eval(opts ...Option) *options {
	o := &options{
		maxHeaderSize: defaultMaxHeaderSize,
		interceptors:  []grpc.UnaryClientInterceptor{},
	}

	for _, opt := range opts {
		opt(o)
	}

	return o
}
