package main

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"html/template"
)

func main() {
	web.BConfig.WebConfig.EnableXSRF = true
	web.BConfig.WebConfig.XSRFExpire = 3
	web.BConfig.WebConfig.XSRFKey = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"

	mc := &MainController{}
	web.Router("/xsrfpage", mc, "get:XsrfPage")
	web.Router("/new_message", mc, "*:NewMessage")

	web.Run(":8080")

}

type MainController struct {
	web.Controller
}

func (mc *MainController) XsrfPage() {
	mc.XSRFExpire = 5
	mc.Data["xsrfdata"] = template.HTML(mc.XSRFFormHTML())
	fmt.Println(mc.XSRFToken())
	mc.TplName = "xsrf.html"
}

func (mc *MainController) NewMessage() {
	v, _ := mc.Input()
	if mc.XSRFToken() == v.Get("_xsrf") {
		mc.Ctx.WriteString("hello" + v.Get("_xsrf"))
	} else {
		mc.Redirect("/xsrfpage", 302)
	}

}
