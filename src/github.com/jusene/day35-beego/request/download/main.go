package main

import "github.com/beego/beego/v2/server/web"

func main() {
	ctrl := &MainController{}
	web.Router("/download", ctrl, "get:Download")
	web.Router("/download_file", ctrl, "get:DownloadFile")
	web.Run()
}

type MainController struct {
	web.Controller
}

func (ctrl *MainController) DownloadFile() {
	ctrl.Ctx.Output.Download("LICENSE", "license.txt")
}

// Download is an example that download the content stored in memory
// when you access localhost:8080/download
// A file named "mytest.txt" with content "Hello" will be downloaded
func (ctrl *MainController) Download() {
	output := ctrl.Ctx.Output
	output.Header("Content-Disposition", "attachment;filename=mytest.txt;")
	output.Header("Content-Description", "File Transfer")
	output.Header("Content-Type", "application/octet-stream")
	output.Header("Content-Transfer-Encoding", "binary")
	output.Header("Expires", "0")
	output.Header("Cache-Control", "must-revalidate")
	output.Header("Pragma", "public")
	ctrl.Ctx.WriteString("Hello")
}
