package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["iplay/go-iplay/controllers:GameController"] = append(beego.GlobalControllerRouter["iplay/go-iplay/controllers:GameController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["iplay/go-iplay/controllers:QuizzesController"] = append(beego.GlobalControllerRouter["iplay/go-iplay/controllers:QuizzesController"],
		beego.ControllerComments{
			Method: "CreateQuizzes",
			Router: `/create_quizzes`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["iplay/go-iplay/controllers:QuizzesController"] = append(beego.GlobalControllerRouter["iplay/go-iplay/controllers:QuizzesController"],
		beego.ControllerComments{
			Method: "QuizzesList",
			Router: `/quizzes`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

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

	beego.GlobalControllerRouter["iplay/go-iplay/controllers:UserQuizzesController"] = append(beego.GlobalControllerRouter["iplay/go-iplay/controllers:UserQuizzesController"],
		beego.ControllerComments{
			Method: "DoQuizzes",
			Router: `/do_quizzes`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["iplay/go-iplay/controllers:UserQuizzesController"] = append(beego.GlobalControllerRouter["iplay/go-iplay/controllers:UserQuizzesController"],
		beego.ControllerComments{
			Method: "UserQuizzesList",
			Router: `/quizzes_list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["iplay/go-iplay/controllers:UserQuizzesController"] = append(beego.GlobalControllerRouter["iplay/go-iplay/controllers:UserQuizzesController"],
		beego.ControllerComments{
			Method: "EndQuizzes",
			Router: `/ui_h*9yh/end_quizzes`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
