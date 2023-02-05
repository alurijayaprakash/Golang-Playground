package main

import (
	"basicgreet/greetApp/greetpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func doUnary(conn greetpb.GreetServiceClient) {
	fmt.Println("doUnary Func init...!")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "JayaPrakash",
			LastName:  "Aluri",
		},
	}
	resp, err := conn.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling Greet func: %v", err)
	}
	log.Printf("Response from Greet : %v", resp)
}

func main() {
	fmt.Println("Client init......")
	// cc = Client Connection
	cc, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client Could not connect to gRPC Server : %v", err)
	}
	defer cc.Close()
	conn := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Client created successfully : %f", conn)

	// Call Unary func
	doUnary(conn)
}
