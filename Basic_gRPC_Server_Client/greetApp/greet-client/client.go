package main

import (
	"basicgreet/greetApp/greetpb"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	PORT = ":8080"
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
	fmt.Println("Response from Greet :", resp)
}

func doServerStream(conn greetpb.GreetServiceClient) {
	fmt.Println("doServerStream Func init...!")
	req := &greetpb.GreetManyRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "JayaPrakash",
			LastName:  "Aluri",
		},
	}
	respStream, err := conn.GreetMany(context.Background(), req)
	if err != nil {
		fmt.Println("Error While calling GreetMany RPC : ", err)
	}
	for {
		msg, err := respStream.Recv()
		if err == io.EOF {
			fmt.Println("Stream Completed")
			break
		}
		if err != nil {
			fmt.Println("Error While reading stream Recv() : ", err)
		}
		fmt.Println("Response from GreetMany ==> ", msg.GetResult())
	}
}

func main() {
	fmt.Println("Client init......")
	// cc = Client Connection
	cc, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client Could not connect to gRPC Server : %v \n", err)
	}
	defer cc.Close()
	conn := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Client created successfully and connected port is, %v \n", PORT)

	// Call Unary func
	doUnary(conn)

	// Call Server Stream func
	doServerStream(conn)
}
