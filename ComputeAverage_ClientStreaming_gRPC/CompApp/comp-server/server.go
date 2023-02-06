package main

import (
	"compavg/CompApp/comppb"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	PORT = ":8080"
)

type myserver struct {
	comppb.CompServiceServer
}

// Client Streaming
// This func will accept client streaming and return computed average value to client
func (s *myserver) CompAvg(stream comppb.CompService_CompAvgServer) error {
	fmt.Println("CompAvg Func was invoked with stream")
	sum := 0
	count := 1
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			result := sum / count
			fmt.Println("CompAvg Func finished reading from client stream and sent response to client => ", result)
			return stream.SendAndClose(&comppb.CompResponse{
				Result: int64(result),
			})
		}
		if err != nil {
			log.Fatal("Error while reading client stream", err)
		}
		sum += int(req.GetNum())
		count++
		fmt.Println("Stream input is : ", req.GetNum())
	}
}

func main() {
	fmt.Println("Server init....")

	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()

	comppb.RegisterCompServiceServer(grpcServer, &myserver{})

	fmt.Println("Successfully Started gRPC Server at PORT is ", PORT)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC Server : %v", err)

	}

}
