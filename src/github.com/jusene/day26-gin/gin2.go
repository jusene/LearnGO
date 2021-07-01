package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func gettinng(ctx *gin.Context) {
	ctx.String(http.StatusOK, "GET")
}

func main() {
	router := gin.Default()
	// http://127.0.0.1:8080/get GET
	router.GET("/get", gettinng)
	// http://127.0.0.1:8080/get/jusene HELLO, jusene
	router.GET("/get/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "HELLO, %s", name)
	})
	// http://127.0.0.1:8080/get/jusene/29 {"age":"29","code":200,"name":"jusene"}
	router.GET("/get/:name/:age", func(context *gin.Context) {
		name := context.Param("name")
		age := context.Param("age")
		context.JSON(200, gin.H{
			"name": name,
			"age":  age,
			"code": 200,
			"path": context.FullPath(),
		})
	})
	// http://127.0.0.1:8080/get/jusene/29/send   jusene is /send
	router.GET("/get/:name/:age/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		message := name + " is " + action
		context.String(http.StatusOK, message)
	})
	// http://127.0.0.1:8080/post/jusene/29/send /post/:name/:age/*action
	router.POST("/post/:name/:age/*action", func(context *gin.Context) {
		context.String(http.StatusOK, context.FullPath())
	})
	router.Run()
}
