package main

import (
	"context"
	"fmt"
	"jpblog/BlogApp/blogpb"
	"log"
	"net"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

const (
	PORT = ":8081"
)

type myserver struct {
	blogpb.BlogServiceServer
}

var collection *mongo.Collection

type blogItem struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}

// https://pkg.go.dev/github.com/mongodb/mongo-go-driver/mongo?utm_source=godoc

func main() {
	// if we crash , we will get file name and line num in cmd
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("BlogService Server init....")

	// connect to mongodb
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("tmdb").Collection("jpblog")
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()

	blogpb.RegisterBlogServiceServer(grpcServer, &myserver{})

	go func() {
		fmt.Println("gRPC Server Starting at PORT is ", PORT)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start gRPC Server : %v", err)
		}
	}()

	// Wait for Control+C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	//Block untill a signal is received
	<-ch
	fmt.Println("Stopping the Server")
	grpcServer.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("End of the Program")
}
