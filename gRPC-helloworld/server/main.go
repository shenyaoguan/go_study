package main

import (
	"context"
	"fmt"
	hello_grpc "gRPC/pb"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	hello_grpc.UnimplementedHelloServiceServer
}

func (s server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "Hello, I'm server"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	hello_grpc.RegisterHelloServiceServer(s, &server{})
	s.Serve(lis)
}
