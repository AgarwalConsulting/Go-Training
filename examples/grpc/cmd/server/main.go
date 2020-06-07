package main

import (
	"context"
	"net"

	pb "algogrit.com/biblioteca-grpc/biblioteca"
	repo "algogrit.com/biblioteca-grpc/pkg/repository"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var books = map[int64]pb.Book{
	1: pb.Book{Id: 1, Title: "The C Book", Author: "Dennis Ritchie"},
	2: pb.Book{Id: 2, Title: "C++", Author: "Bjarne Stroustrop"},
}

var sequence = repo.SequenceGenerator(2)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedBibliotecaServer
}

func (s *server) ShowBook(ctx context.Context, req *pb.BookQuery) (*pb.Book, error) {
	bookID := req.GetId()
	log.Info("Getting bookID: ", bookID)
	book := books[bookID]

	return &book, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBibliotecaServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
