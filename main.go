package main

import (
	"context"
	"flag"
	pb "golang-grpc-protobuf/protobuf/helloworld"
	"golang-grpc-protobuf/protobuf/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var (
	grpcHost = util.GetEnv("GRPC_SERVER_HOST", "192.168.0.2")
	grpcPort = util.GetEnv("GRPC_SERVER_PORT", "5000")
	addr     = flag.String("addr", grpcHost+":"+grpcPort, "the address to connect to")
	name     = flag.String("name", "default", "Name to greet")
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
