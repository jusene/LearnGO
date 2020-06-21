package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	// r.LoadHTMLFiles("templates/template1.html")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
			"body":  "hello world",
		})
	})

	r.Run()
}
