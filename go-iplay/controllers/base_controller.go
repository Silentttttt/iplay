package controllers

import (
	"iplay/go-iplay/utils"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

const (
	Success = 200
	Fail    = 500

	// Error
	NeedLoginErr = "用户未登陆，请登陆以后再执行操作"

	RegisterUserExistErr = "该用户已注册，请直接登录"
	RegisterSystemErr    = "系统错误，注册失败"

	LoginParamsErr = "登陆失败，用户名或者密码无效"

	IDCardAuthenticationRepeatErr = "不可以重复实名认证"
	IDCardAuthenticationErr       = "系统错误，实名认证失败"
)

type BaseController struct {
	beego.Controller
}

type JsonResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Ts   int64       `json:"ts"`
}

func (c *BaseController) json(code int, msg string, data interface{}) {
	c.Data["json"] = &JsonResult{code, msg, data, time.Now().Unix()}
	c.ServeJSON()
	c.StopRun()
}

func CheckAuthToken(authToken string) bool {
	if !utils.IsExist(authToken) {
		return false
	}

	return strings.Split(authToken, ":")[0] == utils.Get(authToken)
}
