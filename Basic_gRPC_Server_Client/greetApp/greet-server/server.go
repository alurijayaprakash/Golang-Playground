package main

import (
	"basicgreet/greetApp/greetpb"
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

const (
	PORT = ":8080"
)

type myserver struct {
	greetpb.GreetServiceServer
}

// Unary Streaming
func (s *myserver) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	// extrat info from request
	// don't directly access fields , only access through getters
	fmt.Println("Greet Func was invoked with req", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

// Server Streaming
func (s *myserver) GreetMany(req *greetpb.GreetManyRequest, stream greetpb.GreetService_GreetManyServer) error {
	fmt.Println("GreetMany Func was invoked with req", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 1; i <= 10; i++ {
		result := "Hello " + firstName + " Number is " + strconv.Itoa(i)
		res := &greetpb.GreetManyResponse{
			Result: result,
		}
		stream.Send(res)
		fmt.Println("GreetMany Func Sent ==> ", result)
		time.Sleep(1 * time.Second) // time is not required here
	}
	return nil
}

// Client Streaming
func (s *myserver) LongGreet(greetpb.GreetService_LongGreetServer) error {

	return nil
}

func main() {
	fmt.Println("Greet Server init....")

	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(grpcServer, &myserver{})

	fmt.Println("Successfully Started gRPC Server at PORT is ", PORT)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC Server : %v", err)

	}

}
