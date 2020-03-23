package main

import (
	"log"
	"net/http"
	"net/rpc"
)

func main() {
	rpc.Register(new(Cal))
	rpc.HandleHTTP()
	log.Printf("Serving RPC server on port %d", 1234)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("error : ", err)
	}
}

type Result struct {
	Num, Ans int
}
type Cal int

// func (t *T) MethodName(argType T1, replyType *T2) error
func (cal *Cal) SquareRpc(num int, result *Result) error {
	result.Ans = num
	result.Num = num * num
	return nil
}
