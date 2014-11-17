package controllers

import (
	"github.com/astaxie/beego"
)

type CommentRouterController struct {
	beego.Controller
}

func (c *CommentRouterController) URLMapping() {
    c.Mapping("StaticBlock", c.StaticBlock)
    c.Mapping("AllBlock", c.AllBlock)
}


// @router /staticblock/:key [get]
func (c *CommentRouterController) StaticBlock() {
    c.Ctx.Output.Body([]byte("StaticBlock test"))
}

// @router /all/:key [get]
func (c *CommentRouterController) AllBlock() {
    c.Ctx.Output.Body([]byte("AllBlock test"))
}