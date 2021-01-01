package main

import (
	"context"
	"fmt"
	"log"

	"../calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client!")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	c := calculatorpb.NewSumServiceClient(cc)
	doUnaryRequest(c)
	defer cc.Close()
}

func doUnaryRequest(c calculatorpb.SumServiceClient) {
	// fmt.Printf("Created Client: %v", c)
	req := &calculatorpb.SumRequest{
		Numbers: &calculatorpb.Numbers{
			Num1: 5,
			Num2: 7,
		},
	}

	res, err := c.Calculator(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to get response from server: %v", err)
	} else {
		log.Printf("Server response: %v", res.Sum)
	}

}
