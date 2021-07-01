package main

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"net/http"
)

func main() {
	web.Router("/", &MainController{})
	web.RunWithMiddleWares(":8080", customMiddleWare1, func(handler http.Handler) http.Handler {
		return &customMiddleWare2{
			next: handler,
		}
	})
}

type MainController struct {
	web.Controller
}

func customMiddleWare1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logs.Info("log: in ths middleware1 from: ", r.RequestURI)
		w.Header().Add("MiddleWare", "1")
		next.ServeHTTP(w, r)
	})
}

type customMiddleWare2 struct {
	next http.Handler
}

func (c *customMiddleWare2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logs.Info("log: in ths middleware2 from: ", r.RequestURI)
	w.Header().Add("MiddleWare", "2")
	c.next.ServeHTTP(w, r)
}

func (m *MainController) Get() {
	m.Ctx.WriteString("hello world")
}
