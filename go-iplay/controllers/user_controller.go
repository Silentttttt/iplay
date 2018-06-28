package controllers

import (
	"encoding/json"
	"fmt"
	"iplay/go-iplay/models"
	"iplay/go-iplay/utils"
	"strings"

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
		c.json(Success, "", &models.LoginResponseData{AuthToken: authToken, Username: user.Username, Avatar: user.Avatar})
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
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	fmt.Println(params)
	username := strings.TrimSpace(params.Username)
	password := strings.TrimSpace(params.Password)

	user, _ := models.GetByUsername(username)
	if user != nil {
		c.json(Fail, RegisterUserExistErr, nil)
	}

	m := &models.User{}
	m.Username = username
	m.Pwd = utils.Md5(password)
	m.Passphrase = utils.RandomString(10)
	m.HashAddress = createAddressWithPassphrase(m.Passphrase)

	o := orm.NewOrm()
	if _, err := o.Insert(m); err != nil {
		c.json(Fail, RegisterSystemErr, nil)
	}
	uuid, _ := uuid.NewV4()
	authToken := username + ":" + uuid.String()
	utils.Put(authToken, username, utils.Month)
	c.json(Success, "", &models.LoginResponseData{AuthToken: authToken, Username: m.Username})
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
		}
	}
	c.json(Fail, NeedLoginErr, nil)

}

func createAddressWithPassphrase(passphrase string) string {
	// params := &AreateAddressWithPassphraseRequest{passphrase: passphrase}
	// b, _ := json.Marshal(params)
	// req, err := http.NewRequest("POST", "url", bytes.NewBuffer(b))
	// req.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
	return ""
}
