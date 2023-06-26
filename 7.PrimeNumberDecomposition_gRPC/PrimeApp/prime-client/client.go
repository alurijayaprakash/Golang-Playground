package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"primegrpc/PrimeApp/primepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	PORT = ":8081"
)

func doPrimeStream(conn primepb.PrimeServiceClient) {
	fmt.Println("doPrimeStream Func init...!")
	req := &primepb.PrimeRequest{
		Num: 123456,
	}
	respStream, err := conn.PrimeMany(context.Background(), req)
	if err != nil {
		fmt.Println("Error While calling PrimeMany RPC : ", err)
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
		fmt.Println("Response from PrimeMany ==> ", msg.GetResult())
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
	conn := primepb.NewPrimeServiceClient(cc)
	fmt.Printf("Client created successfully and connected port is, %v \n", PORT)

	// Call Server Stream func
	doPrimeStream(conn)
}
