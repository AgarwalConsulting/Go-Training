package main

import (
	"context"
	"encoding/base64"
	"net"
	"time"

	pb "algogrit.com/fib-grpc/fibonacci"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// uInterceptor := withServerUnaryInterceptor()
	// sInterceptor := grpc.StreamInterceptor(streamInterceptor)

	s := grpc.NewServer()
	pb.RegisterFibonacciServer(s, NewFibonacciServer())

	log.Info("Starting fib server on port ", port, "...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

const username = "Alice"
const password = "password123"

//AuthFunc is a middleware (interceptor) that extracts token from header
func AuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "basic")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "no basic header found: %v", err)
	}
	auth := username + ":" + password
	authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))
	if token != authEncoded {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth credentials: %v", err)
	}
	newCtx := context.WithValue(ctx, "Authenticated", true)
	return newCtx, nil
}

// type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)

func unaryInterceptor(ctx context.Context,
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
	// return grpc.UnaryInterceptor(unaryInterceptor)
	return grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
		unaryInterceptor,
		// grpc_auth.UnaryServerInterceptor(AuthFunc),
	))
}

// type StreamServerInterceptor func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error

func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()

	log.Info("Streaming: ", info.FullMethod)

	err := handler(srv, ss)

	log.Infof("Request - Method: %s\tDuration:%s\tError:%v\n", info.FullMethod, time.Since(start), err)

	return err
}
