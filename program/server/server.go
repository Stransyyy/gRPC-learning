package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/Stransyyy/gRPC-learning/pb"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedExampleServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Greeting: fmt.Sprintf("Hello %s", in.GetName())}, nil
}

func main() {
	flag.Parse()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterExampleServiceServer(s, &server{})
	log.Printf("Server listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
