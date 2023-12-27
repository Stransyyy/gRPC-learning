package server

import (
	"context"
	"fmt"

	pb "github.com/stransy/grpc-learning/client/home/stransy/gRPC-learning/client"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Hello, I am in the SayHello function")
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}
