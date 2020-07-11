package main

import (
	"context"
	"io"
	"os"
	"time"

	pb "algogrit.com/fib-grpc/api"
	empty "github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func getFirst10(c pb.FibonacciClient) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	log.Info("First 10 fibonacci numbers...")
	r, err := c.GetFirstN(ctx, &pb.FirstNQuery{N: 10})

	if err != nil {
		log.Fatalf("could not get first 10 fib numbers: %v", err)
	}

	// log.Debug("Numbers: ", r)
	for _, el := range r.Values {
		log.Info(el.Value)
	}
}

func getNext(c pb.FibonacciClient) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	nextVal, err := c.GetNext(ctx, &empty.Empty{})

	if err != nil {
		log.Fatalf("could not get next number: %v", err)
	}

	log.Info("Next number: ", nextVal.Value)
}

func streamValues(c pb.FibonacciClient) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	// auth := "algogrit:supersecurepassword"
	// authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))
	// md := metadata.Pairs("authorization", "basic "+authEncoded)
	// ctx = metadata.NewOutgoingContext(ctx, md)

	stream, _ := c.Stream(ctx, &empty.Empty{})

	for {
		nextVal, err := stream.Recv()

		if err != nil {
			if err != io.EOF {
				log.Errorf("Error streaming: %v\n", err)
			}
			break
		}

		log.Info("Received next number: ", nextVal.Value)
	}
}

var address string

func init() {
	var ok bool

	address, ok = os.LookupEnv("SERVER_ADDRESS")
	if !ok {
		address = ":50001"
	}
}

func main() {
	log.Infof("Connecting to fib server on... %s\n", address)
	//// With TLS Authentication
	// creds, _ := credentials.NewClientTLSFromFile("cert.pem", "")
	// tlsOption := grpc.WithTransportCredentials(creds)

	// Set up a connection to the server.
	conn, err := grpc.Dial(address,
		grpc.WithTimeout(time.Second*10), grpc.WithBlock(),
		grpc.WithInsecure(),
		// tlsOption,
		// grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		// grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
	)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Info("Connected!")
	defer conn.Close()
	c := pb.NewFibonacciClient(conn)

	getFirst10(c)
	getNext(c)
	streamValues(c)
}
