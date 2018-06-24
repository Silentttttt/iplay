package controllers

import (
	"iplay/go-iplay/utils"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

const (
	Success = "0"
	Fail    = "1"
)

type BaseController struct {
	beego.Controller
}

type JsonResult struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Ts   int64       `json:"ts"`
}

func (c *BaseController) json(code string, msg string, data interface{}) {
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
