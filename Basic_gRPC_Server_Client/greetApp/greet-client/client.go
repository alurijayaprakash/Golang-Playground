package main

import (
	"basicgreet/greetApp/greetpb"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	PORT = ":8080"
)

// Unary Streaming

// func doUnary(conn greetpb.GreetServiceClient) {
// 	fmt.Println("doUnary Func init...!")
// 	req := &greetpb.GreetRequest{
// 		Greeting: &greetpb.Greeting{
// 			FirstName: "JayaPrakash",
// 			LastName:  "Aluri",
// 		},
// 	}
// 	resp, err := conn.Greet(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("Error While calling Greet func: %v", err)
// 	}
// 	fmt.Println("Response from Greet :", resp)
// }

// Server Streaming

// func doServerStream(conn greetpb.GreetServiceClient) {
// 	fmt.Println("doServerStream Func init...!")
// 	req := &greetpb.GreetManyRequest{
// 		Greeting: &greetpb.Greeting{
// 			FirstName: "JayaPrakash",
// 			LastName:  "Aluri",
// 		},
// 	}
// 	respStream, err := conn.GreetMany(context.Background(), req)
// 	if err != nil {
// 		fmt.Println("Error While calling GreetMany RPC : ", err)
// 	}
// 	for {
// 		msg, err := respStream.Recv()
// 		if err == io.EOF {
// 			fmt.Println("Stream Completed")
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("Error While reading stream Recv() : ", err)
// 		}
// 		fmt.Println("Response from GreetMany ==> ", msg.GetResult())
// 	}
// }

// Client Streaming

// func doClientStream(c greetpb.GreetServiceClient) {
// 	fmt.Println("doClientStream Func init...!")
// 	requests := []*greetpb.LongGreetRequest{
// 		{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "JayaprakashA",
// 				LastName:  "Aluri",
// 			},
// 		},
// 		{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "SonyB",
// 				LastName:  "India",
// 			},
// 		}, {
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "GoogleC",
// 				LastName:  "Inc",
// 			},
// 		}, {
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "MicrosoftD",
// 				LastName:  "365",
// 			},
// 		}, {
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "ChatE",
// 				LastName:  "GPT",
// 			},
// 		}, {
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "OneplusF",
// 				LastName:  "Nord2",
// 			},
// 		},
// 	}
// 	stream, err := c.LongGreet(context.Background())
// 	if err != nil {
// 		log.Fatal("Error while calling LongGreet ", err)
// 	}
// 	// loop over requests slice and send message individually to server
// 	for _, req := range requests {
// 		fmt.Println("Sending req", req)
// 		stream.Send(req)
// 		time.Sleep(1 * time.Second)
// 	}
// 	resp, err := stream.CloseAndRecv()
// 	if err != nil {
// 		log.Fatal("Error while receving response from LongGreet ", err)
// 	}
// 	fmt.Println("LognGreet Response is ", resp)
// }

// Bi-Directional Streaming
func doBiDirectStream(c greetpb.GreetServiceClient) {
	fmt.Println("doBiDirectStream Func init...!")
	allReqs := []*greetpb.GreetEveryoneRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Jayaprakash",
				LastName:  "Aluri",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Sony",
				LastName:  "India",
			},
		}, {
			Greeting: &greetpb.Greeting{
				FirstName: "Google",
				LastName:  "Inc",
			},
		}, {
			Greeting: &greetpb.Greeting{
				FirstName: "Microsoft",
				LastName:  "365",
			},
		}, {
			Greeting: &greetpb.Greeting{
				FirstName: "Chat",
				LastName:  "GPT",
			},
		}, {
			Greeting: &greetpb.Greeting{
				FirstName: "Oneplus",
				LastName:  "Nord2",
			},
		},
	}
	// we create a stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatal("Error while calling GreetEveryone ", err)
		return
	}
	waitCh := make(chan struct{})
	// we send a bunch of messages of the client (goroutine)
	go func() {
		// func to send a bunch of messages
		for _, req := range allReqs {
			fmt.Println("Sending message", req)
			stream.Send(req)
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
				// fmt.Println("Stream Completed")
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
	cc, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client Could not connect to gRPC Server : %v \n", err)
	}
	defer cc.Close()
	conn := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Client created successfully and connected port is, %v \n", PORT)

	// // Call Unary func
	// doUnary(conn)

	// // Call Server Stream func
	// doServerStream(conn)

	// // Call Client Stream func
	// doClientStream(conn)

	// Call Bi-Directional Stream func
	doBiDirectStream(conn)
}
