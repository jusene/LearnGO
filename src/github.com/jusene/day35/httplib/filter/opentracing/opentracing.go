package opentracing

import (
	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/client/httplib/filter/opentracing"
	"github.com/beego/beego/v2/core/logs"
)

func main() {

	// don't forget this to inject the opentracing API's implementation
	// opentracing2.SetGlobalTracer()

	builder := opentracing.FilterChainBuilder{}
	req := httplib.Get("http://beego.me/")
	// only work for this request, or using SetDefaultSetting to support all requests
	req.AddFilters(builder.FilterChain)

	resp, err := req.Response()
	if err != nil {
		logs.Error("could not get response: ", err)
	} else {
		logs.Info(resp)
	}
}
