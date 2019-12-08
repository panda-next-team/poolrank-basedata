package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/panda-next-team/poolrank-proto/basedata"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address string = "localhost:80"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAlgorithmServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CountAlgorithms(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("could not : %v", err)
	}
	log.Printf("Count: %d", r.Count)
}
