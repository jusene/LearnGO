package main

import (
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
)

func main() {
	req := httplib.Post("http://beego.me/")
	req.Header("Accept-Encoding", "gzip,deflate,sdch")
	req.Header("Host", "beego.me")

	header := req.GetRequest().Header
	logs.Info(header)
}
