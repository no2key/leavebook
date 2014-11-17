package controllers

import (
	"github.com/astaxie/beego"
)

type RouterRegexController struct {
	beego.Controller
}

func (c *RouterRegexController) Get() {
	path := c.Ctx.Input.Param(":path")
    c.Ctx.Output.Body([]byte(path))
}
