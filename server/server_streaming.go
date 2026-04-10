package main

import (
	"fmt"
	"log"
	"time"

	pb "github.com/waves2k/go-simple-grpc/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.ListOfNames, stream pb.SimpleService_SayHelloServerStreamingServer) error {
	log.Printf("Got a request with names: %v", req.Names)
	for _, name := range req.Names {
		res := pb.HelloResponse{
			Message: fmt.Sprintf("Hello, %s!", name),
		}
		if err := stream.Send(&res); err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}
