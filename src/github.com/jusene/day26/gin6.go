package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// curl -X POST -d "names[c]=jusne" http://127.0.0.1:8080/post?ids[a]=123&ids[b]=456  {"ids":{"a":"123","b":"456"},"name":{"c":"jusene"}}
	router.POST("/post", func(context *gin.Context) {
		ids := context.QueryMap("ids")
		name := context.PostFormMap("names")

		context.JSON(200, gin.H{
			"ids":  ids,
			"name": name,
		})
	})

	router.Run()
}
