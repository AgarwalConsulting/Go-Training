package main

import (
	"context"

	pb "algogrit.com/hello-grpc/hello"
)

type HelloService struct {
	// pb.UnimplementedHelloServiceServer
}

func (h HelloService) Greeting(ctx context.Context, r *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	return &pb.GreetingResponse{Greeting: "Hello, " + r.Name}, nil
}
