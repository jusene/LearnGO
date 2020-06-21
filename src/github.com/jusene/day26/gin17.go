package main

import "github.com/gin-gonic/gin"

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	route := gin.Default()
	route.GET("/:name/:id", func(context *gin.Context) {
		var person Person
		if err := context.ShouldBindUri(&person); err != nil {
			context.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		context.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
	route.Run()
}
