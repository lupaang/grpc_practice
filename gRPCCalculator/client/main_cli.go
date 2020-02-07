package main

import (
	"log"

	"context"

	pb "../calculator-proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:60000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot dial : %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)

	r, err := client.Add(context.Background(), &pb.Numbers{3, 4})

	if err != nil {
		log.Fatalf("Could not greet : %v", err)
	}

	log.Printf("Calculated : %t", r)
}
