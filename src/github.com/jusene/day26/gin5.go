package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	// curl -d "name=jusene" http://127.0.0.1:8080/post?id=1 {"id":"1","message":"ok","name":"jusene","page":"0"}
	router.POST("/post", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "0")
		name := context.PostForm("name")
		message := context.DefaultPostForm("message", "ok")

		context.JSON(200, gin.H{
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})
	})

	router.Run()
}
