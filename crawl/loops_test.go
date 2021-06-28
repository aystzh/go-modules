package main

import (
	"fmt"
	"testing"
)

func TestLoops(t *testing.T) {
	str := "helloword"
	for k, v := range str {
		fmt.Printf("k:%dv:%x \r\n", k, v)
	}
}

func TestChan(t *testing.T) {
	c := make(chan int)
	go func() {
		c <- 5
		c <- 9
		c <- 1
		close(c)
	}()

	for k := range c {
		fmt.Printf("chan: %d \r\n", k)
	}
}
