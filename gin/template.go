package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func templateTest() {
	htmlTemplate()
}

type Student struct {
	Name string //坑： 首字母必须大写 否则渲染的时候会报错
	Age  int8
}

func htmlTemplate() {
	r := gin.Default()
	r.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "src/github.com/aystzh/go-modules/gin/templates/*")) //这里使用绝对路径

	stu1 := &Student{Name: "zhangsan", Age: 22}
	stu2 := &Student{Name: "lisi", Age: 33}
	r.GET("/arr", func(context *gin.Context) {
		context.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title":  "Gin",
			"stuArr": [2]*Student{stu1, stu2},
		})
	})
	r.Run()
}
