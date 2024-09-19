package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
	"github.com/aqkhan/sample_microservice/sample_service"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/grpc"
)

var response_counter uint64 = 0
type localSampleServiceServer struct {
	sample_service.UnimplementedSampleServiceServer
}

func (s localSampleServiceServer) Create(ctx context.Context, req *sample_service.CreateRequest) (*sample_service.CreateResponse, error) {
	response_counter++
	log.Printf("Response counter: %d", response_counter)
	return &sample_service.CreateResponse{
		DeformedMessage: 	req.Message + " <--- Sample service MARKED ||// Timestamp: " + time.Now().Format(time.RFC3339),
		ResponseCounter: 	response_counter,
		ResponseTimeStamp: 	timestamppb.Now(),
	}, nil
}

func main() {
	fmt.Println("Hello from the Sample Microservice")

	lis, err := net.Listen("tcp", ":6969")
	if err != nil {
		log.Fatalln("Error listning to to TCP port 6969. Error message: " + err.Error())
	}
	serverRegistrar := grpc.NewServer()

	service := &localSampleServiceServer{}

	sample_service.RegisterSampleServiceServer(serverRegistrar, service)

	err = serverRegistrar.Serve(lis)

	if err != nil {
		log.Fatalln("Error serving server registrar. Error message: " +  err.Error())
	}
}

