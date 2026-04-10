package main

import (
	"log"
	"net"

	pb "github.com/waves2k/go-simple-grpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.SimpleServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleServiceServer(grpcServer, &helloServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start grpc Server: %v", err)
	}
}
