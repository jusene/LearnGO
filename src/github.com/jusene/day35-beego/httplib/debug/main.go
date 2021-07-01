package main

import (
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

func main() {
	resp, err := httplib.Get("http://beego.me").SetTimeout(3*time.Second, 5*time.Second).Debug(true).DoRequest()
	if err != nil {
		logs.Error(err)
	}
	logs.Info(resp)
}
