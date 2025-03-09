package main

import (
	"fmt"
	"log"
	"net"
	"primegrpc/PrimeApp/primepb"
	"time"

	"google.golang.org/grpc"
)

const (
	PORT = ":8081"
)

type myserver struct {
	primepb.PrimeServiceServer
}

func (s *myserver) PrimeMany(req *primepb.PrimeRequest, stream primepb.PrimeService_PrimeManyServer) error {
	fmt.Println("PrimeMany Func was invoked with req", req)
	givenNum := req.GetNum()
	k := 2
	resp := &primepb.PrimeResponse{}
	for givenNum > 1 {
		if int(givenNum)%k == 0 {
			// K is factor of givenNum
			resp.Result = int32(k)
			stream.Send(resp)
			fmt.Println("PrimeMany Func Sent ==> ", resp.GetResult())
			time.Sleep(1 * time.Second) // time is not required here
			givenNum = givenNum / int32(k)
		} else {
			k = k + 1
		}
	}
	return nil
}

func main() {
	fmt.Println("Prime Server init....")

	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()

	primepb.RegisterPrimeServiceServer(grpcServer, &myserver{})

	fmt.Println("Successfully Started gRPC Server at PORT is ", PORT)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC Server : %v", err)

	}

}
