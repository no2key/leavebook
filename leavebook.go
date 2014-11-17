package main

import (
	_ "github.com/wskm/leavebook/routers"
	"github.com/astaxie/beego"
)

/*
func page_not_found(rw http.ResponseWriter, r *http.Request){
    t,_:= template.New("404.html").ParseFiles(beego.ViewsPath+"/404.html")
    data :=make(map[string]interface{})
    data["content"] = "page not found"
    t.Execute(rw, data)
}
*/

func main() {
    //默认日志输出到console,可以设置写到文件
    //beego.SetLogger("file", `{"filename":"logs/test.log"}`)
    //beego.BeeLogger.DelLogger("console")  取消输出到命令行，否则文件、命令都有输出。

    //自定义错误页面，如:dbError
    //beego.Errorhandler("404",page_not_found)

    //Bool,int,int64,Float,String
    //beego.AppConfig.String("mysqluser")
    //beego.AppConfig.String("mysqlpass")
    //beego.AppConfig.String("mysqlurls")
    //beego.AppConfig.String("mysqldb")

    //beego.GetConfig(typ, key string)  获取自定义参数

	beego.Run()

}

