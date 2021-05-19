package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/aystzh/go-modules/grpc/example/proto"
	"google.golang.org/grpc"
)

const (
	Address = "127.0.0.1:50052"
)

type helloService struct{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message = "hello" + in.Name + "."
	return resp, nil
}

var HelloServer = helloService{}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Printf("监听端口失败！：%v", err)
	}

	o := grpc.NewServer()
	
	pb.RegisterHelloServer(o, HelloServer)
	fmt.Println("Listen on" + Address)

	o.Serve(listen)
}
