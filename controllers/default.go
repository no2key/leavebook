package controllers

import (
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/validation"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"

    "html/template"
    //"fmt"
)

func init() {
    orm.RegisterModel(new(message))

    orm.RegisterDriver("mysql", orm.DR_MySQL)
    
    //orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
    orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser")+ ":" + beego.AppConfig.String("mysqlpass") + "@tcp(" + beego.AppConfig.String("mysqlurls") +  ":3306)/" + beego.AppConfig.String("mysqldb") +  "?charset=utf8")
}

//字段小写或form为-，则忽略
type message struct {
    Id int              `form:"-"`
    Nick_name string    `form:"name"`
    Sex int             `form:"sex"`
    Email string        `form:"email"`
    Msg string          `form:"message"`
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
    //c.Abort("404") 中断抛出异常
    beego.Warn("this is warn")

    flash := beego.ReadFromRequest(&c.Controller)
    if n,ok:=flash.Data["notice"];ok{
        //显示设置成功
        c.Data["msg"] = n
    }else if n,ok=flash.Data["error"];ok{
        //显示错误
         c.Data["msg"] = n
    }

	c.Data["xsrf"] = template.HTML(c.XsrfFormHtml())
	c.TplNames = "index.html"
    c.Render()
}

func (c *MainController) Post() {
    msg := message{}

    if err := c.ParseForm(&msg); err != nil {
       c.Ctx.WriteString("解析错误")
    }

    flash:=beego.NewFlash()
    valid := validation.Validation{}

    //if msg.Email == "" {
    //valid.Required(msg.Email, "string")
    //if valid.HasErrors() {
    if !valid.Required(msg.Email, "string").Ok {
        flash.Error("邮箱不能为空!")
        flash.Store(&c.Controller)
        c.Redirect("/",302)
        return
    }
    
    
    if  !valid.Email(msg.Email , "string").Ok {
        flash.Error("邮箱格式错误!")
        flash.Store(&c.Controller)
        c.Redirect("/",302)
        return
    }

    if !valid.Required(msg.Nick_name, "string").Ok {
        flash.Error("昵称不能为空!")
        flash.Store(&c.Controller)
        c.Redirect("/",302)
        return
    }

    if !valid.MaxSize(msg.Nick_name, 10, "string").Ok {
        flash.Error("昵称长度不能超过10!")
        flash.Store(&c.Controller)
        c.Redirect("/",302)
        return
    }

    if !valid.Required(msg.Msg, "string").Ok {
        flash.Error("信息不能为空!")
        flash.Store(&c.Controller)
        c.Redirect("/",302)
        return
    }

    if !valid.MaxSize(msg.Msg, 100, "string").Ok {
        flash.Error("信息长度不能超过100!")
        flash.Store(&c.Controller)
        c.Redirect("/",302)
        return
    }
    
    if !valid.Range(msg.Sex, 0, 1, "int").Ok {
        flash.Error("性别选择错误!")
        flash.Store(&c.Controller)
        c.Redirect("/",302)
        return
    }

    //fmt.Println("msg", msg)
    //c.Ctx.WriteString("test:" + msg.Nick_name)

    o := orm.NewOrm()
    id, err := o.Insert(&msg)
    if err == nil {
        beego.Warn(id)
    }
    
    flash.Notice("提交成功!")
    flash.Store(&c.Controller)
    c.Redirect("/",302)
}
