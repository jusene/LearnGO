package main

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type MainController struct {
	web.Controller
}

// http://127.0.0.1:8080/?id=21212&isok=true&ol[0]=1&ol[1]=2&ul[]=all&ul[]=test&user.Name=jusene
func (ctrl *MainController) Get() {
	var id int
	_ = ctrl.Ctx.Input.Bind(&id, "id")
	fmt.Println(id)

	var isok bool
	_ = ctrl.Ctx.Input.Bind(&isok, "isok")
	fmt.Println(isok)

	ol := make([]int, 0, 2)
	_ = ctrl.Ctx.Input.Bind(&ol, "ol")
	fmt.Println(ol)

	ul := make([]string, 0, 2)
	_ = ctrl.Ctx.Input.Bind(&ul, "ul")
	fmt.Println(ul)

	user := struct {
		Name string
	}{}
	_ = ctrl.Ctx.Input.Bind(&user, "user")
	fmt.Println(user)

	ctrl.Ctx.WriteString(strconv.Itoa(id))
}

func main() {
	ctrl := &MainController{}
	web.Router("/", ctrl)
	web.Run()
}
