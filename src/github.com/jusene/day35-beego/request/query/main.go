package main

import "github.com/beego/beego/v2/server/web"

func main() {
	ctrl := &MainController{}
	web.Router("/", ctrl)
	web.Run()
}

type MainController struct {
	web.Controller
}

func (ctrl *MainController) Get() {
	name := ctrl.GetString("name")

	if name == "" {
		ctrl.Ctx.WriteString("Hello World")
		return
	}

	ctrl.Ctx.WriteString("hello " + name)
}
