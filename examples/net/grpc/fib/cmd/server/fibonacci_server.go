package main

import (
	"context"
	"time"

	pb "algogrit.com/fib-grpc/fibonacci"
	gen "algogrit.com/fib-grpc/pkg/generator"
	empty "github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
)

type fibonacciServer struct {
	gen *gen.FibonacciGenerator
}

func (s *fibonacciServer) GetNext(ctx context.Context, e *empty.Empty) (*pb.FibonacciNumber, error) {
	return &pb.FibonacciNumber{Value: s.gen.NextValue()}, nil
}

func (s *fibonacciServer) GetFirstN(ctx context.Context, n *pb.FirstNQuery) (*pb.FirstNResponse, error) {
	s.gen.Reset()
	// values := []*pb.FibonacciNumber{}
	values := make([]*pb.FibonacciNumber, 0, n.N)

	var i int64

	for i = 0; i < n.N; i++ {
		values = append(values, &pb.FibonacciNumber{Value: s.gen.NextValue()})
	}

	return &pb.FirstNResponse{Values: values}, nil
}

func (s *fibonacciServer) Stream(e *empty.Empty, stream pb.Fibonacci_StreamServer) error {
	for {
		nextVal := &pb.FibonacciNumber{Value: s.gen.NextValue()}

		if nextVal.Value < 0 { // Integer overflow
			break
		}

		stream.Send(nextVal)

		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

func NewFibonacciServer() *fibonacciServer {
	log.Info("Initializing a new server...")
	return &fibonacciServer{gen: gen.NewFibonacciGenerator()}
}
