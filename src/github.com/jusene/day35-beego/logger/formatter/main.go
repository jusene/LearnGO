package main

import (
	"github.com/beego/beego/v2/core/logs"
)

func main() {

	f := &logs.PatternLogFormatter{
		Pattern:    "%F:%n|%w%t>> %m",
		WhenFormat: "2006-01-02",
	}
	logs.RegisterFormatter("pattern", f)

	_ = logs.SetLogger("console", `{"formatter": "pattern"}`)

	logs.Info("hello, world")
}
