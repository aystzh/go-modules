package main

import (
	"fmt"
	"os"
	"testing"
)

type calcCase struct {
	A, B, Excepted int
}

func setup() {
	fmt.Println("测试前初始化方法")
}

func teardown() {
	fmt.Println("测试后执行的方法")
}

func createMulTestCase(t *testing.T, c *calcCase) {
	t.Helper() //用于标注该函数是帮助函数，报错时将输出帮助函数调用者的信息，而不是帮助函数的内部信息。
	if ans := Mul(c.A, c.B); ans != c.Excepted {
		t.Fatalf("%d * %d expected %d, but %d got",
			c.A, c.B, c.Excepted, ans)
	}
}

func TestMul(t *testing.T) {
	fmt.Println("createMulTestCase execute")
	createMulTestCase(t, &calcCase{2, 3, 6})
	createMulTestCase(t, &calcCase{2, 2, 4})
	createMulTestCase(t, &calcCase{2, 0, 0})
}

func TestAdd(t *testing.T) {
	fmt.Println("testAdd execute")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
