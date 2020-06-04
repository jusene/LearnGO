package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./config.yaml")
	viper.WatchConfig()
	for {
		viper.OnConfigChange(func(in fsnotify.Event) {
			// 配置文件发生变化之后调用回调函数
			fmt.Println("Config file changed: ", in.Name, in.Op, in.String())
		})
	}
}
