package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"iplay/go-iplay/models"
	"iplay/go-iplay/utils"
	"net/http"
	"strings"

	"github.com/astaxie/beego/orm"
	uuid "github.com/satori/go.uuid"
)

// UserController User info
type UserController struct {
	BaseController
}

// AuthToken auth token
type AuthToken struct {
	AuthToken string
}

// AreateAddressWithPassphraseRequest ..
type AreateAddressWithPassphraseRequest struct {
	passphrase string
}

// LoginParams login params
type LoginParams struct {
	username string
	password string
}

// IDCardAuthenticationParams IDCardAuthentication params
type IDCardAuthenticationParams struct {
	auth_token string
	name       string
	id_card    string
}

// Login for user login
func (c *UserController) Login() {
	var params LoginParams
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	username := strings.TrimSpace(params.username)
	password := strings.TrimSpace(params.password)
	if len(username) == 0 || len(password) == 0 {
		c.json(Fail, LoginParamsErr, nil)
	}
	password = utils.Md5(password)
	user, err := models.GetByUserNameAndPassword(username, password)
	if user != nil && err == nil {
		uuid, _ := uuid.NewV4()
		authToken := username + ":" + uuid.String()
		utils.Put(authToken, username, utils.Month)
		c.json(Success, "", &AuthToken{AuthToken: authToken})
	} else {
		c.json(Fail, LoginParamsErr, nil)
	}
}

// Register for user register
func (c *UserController) Register() {
	var params LoginParams
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	username := strings.TrimSpace(params.username)
	password := strings.TrimSpace(params.password)

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
	c.json(Success, "", &AuthToken{AuthToken: authToken})
}

// IDCardAuthentication 实名认证
func (c *UserController) IDCardAuthentication() {
	var params IDCardAuthenticationParams
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	authToken := strings.TrimSpace(params.auth_token)
	name := strings.TrimSpace(params.name)
	IDCard := strings.TrimSpace(params.id_card)

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
	params := &AreateAddressWithPassphraseRequest{passphrase: passphrase}
	b, _ := json.Marshal(params)
	req, err := http.NewRequest("POST", "url", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return ""
}
