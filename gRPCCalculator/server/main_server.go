package main

import (
	"context"
	"log"
	"net"

	pb "../calculator-proto"

	"google.golang.org/grpc"
)

const (
	port = ":60000"
)

type calculatorServiceServer struct {
}

func (s *calculatorServiceServer) Add(ctx context.Context, numbers *pb.Numbers) (*pb.Response, error) {
	sum := numbers.A + numbers.B

	return &pb.Response{Created: true, Result: float64(sum)}, nil
}

func (s *calculatorServiceServer) Subtract(ctx context.Context, numbers *pb.Numbers) (*pb.Response, error) {
	subtract := numbers.A - numbers.B

	return &pb.Response{Created: true, Result: float64(subtract)}, nil
}

func (s *calculatorServiceServer) Multiply(ctx context.Context, numbers *pb.Numbers) (*pb.Response, error) {
	var multiply int32

	multiply = numbers.A * numbers.B

	return &pb.Response{Created: true, Result: float64(multiply)}, nil
}

func (s *calculatorServiceServer) Divide(ctx context.Context, numbers *pb.Numbers) (*pb.Response, error) {
	divide := int32(0)

	if numbers.B == 0 {
		return &pb.Response{Created: true, Result: float64(divide)}, nil
	}

	divide = numbers.A / numbers.B
	return &pb.Response{Created: true, Result: float64(divide)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &calculatorServiceServer{})

	log.Println("Running on port :", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
