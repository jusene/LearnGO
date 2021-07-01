package main

import (
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
	"io/ioutil"
)

func main() {
	fileReq := httplib.Post("http://beego.me")
	fileReq.Param("username", "astaxie")
	fileReq.Param("password", "123456")
	fileReq.PostFile("uploadfile", "./hello.txt")
	resp, err := fileReq.DoRequest()
	logs.Info(resp.Status)

	// Bigfile
	bigFileReq := httplib.Post("http://beego.me/")
	bt, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		logs.Error(err)
	}
	bigFileReq.Body(bt)

}
