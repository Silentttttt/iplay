// @APIVersion 1.0.0
// @Title iplay API
// @Description api for iplay.
package routers

import (
	"iplay/go-iplay/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	// beego.Router("/login", &controllers.UserController{}, "*:Login")
	// beego.Router("/sendsms", &controllers.UserController{}, "*:SendMobileCode")
	// beego.Router("/reg", &controllers.UserController{}, "*:Register")
	// beego.Router("/authentication", &controllers.UserController{}, "*:Authentication")

	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.UserController{},
				),
			),
			beego.NSNamespace("/game",
				beego.NSInclude(
					&controllers.GameController{},
				),
			),
		)

	beego.AddNamespace(ns)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}
