package main

import (
	"net"

	pb "algogrit.com/grpc-biblioteca/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":5001")

	checkError(err)

	s := grpc.NewServer()

	pb.RegisterBibliotecaServer(s, NewBibliotecaServer())

	err = s.Serve(lis)

	checkError(err)
}

func checkError(err error) {
	if err != nil {
		logrus.Fatal("encountered: ", err)
	}
}
