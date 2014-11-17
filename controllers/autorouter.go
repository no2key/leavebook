package controllers

import (
	"github.com/astaxie/beego"
)

type AutoRouterController struct {
	beego.Controller
}

func (c *AutoRouterController) List() {
    c.Ctx.Output.Body([]byte( c.Ctx.Input.Param(":ext") + "|" + c.Ctx.Input.Params["0"] ))
}
