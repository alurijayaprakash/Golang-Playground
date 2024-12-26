package main

import (
	"calc/CalcApp/calcpb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type myserver struct {
	calcpb.CalcServiceServer
}

func (s *myserver) GetSum(ctx context.Context, req *calcpb.CalcRequest) (*calcpb.CalcResponse, error) {

	fmt.Println("GetSum Invoked")
	num1 := req.GetNum1()
	num2 := req.GetNum2()

	sumX := num1 + num2
	fmt.Printf("GetSum Func ==> Given Num1 : %v Num2 : %v then Sum is : %v ", num1, num2, sumX)
	fmt.Println()
	result := &calcpb.CalcResponse{
		Result: sumX,
	}
	return result, nil
}

func (s *myserver) GetMultiply(ctx context.Context, req *calcpb.CalcRequest) (*calcpb.CalcResponse, error) {

	fmt.Println("GetMultiply Invoked")
	num1 := req.GetNum1()
	num2 := req.GetNum2()

	sumX := num1 * num2
	fmt.Printf("GetMultiply Func ==> Given Num1 : %v Num2 : %v then Result is : %v ", num1, num2, sumX)
	fmt.Println()
	result := &calcpb.CalcResponse{
		Result: sumX,
	}
	return result, nil
}

func (s *myserver) GetDivision(ctx context.Context, req *calcpb.CalcRequest) (*calcpb.CalcResponse, error) {

	fmt.Println("GetDivision Invoked")
	num1 := req.GetNum1()
	num2 := req.GetNum2()

	sumX := num1 / num2
	fmt.Printf("GetDivision Func ==> Given Num1 : %v Num2 : %v then Result is : %v ", num1, num2, sumX)
	fmt.Println()
	result := &calcpb.CalcResponse{
		Result: sumX,
	}
	return result, nil
}

func (s *myserver) GetSubtract(ctx context.Context, req *calcpb.CalcRequest) (*calcpb.CalcResponse, error) {

	fmt.Println("GetSubtract Invoked")
	num1 := req.GetNum1()
	num2 := req.GetNum2()

	sumX := num1 - num2
	fmt.Printf("GetSubtract Func ==> Given Num1 : %v Num2 : %v then Result is : %v ", num1, num2, sumX)
	fmt.Println()
	result := &calcpb.CalcResponse{
		Result: sumX,
	}
	return result, nil
}

func main() {
	fmt.Println("Server init....!")
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	fmt.Printf("Server Listen at Port %v", lis.Addr())
	fmt.Println()
	grpcServer := grpc.NewServer()

	calcpb.RegisterCalcServiceServer(grpcServer, &myserver{})

	fmt.Println("Successfully Started gRPC Server")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC Server : %v", err)
	}

}
