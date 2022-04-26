package main

import (
	"corpuser/server"
	pb "corpuser/user"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const addr = ":50051"

func main() {
	fmt.Println("Hello from Server")
	log.Println("Starting gRPC server")
	StartGrpcServer(addr)

}
func StartGrpcServer(addr string) {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen : %v \n", err)
	}

	// creating a new server
	grpcServer := grpc.NewServer()

	s := server.NewService()

	pb.RegisterUserMethodServer(grpcServer, s)

	log.Printf("GRPC server running on %v\n", addr)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Falied to serve: %v", err)
	}

}
