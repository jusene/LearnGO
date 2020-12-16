package main

import beego "github.com/beego/beego/v2/server/web"

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("Hello World")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
