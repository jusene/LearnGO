package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {
	viper.SetEnvPrefix("spf")
	viper.BindEnv("id")
	viper.BindEnv("name", "name")

	os.Setenv("SPF_ID", "12")
	os.Setenv("NAME", "JUSENE")
	fmt.Println(viper.Get("id"))
	fmt.Println(viper.Get("name"))
}
