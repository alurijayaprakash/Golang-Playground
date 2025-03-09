package main

import (
	"fmt"
	"io"
	"log"
	"maxnum/FindMaxApp/findmaxpb"
	"net"

	"google.golang.org/grpc"
)

const (
	PORT = ":8082"
)

type myserver struct {
	findmaxpb.FindMaxServiceServer
}

// Bi-Directional Streaming
func (s *myserver) GetMaxNum(stream findmaxpb.FindMaxService_GetMaxNumServer) error {
	fmt.Println("GetMaxNum Func was invoked with stream")
	var MaxNum int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		respNum := req.GetNum()
		fmt.Println("GetMaxNum : Received respNum is : ", respNum)
		if MaxNum < respNum {
			// we received max num
			MaxNum = respNum
			// sending calculated max num to client
			sendErr := stream.Send(&findmaxpb.FindMaxResponse{
				Result: MaxNum,
			})
			if sendErr != nil {
				log.Fatalf("Error While sending data to client %v ", err)
				return err
			}
			fmt.Println("GetMaxNum : Stream MaxNum ==> ", MaxNum)
		} else {
			continue
		}
	}
}

//main
func main() {
	fmt.Println("FindMax Server init....")

	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()

	findmaxpb.RegisterFindMaxServiceServer(grpcServer, &myserver{})

	fmt.Println("Successfully Started gRPC Server at PORT is ", PORT)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC Server : %v", err)

	}

}
