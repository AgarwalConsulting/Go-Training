package main

import (
	"net"
	"net/http"
	"os"

	pb "algogrit.com/fib-grpc/api"
	"algogrit.com/fib-grpc/fibonacci/service"
	"algogrit.com/fib-grpc/pkg/auth"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func withServerStreamInterceptor() grpc.ServerOption {
	basic := auth.NewBasicAuth()

	return grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
		logStreamInterceptor,
		grpc_auth.StreamServerInterceptor(basic.Interceptor),
	))
}

func withServerUnaryInterceptor() grpc.ServerOption {
	basic := auth.NewBasicAuth()

	return grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
		logUnaryInterceptor,
		grpc_auth.UnaryServerInterceptor(basic.Interceptor),
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

	// With TLS
	// creds, _ := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
	// s := grpc.NewServer(grpc.Creds(creds))

	// With logging
	// uInterceptor := grpc.UnaryInterceptor(logUnaryInterceptor)
	// sInterceptor := grpc.StreamInterceptor(logStreamInterceptor)

	// With logging & basic auth
	// uInterceptor := withServerUnaryInterceptor()
	// sInterceptor := withServerStreamInterceptor()
	// s := grpc.NewServer(uInterceptor, sInterceptor)

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	pb.RegisterFibonacciServer(s, service.NewFibonacciServer())

	grpc_prometheus.Register(s)
	http.Handle("/metrics", promhttp.Handler())

	log.Infof("Starting diagnostice server on port %s...\n", diagnosticsPort)
	go http.ListenAndServe(":"+diagnosticsPort, nil)

	log.Infof("Starting fib server on port %s...\n", grpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
