package main

import (
	"crypto/tls"
	"log"
	"net/rpc"
)

func main() {
	rpc.Register(new(Cal))
	cert, _ := tls.LoadX509KeyPair("server.crt", "server.key")
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	listener, _ := tls.Listen("tcp", ":1234", config)
	log.Printf("Serving RPC server on port %d", 1234)
	for {
		conn, _ := listener.Accept()
		defer conn.Close()
		go rpc.ServeConn(conn)
	}
}

type Result struct {
	Num, Ans int
}
type Cal int

// func (t *T) MethodName(argType T1, replyType *T2) error
func (cal *Cal) Square(num int, result *Result) error {
	result.Ans = num
	result.Num = num * num
	return nil
}
