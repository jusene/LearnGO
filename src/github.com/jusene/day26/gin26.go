package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/mod", "./")
	r.StaticFS("/more", http.Dir("./"))
	r.StaticFile("/favicon.ico", "./favicon.ico")

	r.Run()
}
