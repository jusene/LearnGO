package main

import (
	"github.com/beego/beego/v2/server/web"
)

func main() {
	ctrl := &MainController{}
	web.Router("/", ctrl)
	web.Router("/health", ctrl, "get:Health")
	web.Router("/update", ctrl, "get,post:GetOrPost")
	web.Router("/any", ctrl, "*:Any")
	web.Run()
}

type MainController struct {
	web.Controller
}

func (ctrl *MainController) Health() {
	ctrl.Ctx.Output.Body([]byte("hello world"))
}

func (ctrl *MainController) GetOrPost() {
	ctrl.Ctx.WriteString("HELLO WORLD")
}

func (ctrl *MainController) Any() {
	var msg = map[string]string{
		"method": ctrl.Ctx.Request.Method,
	}
	ctrl.Data["json"] = msg
	ctrl.ServeJSON()
}
