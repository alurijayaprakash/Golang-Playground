package main

import (
	"basicgreet/greetApp/greetpb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type myserver struct {
	greetpb.GreetServiceServer
}

func (s *myserver) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	// extraxt info from request
	// don't directly access fields , only access through getters
	fmt.Printf("Greet Func was invoked with %v ", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Greet Server init....")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(grpcServer, &myserver{})

	log.Printf("Listing by the address : %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC Server : %v", err)

	}

	log.Printf("Successfully Started gRPC Server. Address is %v and Port no %v", lis.Addr(), lis.Addr().Network())

}
