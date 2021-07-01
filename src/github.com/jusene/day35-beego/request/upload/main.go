package main

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"os"
)

func main() {
	ctrl := &MainController{}

	web.Router("/upload", ctrl, "post:Upload")
	web.Router("/upload", ctrl, "get:UploadPage")
	web.Router("/save", ctrl, "post:Save")
	web.Router("/down/:filename", ctrl, "get:Down")
	web.Run()
}

type MainController struct {
	web.Controller
}

func (ctrl *MainController) UploadPage() {
	ctrl.TplName = "upload.html"
}

func (ctrl *MainController) Upload() {
	file, fileHeader, err := ctrl.GetFile("upload.txt")
	if err != nil {
		logs.Error("save file failed, ", err)
		ctrl.Ctx.Output.Body([]byte(err.Error()))
	} else {
		defer file.Close()

		logs.Info(fileHeader.Filename)
		ctrl.Ctx.Output.Body([]byte("success"))
	}
}

func (ctrl *MainController) Save() {
	err := ctrl.SaveToFile("save.txt", "./upload.txt")
	if err != nil {
		logs.Error("save file failed, ", err.Error())
		ctrl.Ctx.Output.Body([]byte(err.Error()))
	} else {
		ctrl.Ctx.Output.Body([]byte("success"))
	}
}

func (ctrl *MainController) Down() {
	filename := ctrl.Ctx.Input.Param(":filename")
	if _, err := os.Stat("./" + filename); err != nil {
		ctrl.Ctx.WriteString(err.Error())
	} else {
		ctrl.Ctx.Output.Download("./"+filename, "test.png")
	}
}
