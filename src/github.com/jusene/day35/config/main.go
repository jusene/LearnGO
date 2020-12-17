package main

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

//var ConfigFile = "./conf/app.conf"

func main() {
	//	err := web.LoadAppConfig("ini", ConfigFile)
	//	if err != nil {
	//		logs.Critical("ERROR")
	//		panic(err)
	//	}

	val, err := web.AppConfig.String("name")
	if err != nil {
		panic(err)
	}

	logs.Info("load config name is", val)
}
