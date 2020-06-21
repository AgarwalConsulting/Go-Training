package main

import (
	"context"
	"net"
	"time"

	pb "algogrit.com/fib-grpc/fibonacci"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// s := grpc.NewServer()
	s := grpc.NewServer(withServerUnaryInterceptor(), grpc.StreamInterceptor(streamInterceptor))
	pb.RegisterFibonacciServer(s, NewFibonacciServer())

	log.Info("Starting fib server on port ", port, "...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)

func serverInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	// Calls the handler
	h, err := handler(ctx, req)

	log.Infof("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}

func withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

// type StreamServerInterceptor func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error

func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()

	log.Info("Streaming: ", info.FullMethod)

	err := handler(srv, ss)

	log.Infof("Request - Method: %s\tDuration:%s\tError:%v\n", info.FullMethod, time.Since(start), err)

	return err
}
