package controllers

import (
	"github.com/astaxie/beego"
)

type RestfullController struct {
	beego.Controller
}

func (c *RestfullController) List() {
    path := c.Ctx.Input.Param(":id")
    c.Ctx.Output.Body([]byte(path))
}


func (c *RestfullController) GetTest() {
    c.Ctx.Output.Body([]byte("get test"))
}

func (c *RestfullController) PostTest() {
    c.Ctx.Output.Body([]byte("post test"))
}