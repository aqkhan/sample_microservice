package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/aqkhan/sample_microservice/sample_service"
	ss "github.com/aqkhan/sample_microservice/streaming_service"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var response_counter uint64 = 0

type localSampleServiceServer struct {
	sample_service.UnimplementedSampleServiceServer
}

type streaming_server struct {}

func (s localSampleServiceServer) Create(ctx context.Context, req *sample_service.CreateRequest) (*sample_service.CreateResponse, error) {
	response_counter++
	
	return &sample_service.CreateResponse{
		DeformedMessage: 	req.Message + " <--- Sample service MARKED ||// Timestamp: " + time.Now().Format(time.RFC3339),
		ResponseCounter: 	response_counter,
		ResponseTimeStamp: 	timestamppb.Now(),
	}, nil
}

func (s streaming_server) DataStream(req *ss.StreamRequest, srv ss.StreamData_DataStreamServer) error {
	for i := 0; i <= 10; i++ {
		value := randStringBytes(128)
		resp := ss.StreamResponse {
			Part: 	uint32(i),
			Buffer: value,
		}

		if err := srv.Send(&resp); err != nil {
			log.Fatalf("Error generating response in data streaming service. Error message: %v", err)
			return err
		}
	}
	return nil
}

func randStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	fmt.Println("Hello from the Sample Microservice")

	lis, err := net.Listen("tcp", ":6969")
	if err != nil {
		log.Fatalln("Error listning to to TCP port 6969. Error message: " + err.Error())
	}
	grpcServer := grpc.NewServer()

	basic_service := &localSampleServiceServer{}
	streaming_service := &streaming_server{}

	sample_service.RegisterSampleServiceServer(grpcServer, basic_service)
	// ss.RegisterStreamDataServer(grpcServer, streaming_service)
	ss.RegisterStreamDataServer(grpcServer, streaming_service)

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalln("Error serving server registrar. Error message: " +  err.Error())
	}
}

