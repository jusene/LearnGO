package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	// set a lower memory limit for multipart forms(default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8MiB
	// curl -X POST http://127.0.0.1:8080/upload -F "upload[]=@gin1.go" -H "Content-Typ
	//e: multipart/form-data"
	router.POST("/upload", func(context *gin.Context) {
		// Multipart form
		form, _ := context.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			context.SaveUploadedFile(file, "./test.go")
		}
		context.String(http.StatusOK, fmt.Sprintf("%d files uploads!", len(files)))
	})

	router.Run()
}
