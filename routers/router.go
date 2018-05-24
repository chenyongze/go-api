package routers

import (
	"github.com/astaxie/beego"
	"github.com/chenyongze/go-api/controllers"
)

func init() {

	beego.Informational("init routers start ...")

	// 默认登录
	beego.Router("/", &controllers.ApiDocController{}, "*:Index")
	beego.Router("/login", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")
	beego.Router("/no_auth", &controllers.LoginController{}, "*:NoAuth")

	beego.Router("/home", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")

	beego.AutoRouter(&controllers.ApiController{})
	beego.AutoRouter(&controllers.ApiSourceController{})
	beego.AutoRouter(&controllers.ApiPublicController{})
	beego.AutoRouter(&controllers.TemplateController{})
	beego.AutoRouter(&controllers.ApiDocController{})
	//add lots
	beego.AutoRouter(&controllers.LotsController{})
	beego.AutoRouter(&controllers.DingController{})
	beego.AutoRouter(&controllers.WeChatController{})
	beego.AutoRouter(&controllers.AuctionController{})
	beego.AutoRouter(&controllers.TestController{})

	// beego.AutoRouter(&controllers.ApiMonitorController{})
	beego.AutoRouter(&controllers.EnvController{})
	beego.AutoRouter(&controllers.CodeController{})

	beego.AutoRouter(&controllers.GroupController{})
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.UserController{})

	beego.Informational("init routers end.")

}
