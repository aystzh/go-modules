package main

import "github.com/gin-gonic/gin"

func ginTest() {
	res := gin.Default() //使用了gin.Default()生成了一个实例，这个实例即 WSGI 应用程序。
	res.GET("/", func(c *gin.Context) {
		c.String(200, "Hello,Gin!!!")
	}) //声明一个路由
	res.Run(":3333") //监听9999端口并启动本地服务 注意前面必须加 :
}
