package main

import (
	"context"
	"fmt"
	"time"

	pb "algogrit.com/grpc-biblioteca/api"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connCtx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	conn, err := grpc.DialContext(connCtx, ":5001", grpc.WithTransportCredentials(insecure.NewCredentials()))

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
