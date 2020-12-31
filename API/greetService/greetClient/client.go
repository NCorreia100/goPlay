package main

import (
	"context"
	"fmt"
	"log"

	"../greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client!")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	c := greetpb.NewGreetServiceClient(cc)
	doUnaryRequest(c)
	defer cc.Close()
}

func doUnaryRequest(c greetpb.GreetServiceClient) {
	// fmt.Printf("Created Client: %v", c)
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Nuno",
			LastName:  "Correia",
			Age:       33,
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to get response from server: %v", err)
	} else {
		log.Printf("Server response: %v", res.Result)
	}

}
