package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// http://localhost:8080/
func routeNoParamsTest() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "无参数请求")
	})
	r.Run()
}

// http://localhost:8081/user/zhangsan
func routePathParamsTest() {
	r := gin.Default()
	r.GET("/user/:name", func(context *gin.Context) {
		param := context.Param("name")
		context.String(http.StatusOK, param)
	})
	r.Run(":8081")
}

// http://localhost:8082/users?name=zhangsan&role=doctor
func routeQueryParamsTest() {
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "teacher") //role 可选参数
		c.String(http.StatusOK, "%s is %s ", name, role)
	})
	r.Run(":8082")
}

// POST form data 格式
func routePostParamsTest() {
	r := gin.Default()
	r.POST("/form", func(context *gin.Context) {
		username := context.PostForm("username")
		role := context.DefaultPostForm("role", "teacher")
		context.JSON(http.StatusOK, gin.H{
			"username": username,
			"role":     role,
		})
	})
	r.Run(":8083")
}

func queryMixPostParams() {
	r := gin.Default()
	r.POST("/form", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "2")
		name := context.PostForm("name")
		role := context.DefaultPostForm("role", "doctor")
		context.JSON(http.StatusOK, gin.H{
			"id":   id,
			"page": page,
			"name": name,
			"role": role,
		})
	})
	r.Run(":8084")
}

func queryMapParams() {
	r := gin.Default()
	r.POST("/post", func(context *gin.Context) {
		ids := context.QueryMap("ids")
		names := context.PostFormMap("names")
		context.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})
	r.Run()
}

func redirectTest() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "重定向测试")
	})

	r.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/index")
	})
	r.GET("/index", func(context *gin.Context) {
		context.Request.URL.Path = "/"
		r.HandleContext(context)
	})
	r.Run(":9900")

}

func groupRouteTest() {
	r := gin.Default()
	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}
	//group v1
	v1 := r.Group("/v1")
	{
		v1.GET("/posts", defaultHandler)
		v1.GET("/series", defaultHandler)
	}
	//group v2
	v2 := r.Group("/v2")
	{
		v2.GET("/posts", defaultHandler)
		v2.GET("/series", defaultHandler)
	}
	r.Run(":9998")
}

func routeTest() {
	//routeNoParamsTest()
	//routePathParamsTest()
	//routeQueryParamsTest()
	//routePostParamsTest()
	//queryMixPostParams()
	//queryMapParams() 不常用
	//redirectTest()
	groupRouteTest() //分组路由
}
