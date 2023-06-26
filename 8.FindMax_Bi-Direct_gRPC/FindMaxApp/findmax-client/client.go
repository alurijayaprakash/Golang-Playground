package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"maxnum/FindMaxApp/findmaxpb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	PORT = ":8082"
)

// Call Bi-Directional Streaming
func doBiDirectStream(c findmaxpb.FindMaxServiceClient) {
	fmt.Println("doBiDirectStream Func init...!")
	allReqs := []*findmaxpb.FindMaxRequest{
		{
			Num: 11,
		},
		{
			Num: 45,
		},
		{
			Num: 2,
		},
		{
			Num: 66,
		},
		{
			Num: 46,
		},
		{
			Num: 112,
		},
		{
			Num: 117,
		},
		{
			Num: 1,
		},
		{
			Num: 100,
		},
	}
	// we create a stream by invoking the client
	stream, err := c.GetMaxNum(context.Background())
	if err != nil {
		log.Fatal("Error while calling GetMaxNum ", err)
		return
	}
	fmt.Println("GetMaxNum Called Successfully")
	waitCh := make(chan struct{})
	// we send a bunch of messages of the client (goroutine)
	go func() {
		// func to send a bunch of messages
		for _, req := range allReqs {
			stream.Send(req)
			fmt.Println("Sent Num", req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()
	// we receive a bunch of messages from the client (go routine)
	go func() {
		// func to receive a bunch of messages
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("Stream Completed")
				break
			}
			if err != nil {
				log.Fatalf("Error While reading stream Recv() : %v", err)
				break
			}
			fmt.Println("Received : ==> ", resp.GetResult())
		}
		close(waitCh)
	}()
	// block untill everything is done
	<-waitCh
}

func main() {
	fmt.Println("Client init......")
	// cc = Client Connection
	cc, err := grpc.Dial(PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client Could not connect to gRPC Server : %v \n", err)
	}
	defer cc.Close()
	conn := findmaxpb.NewFindMaxServiceClient(cc)
	fmt.Printf("Client created successfully and connected port is, %v \n", PORT)

	// Call Bi-Directional Stream func
	doBiDirectStream(conn)
}
