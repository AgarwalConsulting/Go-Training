package main

import (
	"log"
	"net"
	"os"

	pb "algogrit.com/hello-grpc/hello"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50001")

	checkError(err)

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &HelloService{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Printf("Error encountered: %v\n", err)
		os.Exit(1)
	}
}
