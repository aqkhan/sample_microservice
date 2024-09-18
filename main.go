package main

import (
	"fmt"
	"time"
)

type SampleMessage struct {
	Message 			string
	TimeStamp			time.Time
}

func main() {
	fmt.Println("Hello from the Sample Microservice")
}

