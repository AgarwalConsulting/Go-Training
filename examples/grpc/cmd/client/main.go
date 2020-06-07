package main

import (
	"context"
	"log"
	"time"

	pb "algogrit.com/biblioteca-grpc/biblioteca"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBibliotecaClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ShowBook(ctx, &pb.BookQuery{Id: 1})

	if err != nil {
		log.Fatalf("could not book: %v", err)
	}

	log.Printf("Book: %s", r)
}
