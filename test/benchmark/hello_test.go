package main

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Println("hello")
	}
}

//并发测试 基准测试  压测
func BenchmarkParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			// 所有 goroutine 一起，循环一共执行 b.N 次
			buf.Reset()
			err := templ.Execute(&buf, "Word")
			if err != nil {
				fmt.Printf("err")
			}
		}
	})
}
