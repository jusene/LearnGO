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

type user struct {
	Name     string `form:"name"`
	Password string `form:"password" json:"custom_password"`
}

func (ctrl *MainController) Post() {
	input := user{}

	if err := ctrl.ParseForm(&input); err != nil {
		ctrl.Ctx.WriteString(err.Error())
	}

	ctrl.Data["json"] = input
	ctrl.ServeJSON()
}
