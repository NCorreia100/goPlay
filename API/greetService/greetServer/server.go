package main

import (
	"fmt"
	"log"
	"net"

	"../greetpb"
	"google.golang.org/grpc"
)

// type server struct{}

// func (ugs) mustEmbedUnimplementedGreetServiceServer() {}

func main() {
	fmt.Println("Server booting up...")

	lis, err := net.Listen("tcp", "0.0.0..0:50051")
	if err != nil {
		log.Fatalf("Failed server to listen: %v", err)
	}

	s := grpc.NewServer()
	unsafe := greetpb.UnimplementedGreetServiceServer{}
	greetpb.RegisterGreetServiceServer(s, unsafe)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
