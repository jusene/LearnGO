package main

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
	apmbeego "github.com/opentracing-contrib/beego"
)

func main() {
	web.Router("/", &MainController{})
	// Actually, I just use the opentracing-contrib/web as example but I do not check
	// whether it is a good middleware
	web.RunWithMiddleWares("localhost:8080", apmbeego.Middleware("bee-go-demo"))

	// start the server and then request GET http://localhost:8080/
}

type MainController struct {
	web.Controller
}

func (m *MainController) Get() {
	m.Ctx.WriteString(fmt.Sprintf("hello world"))
}
