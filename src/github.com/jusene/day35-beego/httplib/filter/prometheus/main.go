package main

import (
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/client/httplib/filter/prometheus"
	"github.com/beego/beego/v2/core/logs"
)

func main() {
	builder := prometheus.FilterChainBuilder{
		AppName:    "My-test",
		ServerName: "User-server-1",
		RunMode:    "dev",
	}
	req := httplib.Get("http://www.baidu.com/")
	// only work for this request, or using SetDefaultSetting to support all requests
	req.AddFilters(builder.FilterChain)

	resp, err := req.Response()
	if err != nil {
		logs.Error("could not get response: ", err)
	} else {
		logs.Info(resp)
	}
}
