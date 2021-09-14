package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jusene/day27/controllers"
	_ "github.com/jusene/day27/docs"
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
// @schemes http https
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
func main() {
	gin.SetMode("release")
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(Cors())
	//url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	v1 := r.Group("/v1")
	{
		v1.GET("/get/:name", controllers.GetPathParam)
		v1.GET("/get", controllers.Get)
		v1.GET("/download", controllers.Download)

		v1.POST("/post/:name", controllers.PostPathParam)
		v1.POST("/post", controllers.Post)
		v1.POST("/postheader", controllers.PostHeader)
		v1.POST("/upload", controllers.Upload)

		v1.DELETE("/delete/:name", controllers.Delete)
		v1.DELETE("/deletequery/:id", controllers.DeleteQuery)

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Server", "GIN")
		context.Next()
	}
}

