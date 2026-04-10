package main

import (
	"context"
	"io"
	"log"

	pb "github.com/waves2k/go-simple-grpc/proto"
)

func callSayHelloServerStreaming(client pb.SimpleServiceClient, names *pb.ListOfNames) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming :%v", err)
		}
		log.Println(res.Message)
		log.Println(res.String())
		log.Println(res.GetMessage())
	}
	log.Println("Streaming finished")
}
