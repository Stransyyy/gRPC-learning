package client

import (
	"context"
	"log"
	"time"

	"github.com/Stransyyy/gRPC-learning/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunClient() {

	creds := insecure.NewCredentials()

	trprtCred := grpc.WithTransportCredentials(creds)

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8080", trprtCred, grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new example service client
	client := pb.NewExampleServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Send request
	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Stransyyy"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Print response
	log.Printf("Greeting: %s", r.Message)
}
