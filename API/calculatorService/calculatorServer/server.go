package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"../calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedSumServiceServer
}

func (s *server) Calculator(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Request received: %v", req)
	num1 := req.GetNumbers().GetNum1()
	num2 := req.GetNumbers().GetNum2()

	res := &calculatorpb.SumResponse{
		Sum: &calculatorpb.Sum{
			Sum: num1 + num2,
		},
	}
	return res, nil
}

// func (ugs) mustEmbedUnimplementedGreetServiceServer() {}

func main() {
	fmt.Println("Server booting up...")

	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("Failed server to listen: %v", err)
	}

	s := grpc.NewServer()
	// server := greetpb.UnimplementedGreetServiceServer{}
	calculatorpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
