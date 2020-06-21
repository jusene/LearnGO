package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/redic", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	r.GET("/test", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/redic")
	})

	r.GET("/test1", func(context *gin.Context) {
		context.Request.URL.Path = "/test2"
		r.HandleContext(context)
	})

	r.GET("/test2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	r.Run()
}
