package main

import (
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	c := &MainController{}
	web.BConfig.CopyRequestBody = true

	web.Router("/", c)
	web.Run()
}

type MainController struct {
	web.Controller
}

type user struct {
	Name     string                 `json:"name"`
	Password string                 `json:"password"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

func (c *MainController) Post() {
	input := user{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input); err != nil {
		c.Data["json"] = err.Error()
	}

	c.Data["json"] = input
	c.ServeJSON()
}
