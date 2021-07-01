package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type config struct {
	Port    int
	Name    string
	PathMap string `mapstructure:"path_map"`
}

func main() {
	var C config
	err := viper.Unmarshal(&C)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}
