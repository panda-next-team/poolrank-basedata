package main

import (
	"context"
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

	r, err := c.GetAlgorithm(ctx, &pb.GetAlgorithmRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not : %v", err)
	}
	log.Printf("entity: %v", r.Algorithm)

}
