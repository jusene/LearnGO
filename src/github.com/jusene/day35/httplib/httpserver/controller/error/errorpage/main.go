package main

import (
	"github.com/beego/beego/v2/server/web"
	"net/http"
)

func serverError(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("error 500"))
}

func main() {
	web.ErrorHandler("500", serverError)
	web.Router("/", &MainController{})
	web.Run()
}

type MainController struct {
	web.Controller
}

func (m *MainController) Get() {
	m.Abort("500")
}
