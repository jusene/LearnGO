package main

import (
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"reflect"
)

func main() {
	err := viper.AddRemoteProvider("consul", "192.168.66.100:8500", "/oneci/config/arch")
	if err != nil {
		fmt.Println(err)
	}
	viper.SetConfigType("yaml")
	err = viper.ReadRemoteConfig()
	if err != nil {
		fmt.Println(err)
	}

	apps := viper.AllSettings()["apps"]
	fmt.Println(reflect.ValueOf(apps))
	for _, in := range apps.([]interface{}) {
		fmt.Println(in.(map[interface{}]interface{})["name"])
	}
}
