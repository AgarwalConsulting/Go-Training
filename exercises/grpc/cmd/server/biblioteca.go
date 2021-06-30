package main

import (
	"context"

	pb "algogrit.com/grpc-biblioteca/api"
)

type BibliotecaServerImpl struct {
	books      map[int64]pb.Book
	idSequence int64
	pb.UnimplementedBibliotecaServer
}

func (bs *BibliotecaServerImpl) nextID() int64 {
	bs.idSequence++
	return bs.idSequence
}

func (bs *BibliotecaServerImpl) Show(ctx context.Context, br *pb.BookRequest) (*pb.Book, error) {
	book := bs.books[br.ID]

	return &book, nil
}

func NewBibliotecaServer() pb.BibliotecaServer {
	server := &BibliotecaServerImpl{}

	server.books = map[int64]pb.Book{
		1: {ID: server.nextID(), Title: "The C Book", Author: "Dennis Ritchie"},
		2: {ID: server.nextID(), Title: "C++", Author: "Bjarne Stroustrop"},
	}

	return server
}
