package main

import (
	"algogrit.com/fib-grpc/pkg/auth"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

func withServerUnaryInterceptor(enableAuth bool) grpc.ServerOption {
	uInterceptorChain := grpcMiddleware.ChainUnaryServer(
		logUnaryInterceptor,
		grpc_prometheus.UnaryServerInterceptor,
	)

	if enableAuth {
		basic := auth.NewBasicAuth()
		uInterceptorChain = grpcMiddleware.ChainUnaryServer(
			uInterceptorChain,
			grpc_auth.UnaryServerInterceptor(basic.Interceptor),
		)
	}

	return grpc.UnaryInterceptor(uInterceptorChain)
}

func withServerStreamInterceptor(enableAuth bool) grpc.ServerOption {
	sInterceptorChain := grpcMiddleware.ChainStreamServer(
		logStreamInterceptor,
		grpc_prometheus.StreamServerInterceptor,
	)

	if enableAuth {
		basic := auth.NewBasicAuth()
		sInterceptorChain = grpcMiddleware.ChainStreamServer(
			sInterceptorChain,
			grpc_auth.StreamServerInterceptor(basic.Interceptor),
		)
	}

	return grpc.StreamInterceptor(sInterceptorChain)
}
