package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// GROUP v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/get", func(context *gin.Context) {
			context.String(200, "GET")
		})
		v1.POST("/post", func(context *gin.Context) {
			context.String(200, "POST")
		})
	}
	router.Run()
}
