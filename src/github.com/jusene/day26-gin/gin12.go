package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()

	router.Use(Cors)
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.Run()
}

func Cors(ctx *gin.Context) {
	ctx.Header("BY-X-SERVER", "GIN")
	ctx.Next()
}
