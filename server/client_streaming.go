package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/waves2k/go-simple-grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.SimpleService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.ListOfMessages{
				Messages: messages,
			})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name %v", res.Name)
		messages = append(messages, fmt.Sprintf("Hello, %s!", res.Name))
	}
}
