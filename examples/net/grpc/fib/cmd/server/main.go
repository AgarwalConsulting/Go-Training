package main

import (
	"net"
	"os"

	pb "algogrit.com/fib-grpc/api"
	"algogrit.com/fib-grpc/fibonacci/service"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func withServerUnaryInterceptor() grpc.ServerOption {
	// basic := auth.NewBasicAuth()
	// return grpc.UnaryInterceptor(logUnaryInterceptor)
	return grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
		logUnaryInterceptor,
		// grpc_auth.UnaryServerInterceptor(basic.Interceptor),
	))
}

var grpcPort, diagnosticsPort string

func init() {
	var ok bool

	grpcPort, ok = os.LookupEnv("GRPC_PORT")
	if !ok {
		grpcPort = "50001"
	}

	diagnosticsPort, ok = os.LookupEnv("DIAGNOSTICS_PORT")
	if !ok {
		diagnosticsPort = "8081"
	}
}

func main() {
	lis, err := net.Listen("tcp", ":"+grpcPort)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	uInterceptor := withServerUnaryInterceptor()
	sInterceptor := grpc.StreamInterceptor(logStreamInterceptor)

	s := grpc.NewServer(uInterceptor, sInterceptor)
	pb.RegisterFibonacciServer(s, service.NewFibonacciServer())

	log.Infof("Starting fib server on port %s...\n", grpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
