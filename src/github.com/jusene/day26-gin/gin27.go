package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/local/file", func(context *gin.Context) {
		context.File("./gin.log")
	})

	var fs http.FileSystem
	router.GET("/fs/file", func(context *gin.Context) {
		context.FileFromFS("gin1.go", fs)
	})

	router.Run()
}
