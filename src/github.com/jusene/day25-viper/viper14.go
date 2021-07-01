package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config.yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}

	// 监控配置文件变化
	viper.WatchConfig()

	r := gin.Default()
	r.GET("/version", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"CODE":    200,
			"VERSION": viper.GetString("version"),
		})
	})

	if err := r.Run(fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
		panic(err)
	}
}
