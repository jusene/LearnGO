package main

import (
	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Router("/abort", &AbortController{})
	web.Run()
}

type AbortController struct {
	web.Controller
}

func (a *AbortController) Get() {
	a.Abort(strconv.Itoa(http.StatusUnauthorized))

	user := a.GetString("user")
	if user != "" {
		a.Ctx.WriteString("hello " + user)
		return
	} else {
		a.Ctx.WriteString("hello world")
		return
	}

}
