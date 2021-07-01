package main

import (
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/beego/beego/v2/core/logs/es"
)

func main() {
	err := logs.SetLogger(logs.AdapterEs, `{"dsn":"http://localhost:9200/","level":1}`)
	if err != nil {
		panic(err)
	}
	logs.Info("hello beego")
}
