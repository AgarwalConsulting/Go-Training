package main

import (
	"context"
	"io"
	"log"

	pb "algogrit.com/hello-grpc/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	log.Println("Connecting to greeting server on: 12345...")

	ctx, _ := context.WithCancel(context.Background())

	conn, err := grpc.DialContext(ctx, ":12345", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("Unable to connect:", err)
	}
	defer conn.Close()

	c := pb.NewGreetingClient(conn)

	msg := pb.Hello{Name: "World"}
	res, err := c.Greet(ctx, &msg)

	if err != nil {
		log.Fatalln("Unable to greet:", err)
	}

	log.Println("Response:", res.Value)

	log.Println("Streaming from server:")
	streamClient, err := c.Stream(ctx, &emptypb.Empty{})

	if err != nil {
		log.Fatalln("Unable to stream:", err)
	}

	for {
		res, err := streamClient.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Unable to receive:", err)
		}

		log.Println("\tReceived:", res.Text)
	}
}
