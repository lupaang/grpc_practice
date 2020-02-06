package main

import (
	"log"

	"context"

	pb "./proto/test"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:60000"
	defaultFileName = "number.json"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot dial : %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	r, err := client.Add(context.Background(), numbers)

	if err != nil {
		log.Fatalf("Could not greet : %v", err)
	}

	log.Printf("Created : %t", r)
}
