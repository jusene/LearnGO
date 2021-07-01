package main

import (
	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/captcha"
)

// init captcha
var cpt *captcha.Captcha

// Controller
type Controller struct {
	web.Controller
}

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 100
	cpt.StdHeight = 40
}

func main() {
	ctrl := &Controller{}

	web.Router("/", ctrl)
	web.Router("/sendCaptcha", ctrl, "post:Captcha")
	web.Run()
}

// Get
func (ctrl *Controller) Get() {
	ctrl.TplName = "captcha.html"
	ctrl.Data["name"] = "Home"
	_ = ctrl.Render()
}

// Captcha
func (ctrl *Controller) Captcha() {
	ctrl.TplName = "captcha.html"

	if !cpt.VerifyReq(ctrl.Ctx.Request) {
		logs.Error("Captcha does not match")
		_ = ctrl.Render()
		return
	}

	logs.Info("matched")
	_ = ctrl.Render()
}
