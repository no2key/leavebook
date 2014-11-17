package routers

import (
	"github.com/wskm/leavebook/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //一般路由
    //-----------------------------------------------------------
    beego.Router("/", &controllers.MainController{})

    beego.Router("/testc", &controllers.TestCtController{})

    //正则路径测试
    beego.Router("/routerreg/:path", &controllers.RouterRegexController{})

    //restfull风格路径设置，设置的函数名首字母要大写
    beego.Router("/restfull/list/:id:int",&controllers.RestfullController{},"*:List")   //http://127.0.0.1:8080/restfull/list/343243232
    beego.Router("/restfull/simple/*",&controllers.RestfullController{},"get:GetTest;post:PostTest")
    //beego.Router("/simple",&SimpleController{},"*:AllFunc;post:PostFunc")   //那么执行 POST 请求的时候，执行 PostFunc 而不执行 AllFunc。

    //自动路由  http://127.0.0.1:8080/autorouter/list/xxxxxx
    //-----------------------------------------------------------
    beego.AutoRouter(&controllers.AutoRouterController{})

    //注解路由
    //-----------------------------------------------------------
    beego.Include(&controllers.CommentRouterController{})
    //http://127.0.0.1:8080/all/ddd  or http://127.0.0.1:8080/staticblock/ddd

    //命名空间请看官方文档
   
}
