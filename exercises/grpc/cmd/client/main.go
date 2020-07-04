package main

import (
	"context"
	"fmt"
	"time"

	pb "algogrit.com/grpc-biblioteca/api"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":5001", grpc.WithInsecure(), grpc.WithTimeout(time.Second))

	checkError(err)

	client := pb.NewBibliotecaClient(conn)
	ctx, _ := context.WithCancel(context.Background())

	var bookID int64
	bookID = 2
	log.Info("fetching book by id: ", bookID)
	br := pb.BookRequest{ID: bookID}
	book, err := client.Show(ctx, &br)
	checkError(err)

	fmt.Println(book)
}

func checkError(err error) {
	if err != nil {
		log.Fatal("encountered: ", err)
	}
}
