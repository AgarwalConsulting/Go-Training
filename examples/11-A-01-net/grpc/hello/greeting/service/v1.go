package service

import (
	"context"
	"log"
	"time"

	pb "algogrit.com/hello-grpc/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

type greetingSvcV1 struct {
	pb.UnimplementedGreetingServer
}

func (svc greetingSvcV1) Greet(ctx context.Context, in *pb.Hello) (out *pb.GreetResponse, err error) {
	var gr pb.GreetResponse

	gr.Value = "Hello, " + in.Name

	return &gr, nil
}

func (svc greetingSvcV1) Stream(in *emptypb.Empty, streamSvr pb.Greeting_StreamServer) error {
	msg := []string{"One", "Two", "Three", "Four"}

	for _, el := range msg {
		res := pb.StreamResponse{Text: el}
		err := streamSvr.Send(&res)

		if err != nil {
			log.Fatalln("Unable to send:", err)
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}

func NewV1() pb.GreetingServer {
	return greetingSvcV1{}
}
