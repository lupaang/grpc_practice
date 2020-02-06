package main

import (
	"context"
	"log"
	"net"

	pb "./proto/test"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":60000"
)

func (s *service) Add(ctx context.Context, numbers *pb.Numbers) (pb.Response, error) {
	if err != nil {
		return nil, err
	}

	sum := 0
	for _, nums := range numbers {
		sum += nums
	}

	return &pb.Response{Created: true, Result: sum, exception: "None"}
}

func (s *service) Subtract(ctx context.Context, numbers *pb.Numbers) (pb.Response, error) {
	if err != nil {
		return nil, err
	}

	subtract := 0
	for _, nums := range numbers {
		subtract -= nums
	}
	return s & pb.Response{Created: true, Result: subtract, exception: "None"}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	reflection.Register(s)

	log.Println("Running on port :", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
