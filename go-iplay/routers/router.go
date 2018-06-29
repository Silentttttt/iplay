// @APIVersion 1.0.0
// @Title iplay API
// @Description api for iplay.
package routers

import (
	"iplay/go-iplay/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	// beego.Router("/login", &controllers.UserController{}, "*:Login")
	// beego.Router("/sendsms", &controllers.UserController{}, "*:SendMobileCode")
	// beego.Router("/reg", &controllers.UserController{}, "*:Register")
	// beego.Router("/authentication", &controllers.UserController{}, "*:Authentication")

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.UserController{},
					&controllers.UserQuizzesController{},
				),
			),
			beego.NSNamespace("/game",
				beego.NSInclude(
					&controllers.GameController{},
					&controllers.QuizzesController{},
				),
			),
		)

	beego.AddNamespace(ns)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}
