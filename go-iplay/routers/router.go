package routers

import (
	"star/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.UserController{}, "*:Login")
	beego.Router("/sendsms", &controllers.UserController{}, "*:SendMobileCode")
	beego.Router("/reg", &controllers.UserController{}, "*:Register")
	beego.Router("/authentication", &controllers.UserController{}, "*:Authentication")
}
