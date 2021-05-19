package main

import (
	"context"
	"fmt"

	pb "github.com/aystzh/go-modules/grpc/example/proto"

	"google.golang.org/grpc"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {
	//连接gRPC服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	//初始化客户端
	c := pb.NewHelloClient(conn)

	reqBody := new(pb.HelloRequest)

	reqBody.Name = "gRPC"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.Message)

}
