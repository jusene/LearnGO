package main

import "github.com/beego/beego/v2/server/web"

func main() {
	ctrl := &MainController{}
	web.Router("/cookie", ctrl, "post:PutCookie")
	web.Router("/cookie", ctrl, "get:ReadCookie")
	web.Run()
}

type MainController struct {
	web.Controller
}

func (ctrl *MainController) PutCookie() {
	ctrl.Ctx.SetCookie("name", "web-cookie", 10)
	ctrl.TplName = "hello_world.html"
	ctrl.Data["name"] = "PutCookie"
	_ = ctrl.Render()
}

func (ctrl *MainController) ReadCookie() {
	ctrl.TplName = "hello_world.html"
	ctrl.Data["name"] = ctrl.Ctx.GetCookie("name")
	_ = ctrl.Render()
}
