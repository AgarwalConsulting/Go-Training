package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "algogrit.com/hello-grpc/hello"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50001", grpc.WithInsecure(), grpc.WithTimeout(time.Second))

	checkError(err)
	client := pb.NewHelloServiceClient(conn)

	gr := pb.GreetingRequest{Name: "Gaurav"}

	ctx, _ := context.WithCancel(context.Background())

	response, err := client.Greeting(ctx, &gr)
	checkError(err)

	fmt.Println(response)
}

func checkError(err error) {
	if err != nil {
		log.Printf("Error encountered: %v\n", err)
		os.Exit(1)
	}
}
