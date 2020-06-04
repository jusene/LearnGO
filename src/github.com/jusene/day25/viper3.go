package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.Set("KEY", "golang")
	viper.AddConfigPath(".")
	viper.SetConfigName("test")
	viper.SetConfigType("yaml")

	// 文件不存在，会报错  会覆盖
	if err := viper.WriteConfig(); err != nil {
		fmt.Printf("%v\n", err)
	}

	// 文件不存在，会创建 不会覆盖
	if err := viper.SafeWriteConfig(); err != nil {
		fmt.Printf("%v\n", err)
	}

	// 文件不存在，会创建   会覆盖
	if err := viper.WriteConfigAs("./.config.yaml"); err != nil {
		fmt.Printf("%v\n", err)
	}

	// 文件不存在，会创建  不会覆盖
	if err := viper.SafeWriteConfigAs("./.config.yaml"); err != nil {
		fmt.Printf("%v\n", err)
	}
}
