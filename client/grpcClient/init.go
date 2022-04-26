package grpcclient

import (
	"log"
	"os"
)

var address = "0.0.0.0:50051"

func init() {
	log.Println("Getting user_endpoint")
	if ep, ok := os.LookupEnv("USER_ENDPOINT"); ok {
		address = ep
	}
	NewGrpcClient(address)
}

func Get() *Client {
	return s
}
