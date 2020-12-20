package main

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	ctrl := &MainController{}
	web.Router("/:id/:name", ctrl)
	web.Run()
}

type MainController struct {
	web.Controller
}

func (ctrl *MainController) Get() {
	n := ctrl.Ctx.Input.Param(":name")
	id := ctrl.Ctx.Input.Param(":id")
	//name := ctrl.Ctx.Input.Params()
	fmt.Println(n)
	ctrl.Ctx.WriteString("Your router id is " + id)
}
