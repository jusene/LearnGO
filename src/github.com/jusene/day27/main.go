package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jusene/day27/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample api server
// @termsOfService http://swagger.io/terms

// @contact.name API Support
// @contact.url http://www.swagger.io/terms
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	v1 := r.Group("/v1")
	{
		v1.GET("/get/:name", controllers.GetPathParam)
		v1.GET("/get", controllers.Get)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run()
}
