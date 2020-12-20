package main

import (
	"github.com/beego/beego/v2/server/web"
	"html/template"
)

func main() {
	web.BConfig.WebConfig.EnableXSRF = true
	web.BConfig.WebConfig.XSRFExpire = 3600
	web.BConfig.WebConfig.XSRFKey = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"

	mc := &MainController{}
	web.Router("/xsrfpage", mc, "get:XsrfPage")
	web.Router("/new_message", mc, "post:NewMessage")

	web.Run(":8080")

}

type MainController struct {
	web.Controller
}

func (mc *MainController) XsrfPage() {
	mc.XSRFExpire = 7200
	mc.Data["xsrfdata"] = template.HTML(mc.XSRFFormHTML())
	mc.TplName = "xsrf.html"
}

func (mc *MainController) NewMessage() {
	v, _ := mc.Input()
	mc.Ctx.WriteString("hello" + v.Get("message"))
}
