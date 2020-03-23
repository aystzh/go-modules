package main

import (
	"log"
	"net/rpc"
)

func main() {
	//SyncInvoke()
	AsyncInvoke()
}

func SyncInvoke() {
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
	var result Result
	//同步调用
	if err := client.Call("Cal.SquareRpc", 12, &result); err != nil {
		log.Fatal("Failed to call Cal.Square.", err)
	}
	log.Printf("%d^2 = %d", result.Num, result.Ans)
}

//异步调用
func AsyncInvoke() {
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
	var result Result
	asyncCall := client.Go("Cal.SquareRpc", 12, &result, nil)
	log.Printf("%d^2 = %d", result.Num, result.Ans)
	<-asyncCall.Done
	log.Printf("%d^2 = %d", result.Num, result.Ans)
}

type Result struct {
	Num, Ans int
}
