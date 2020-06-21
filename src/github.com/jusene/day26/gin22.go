package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/someJSON", func(context *gin.Context) {
		names := []string{"name", "ga", "age"}

		context.SecureJSON(http.StatusOK, names)
	})
	// while(1);["name","ga","age"]
	r.Run()
}
