package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("Tag", map[string]string{"tag": "tags", "info": "infos"})

	fmt.Println(viper.Get("ContentDir"))
	fmt.Println(viper.Get("Tag"))
}
