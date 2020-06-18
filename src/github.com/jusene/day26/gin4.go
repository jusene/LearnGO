package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	// curl -d "message=you are work" http://127.0.0.1:8080/form  {"message":"you are work","nick":"jusene","status":"post"}
	router.POST("/form", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultPostForm("nick", "jusene")

		context.JSON(200, gin.H{
			"status":  "post",
			"message": message,
			"nick":    nick,
		})
	})

	router.Run()
}
