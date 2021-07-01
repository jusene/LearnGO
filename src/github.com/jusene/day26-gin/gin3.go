package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// http://127.0.0.1:8080/welcome?firstname=zhang&lastname=jusene  hello zhang jusene
	router.GET("/welcome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "Guest") // 默认query param
		lastname := context.Query("lastname")

		context.String(http.StatusOK, "hello %s %s", firstname, lastname)
	})

	router.Run()
}
