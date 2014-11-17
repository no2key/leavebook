package controllers

import (
	"github.com/astaxie/beego"
    "strconv"
)

var html string

type TestCtController struct {
	beego.Controller
}

func (c *TestCtController) Prepare(){
    html = " Prepare test "
}

func (c *TestCtController) Finish() {
    c.Ctx.Output.Body([]byte(html + " Finish test"))

}


func (c *TestCtController) Get() {
	//c.Ctx.Output.Body([]byte("get test"))
    html += " get test "
    
    /*
    GetString(key string) string
    GetStrings(key string) []string
    GetInt(key string) (int64, error)
    GetBool(key string) (bool, error)
    GetFloat(key string) (float64, error)
    */

    jsoninfo := c.GetString("test")
    if jsoninfo == "" {
        c.Ctx.WriteString("test is empty")
        return
    }else{
        c.Ctx.WriteString("test is not empty")
    }

    var num int
    c.Ctx.Input.Bind(&num, "num")
    
    c.Ctx.WriteString(c.Input().Get("id")+strconv.Itoa(num))

}

func (c *TestCtController) Post() {
    c.Ctx.Output.Body([]byte(" post test "))
}
