package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/waves2k/go-simple-grpc/proto"
)

func callSayHelloBidirectionalStreaming(client pb.SimpleServiceClient, names *pb.ListOfNames) {
	log.Printf("Bidirectional streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	wait := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error whiel streaming: %v\n", err)
			}
			log.Println(message)
		}
		close(wait)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error whiel sending the request: %v\n", err)
		}
		time.Sleep(time.Second * 2)
	}
	stream.CloseSend()
	<-wait
	log.Println("Bidirectional streaming finished")
}
