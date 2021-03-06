package controllers

import (
	"iplay/go-iplay/models"
	"iplay/go-iplay/utils"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

const (
	Success   = 200
	Fail      = 500
	NeedLogin = 501

	FreeToken = 2018 * 1000

	// Error
	NeedLoginErr = "用户未登陆，请登陆以后再执行操作"

	RegisterUserExistErr         = "该用户已注册，请直接登录"
	RegisterSystemErr            = "系统错误，注册失败"
	RegisterCreateHashAddressErr = "注册失败，创建用户钱包地址出错"
	RegisterTransferFreeTokenErr = "注册失败，发放免费token出错"

	LoginParamsErr = "登陆失败，用户名或者密码无效"

	IDCardAuthenticationRepeatErr = "不可以重复实名认证"
	IDCardAuthenticationErr       = "系统错误，实名认证失败"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) json(code int, msg string, data interface{}) {
	c.Data["json"] = &models.Response{code, msg, data, time.Now().Unix()}
	c.ServeJSON()
	c.StopRun()
}

func CheckAuthToken(authToken string) bool {
	if !utils.IsExist(authToken) {
		return false
	}

	return strings.Split(authToken, ":")[0] == utils.Get(authToken)
}
