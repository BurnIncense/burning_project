package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建默认的 gin 路由器，并且已经附加了 Logger 和 Recovery 中间件
	router := gin.Default()

	// 创建 API 路由组
	api := router.Group("/api")
	{
		// 将 /hello GET 路由添加到路由器并定义路由处理程序
		api.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "world"})
		})
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	// 开始监听服务请求
	router.Run(":8080")
}
