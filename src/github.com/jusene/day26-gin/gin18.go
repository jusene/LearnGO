package main

import "github.com/gin-gonic/gin"

type Header struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		H := Header{}
		if err := context.ShouldBindHeader(&H); err != nil {
			context.JSON(200, err)
		}

		context.JSON(200, gin.H{"Rate": H.Rate, "Domain": H.Domain})
	})

	r.Run()
}
