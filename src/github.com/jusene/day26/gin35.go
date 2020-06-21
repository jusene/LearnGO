package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/long_async", func(context *gin.Context) {
		cCp := context.Copy()
		go func() {
			time.Sleep(5 * time.Second)

			log.Println("Done! " + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(context *gin.Context) {
		time.Sleep(5 * time.Second)

		log.Println("Done! " + context.Request.URL.Path)
	})

	r.Run()
}
