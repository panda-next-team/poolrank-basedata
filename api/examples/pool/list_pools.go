package main

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	pb "github.com/panda-next-team/poolrank-proto/basedata"
	"google.golang.org/grpc"
	"log"
	"math"
	"time"
)

const (
	address string = "localhost:3001"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt64),
		grpc.MaxCallSendMsgSize(math.MaxInt64)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPoolServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	queries := make([]*pb.ListPoolsRequest_Query, 1)
	query := new(pb.ListPoolsRequest_Query)
	query.Value, _ = ptypes.MarshalAny(&pb.ListPoolsRequest_QueryBoolValue{Value:true})
	query.Operator = pb.ListPoolsRequest_EQ
	query.Field = pb.ListPoolsRequest_Q_STATUS
	queries[0] = query

	r, err := c.ListPools(ctx, &pb.ListPoolsRequest{Queries: queries, Skip: 0, Limit: 10})
	if err != nil {
		log.Fatalf("could not : %v", err)
	}
	log.Printf("entity: %v", r)

}
