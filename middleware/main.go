package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(RequestIdMiddleware())
	router.Use(DummyMiddleware)
	router.GET("/middleware", GetDummyEndpoint) //基准应用程序
	router.Run()
}

func GetDummyEndpoint(context *gin.Context) {
	res := map[string]string{"hello": "word!"}
	context.JSON(http.StatusOK, res)
}

//第一种方式
func DummyMiddleware(c *gin.Context) {
	fmt.Println("Im a dummy!!!")
	c.Next()
}

//第二种方式 自定义中间件  请求ID 中间件
func RequestIdMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//请求前
		t := time.Now()
		context.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
		context.Next()
		//请求后
		payTime := time.Since(t)
		fmt.Println("payTime:", payTime)
	}
}
