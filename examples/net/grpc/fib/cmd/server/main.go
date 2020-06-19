package main

import (
	"net"

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

	s := grpc.NewServer()
	pb.RegisterFibonacciServer(s, NewFibonacciServer())

	log.Info("Starting fib server on port ", port, "...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
