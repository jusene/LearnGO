package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	/*
		viper.SetConfigType("json")
		viper.SetConfigName("server")
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				panic(fmt.Errorf("Fatal error config file: %s \n", err))
			}
		}
		fmt.Println(viper.GetString("host.address"))
		fmt.Println(viper.Get("datastore.metric"))
	*/
	viper.SetConfigName("server.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
	fmt.Println(viper.GetString("server.datastore.metric.address"))
}
