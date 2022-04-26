package grpcclient

import (
	"log"

	"google.golang.org/grpc"

	pb "corpuser/user"
)

var s *Client

type Client struct {
	UserClient pb.UserMethodClient
}

func NewGrpcClient(address string) {
	c, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// not doing defer c.close() , it is disconnecting

	client := pb.NewUserMethodClient(c)
	s = &Client{
		UserClient: client,
	}
}
