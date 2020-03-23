package main

import "fmt"

func main() {
	cal := new(Cal)
	res := cal.Square(12)
	fmt.Println("ans :", res.Ans)

	calRpc := new(Cal)
	var result Result
	calRpc.SquareRpc(11, &result) //引用传递
	fmt.Printf("SquareRpc %d", result.Num)
}

type Result struct {
	Num, Ans int
}

type Cal int

//不是rpc
func (cal *Cal) Square(num int) *Result {
	return &Result{
		Num: num,
		Ans: num * num,
	}
}

// func (t *T) MethodName(argType T1, replyType *T2) error
func (cal *Cal) SquareRpc(num int, result *Result) error {
	result.Ans = num
	result.Num = num * num
	return nil
}
