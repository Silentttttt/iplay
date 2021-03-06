package controllers

import (
	"encoding/json"
	"fmt"
	"iplay/go-iplay/models"
	smartcontract "iplay/go-iplay/smartContract"
	"iplay/go-iplay/utils"
	"iplay/go-iplay/wallet"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	uuid "github.com/satori/go.uuid"
)

type UserController struct {
	BaseController
}

func (c *UserController) URLMapping() {
	c.Mapping("login", c.Login)
	c.Mapping("reg", c.Register)
}

// Login user login
// @Title Login
// @Description user login
// @Param   data body models.LoginParams true "user login request params"
// @Success 200 {object} models.LoginResponse
// @Failure 500
// @router /login [post]
func (c *UserController) Login() {
	var params models.LoginParams
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	username := strings.TrimSpace(params.Username)
	password := strings.TrimSpace(params.Password)
	if len(username) == 0 || len(password) == 0 {
		c.json(Fail, LoginParamsErr, nil)
	}
	password = utils.Md5(password)
	user, err := models.GetByUserNameAndPassword(username, password)
	if user != nil && err == nil {
		uuid, _ := uuid.NewV4()
		authToken := username + ":" + uuid.String()
		utils.Put(authToken, username, utils.Month)
		c.json(Success, "", &models.LoginResponseData{AuthToken: authToken, User: user})
	} else {
		c.json(Fail, LoginParamsErr, nil)
	}
}

// Register for user register
// @Title Register
// @Description user register
// @Param   data body models.LoginParams true "user register request params"
// @Success 200 {object} models.LoginResponse
// @Failure 500
// @router /reg [post]
func (c *UserController) Register() {
	var params models.LoginParams
	var err error
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	fmt.Println(params)
	username := strings.TrimSpace(params.Username)
	password := strings.TrimSpace(params.Password)

	user, _ := models.GetByUsername(username)
	if user != nil {
		c.json(Fail, RegisterUserExistErr, nil)
		return
	}

	m := &models.User{}
	m.Username = username
	m.Pwd = utils.Md5(password)
	m.Passphrase = utils.RandomString(10)
	// m.HashAddress = createAddressWithPassphrase(m.Passphrase)
	m.HashAddress, err = wallet.CreateAccount(m.Passphrase)
	if err != nil {
		logs.Error("[Register]Failed to Create hash address,", err)
		c.json(Fail, RegisterCreateHashAddressErr, nil)
		return
	}

	o := orm.NewOrm()
	o.Begin()
	var userID int64
	if userID, err = o.Insert(m); err != nil {
		c.json(Fail, RegisterSystemErr, nil)
		return
	}
	m.Id = userID
	// 注册成功 送用户2018*1000NAS
	_, err = smartcontract.Transfer(o, m.HashAddress, FreeToken)
	if err != nil {
		o.Rollback()
		logs.Error("[Register]Failed to Transfer free token,", err)
		c.json(Fail, RegisterTransferFreeTokenErr, nil)
		return
	}
	o.Commit()
	m.Balance = FreeToken
	if _, err = o.Update(m); err != nil {
		logs.Error("[Register]Failed to update balance,", err)
	}
	uuid, _ := uuid.NewV4()
	authToken := username + ":" + uuid.String()
	utils.Put(authToken, username, utils.Month)

	c.json(Success, "", &models.LoginResponseData{AuthToken: authToken, User: m})
}

// IDCardAuthentication 实名认证
func (c *UserController) IDCardAuthentication() {
	var params models.IDCardAuthenticationParams
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	authToken := strings.TrimSpace(params.AuthToken)
	name := strings.TrimSpace(params.Name)
	IDCard := strings.TrimSpace(params.IdCard)

	if CheckAuthToken(authToken) {
		user, _ := models.GetByUsername(utils.Get(authToken).(string))
		if user.IdCard != "" {
			c.json(Fail, IDCardAuthenticationRepeatErr, nil)
			return
		}
		user.Realname = name
		user.IdCard = IDCard
		user.Status = 1
		o := orm.NewOrm()
		if _, err := o.Update(user); err != nil {
			c.json(Fail, IDCardAuthenticationErr, nil)
			return
		}
	}
	c.json(NeedLogin, NeedLoginErr, nil)

}
