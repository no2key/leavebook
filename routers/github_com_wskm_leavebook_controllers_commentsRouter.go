package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["github.com/wskm/leavebook/controllers:CommentRouterController"] = append(beego.GlobalControllerRouter["github.com/wskm/leavebook/controllers:CommentRouterController"],
		beego.ControllerComments{
			"StaticBlock",
			`/staticblock/:key`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/wskm/leavebook/controllers:CommentRouterController"] = append(beego.GlobalControllerRouter["github.com/wskm/leavebook/controllers:CommentRouterController"],
		beego.ControllerComments{
			"AllBlock",
			`/all/:key`,
			[]string{"get"},
			nil})

}
