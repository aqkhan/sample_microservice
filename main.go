package main

import (
	"fmt"
	"time"
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
}

