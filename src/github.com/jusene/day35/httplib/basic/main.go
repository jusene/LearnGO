package main

import (
	"crypto/tls"
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
)

func main() {
	// Get
	req := httplib.Get("http://www.baidu.com")
	str, err := req.String()
	if err != nil {
		logs.Error(err)
	}

	logs.Info(str)

	// Post
	req = httplib.Post("http://www.baidu.com")
	req.Param("password", "aaaa")
	resp, err := req.DoRequest()
	logs.Info(req.GetRequest().URL)
	if err != nil {
		logs.Error(err)
	}
	logs.Info(resp.Request)

	// https
	req = httplib.Get("https://www.baidu.com")
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err = req.DoRequest()
	if err != nil {
		logs.Error(err)
	}
	logs.Info(resp.Request)
}
