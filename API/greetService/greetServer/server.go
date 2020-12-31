package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"../greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Request received: %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := greetpb.Greeting{
		FirstName: firstName,
	}
	res := &greetpb.GreetResponse{
		Result: &result,
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
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
