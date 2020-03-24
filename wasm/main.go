package main

import "syscall/js"

func main() {
	alert := js.Global().Get("alert")
	alert.Invoke("hello Golang") //等价于 window.alert("hello Golang")
}
