package main

import (
	"context"
	"log"
	"time"

	pb "github.com/waves2k/go-simple-grpc/proto"
)

func callSayHelloClientStreaming(client pb.SimpleServiceClient, names *pb.ListOfNames) {
	log.Println("Client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending: %v", err)
		}
		log.Printf("Sent the request with name: %s\n", name)
		time.Sleep(time.Second * 2)
	}

	resp, err := stream.CloseAndRecv()

	log.Println("Client streaming finished")
	if err != nil {
		log.Fatalf("Error while recieving: %v\n", err)
	}

	log.Printf("%v", resp.Messages)
}
