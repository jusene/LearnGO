package main

import (
	"context"
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
	"net/http"
	"time"
)

func main() {
	httplib.SetDefaultSetting(httplib.BeegoHTTPSettings{
		FilterChains: []httplib.FilterChain{
			myFilter,
		},
		UserAgent:        "beegoServer",
		ConnectTimeout:   60 * time.Second,
		ReadWriteTimeout: 60 * time.Second,
		Gzip:             true,
		DumpBody:         true,
	})

	resp, err := httplib.Get("http://www.baidu.com").Response()
	if err != nil {
		logs.Error("could not get response: ", err)
	} else {
		logs.Info(resp)
	}

}

func myFilter(next httplib.Filter) httplib.Filter {
	return func(ctx context.Context, req *httplib.BeegoHTTPRequest) (response *http.Response, e error) {
		r := req.GetRequest()
		logs.Info(r.URL)

		return next(ctx, req)
	}
}
