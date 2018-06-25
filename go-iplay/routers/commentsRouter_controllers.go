package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["iplay/go-iplay/controllers:UserController"] = append(beego.GlobalControllerRouter["iplay/go-iplay/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["iplay/go-iplay/controllers:UserController"] = append(beego.GlobalControllerRouter["iplay/go-iplay/controllers:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/reg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
