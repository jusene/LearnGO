package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")       // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")         // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("/etc/appname") // 查找配置文件所在的路径
	viper.AddConfigPath(".")            // 多次调用以添加多个搜索路径
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	} // 查找并读取配置文件
	fmt.Println(viper.Get("apiVersion"))
}
