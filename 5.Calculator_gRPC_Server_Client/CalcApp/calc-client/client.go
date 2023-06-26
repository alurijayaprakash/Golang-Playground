package main

import (
	"calc/CalcApp/calcpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func sumFunc(conn calcpb.CalcServiceClient) {
	fmt.Println("sumFunc Func init...!")
	req := &calcpb.CalcRequest{
		Num1: 11,
		Num2: 22,
	}
	resp, err := conn.GetSum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling Calc func: %v", err)
	}
	log.Printf("sumFunc ==> Response from Calc : %v", resp)
}

func MultiplyFunc(conn calcpb.CalcServiceClient) {
	fmt.Println("MultiplyFunc Func init...!")
	req := &calcpb.CalcRequest{
		Num1: 11,
		Num2: 22,
	}
	resp, err := conn.GetMultiply(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling GetMultiply: %v", err)
	}
	log.Printf("MultiplyFunc ==> Response from Calc : %v", resp)
}

func DivisionFunc(conn calcpb.CalcServiceClient) {
	fmt.Println("DivisionFunc Func init...!")
	req := &calcpb.CalcRequest{
		Num1: 11,
		Num2: 22,
	}
	resp, err := conn.GetDivision(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling GetDivision: %v", err)
	}
	log.Printf("MultiplyFunc ==> Response from Calc : %v", resp)
}

func SubtractFunc(conn calcpb.CalcServiceClient) {
	fmt.Println("SubtractFunc Func init...!")
	req := &calcpb.CalcRequest{
		Num1: 11,
		Num2: 22,
	}
	resp, err := conn.GetSubtract(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling GetSubtract: %v", err)
	}
	log.Printf("SubtractFunc ==> Response from Calc : %v", resp)
}

func main() {
	fmt.Println("Client init.....!")
	// cc = Client Connection
	cc, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client Could not connect to gRPC Server : %v", err)
	}
	defer cc.Close()
	conn := calcpb.NewCalcServiceClient(cc)
	fmt.Printf("Client created successfully : %f", conn)

	// Call Unary func

	sumFunc(conn)
	MultiplyFunc(conn)
	DivisionFunc(conn)
	SubtractFunc(conn)

}
