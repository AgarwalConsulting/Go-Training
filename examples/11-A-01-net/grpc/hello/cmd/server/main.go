package main

import (
	"log"
	"net"

	pb "algogrit.com/hello-grpc/api"
	"algogrit.com/hello-grpc/greeting/service"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":12345")

	if err != nil {
		log.Fatalln("failed to create listener:", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreetingServer(s, service.NewV1())

	log.Println("Starting greeting server on port: 12345...")
	if err := s.Serve(lis); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
