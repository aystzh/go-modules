package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func uploadFileTest() {
	//uploadFile()
	uploadFiles()
}

func uploadFile() {
	r := gin.Default()
	r.POST("/uploadSingle", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		context.String(http.StatusOK, "%s uploaded!", file.Filename)
	})
	r.Run()
}

//多文件上传
func uploadFiles() {
	r := gin.Default()
	r.POST("/uploadMany", func(context *gin.Context) {
		form, _ := context.MultipartForm()
		files := form.File["file"]
		for _, file := range files {
			log.Println(file.Filename)
		}
		context.String(http.StatusOK, "%d files uploaded!", len(files))
	})
	r.Run()
}
