package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
