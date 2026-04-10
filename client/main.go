package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/waves2k/go-simple-grpc/proto"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error to connect from client: %v", err)
	}
	defer conn.Close()

	client := pb.NewSimpleServiceClient(conn)
	names := &pb.ListOfNames{
		Names: []string{"Alex", "Victor", "Bob"},
	}

	// callSayHello(client)
	// callSayHelloServerStreaming(client, names)
	callSayHelloClientStreaming(client, names)
}
