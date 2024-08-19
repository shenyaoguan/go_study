package main

import (
	"context"
	"flag"
	hello_grpc "gRPC/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "localhost:50051", "server address")
	name = flag.String("name", "你好", "name")
)

func main() {
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.NewClient err: %v", err)
	}
	defer conn.Close()
	c := hello_grpc.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHi(ctx, &hello_grpc.Req{Message: *name})
	if err != nil {
		log.Fatalf("c.SayHi err: %v", err)
	}
	log.Printf("Response: %s", r.GetMessage())
}
