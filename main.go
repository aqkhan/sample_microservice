package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/aqkhan/sample_microservice/sample_service"
	"google.golang.org/grpc"
)

type SampleMessage struct {
	RequestCount		uint64
	ResponseCount		uint64
	Message				string
	RequestTimeStamp	time.Time
	ResponseTimeStamp	time.Time
}

func main() {
	fmt.Println("Hello from the Sample Microservice")

	lis, err := net.Listen("tcp", "6969")
	if err != nil {
		log.Fatalln("Error listning to to TCP port 6969. Error message: " + err.Error())
	}
	serverRegistrar := grpc.NewServer()

	sample_service.RegisterSampleServiceServer(s serserverRegistrar, )
}

