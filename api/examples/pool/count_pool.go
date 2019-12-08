package main

import (
	"github.com/golang/protobuf/ptypes/empty"
	"log"
	"time"
	"google.golang.org/grpc"
	pb "github.com/panda-next-team/poolrank-proto/basedata"
	"golang.org/x/net/context"
)

const (
	address string= "localhost:80"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPoolServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	r, err := c.CountPools(ctx, &empty.Empty{});
	if err != nil {
		log.Fatalf("could not : %v", err)
	}
	log.Printf("count: %d", r.Count)
}




