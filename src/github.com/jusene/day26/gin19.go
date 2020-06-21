package main

import "github.com/gin-gonic/gin"

type myForm struct {
	Colors []string `form:"colors[]"`
}

func main() {
	r := gin.Default()

	r.POST("/", func(context *gin.Context) {
		var fakeForm myForm
		context.ShouldBind(&fakeForm)
		context.JSON(200, gin.H{
			"color": fakeForm.Colors,
		})
	})
	r.Run()
}
