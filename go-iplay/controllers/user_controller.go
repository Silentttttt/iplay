package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"star/models"
	"star/utils"
	"strings"
	"time"

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

func (c *UserController) Login() {
	mobile := strings.TrimSpace(c.GetString("mobile"))
	password := strings.TrimSpace(c.GetString("password"))
	if len(mobile) == 0 || len(password) == 0 {
		c.json(Fail, "mobile or password is invalid.", nil)
	}
	password = utils.Md5(password)
	user, err := models.GetByMobileAndPassword(mobile, password)
	if user != nil && err == nil {
		uuid, _ := uuid.NewV4()
		authToken := mobile + ":" + uuid.String()
		utils.Put(authToken, mobile, utils.Month)
		c.json(Success, "", &AuthToken{AuthToken: authToken})
	} else {
		c.json(Fail, "mobile or password is invalid.", nil)
	}
}

func (c *UserController) Register() {
	mobile := strings.TrimSpace(c.GetString("mobile"))
	password := strings.TrimSpace(c.GetString("password"))
	code := strings.TrimSpace(c.GetString("code"))

	user, _ := models.GetByMobile(mobile)
	if user != nil {
		c.json(Fail, "手机号码已注册，请直接登录.", nil)
	}
	// TODO: 校验code
	if verifyCode(mobile, code) == false {
		c.json(Fail, "验证码不正确.", nil)
	}

	m := &models.User{}
	m.Mobile = mobile
	m.Pwd = utils.Md5(password)
	m.Passphrase = utils.RandomString(10)
	m.HashAddress = createAddressWithPassphrase(m.Passphrase)
	nick := fmt.Sprintf("star_%v", time.Now().Unix())
	m.NickName = nick

	o := orm.NewOrm()
	if _, err := o.Insert(m); err != nil {
		c.json(Fail, "注册失败", nil)
	}
	uuid, _ := uuid.NewV4()
	authToken := mobile + ":" + uuid.String()
	utils.Put(authToken, mobile, utils.Month)
	c.json(Success, "", &AuthToken{AuthToken: authToken})
}

func (c *UserController) Authentication() {
	authToken := strings.TrimSpace(c.GetString("auth_token"))
	name := strings.TrimSpace(c.GetString("name"))
	IDCard := strings.TrimSpace(c.GetString("ID"))

	if CheckAuthToken(authToken) {
		user, _ := models.GetByMobile(utils.Get(authToken).(string))
		if user.IdCard != "" {
			c.json(Fail, "不可重复认证", nil)
			return
		}
		user.RealName = name
		user.IdCard = IDCard
		user.Status = 1
		o := orm.NewOrm()
		if _, err := o.Update(user); err != nil {
			c.json(Fail, "实名认证失败", nil)
		}
	}
	c.json(Fail, "未授权", nil)

}

func (c *UserController) SendMobileCode() {
	mobile := strings.TrimSpace(c.GetString("mobile"))
	usage := strings.TrimSpace(c.GetString("usage"))
	if usage == "1" {
		user, _ := models.GetByMobile(mobile)
		if user != nil {
			c.json(Fail, "手机号码已注册，请直接登录", nil)
		}
	} else if usage == "2" {
		user, _ := models.GetByMobile(mobile)
		if user == nil {
			c.json(Fail, "账号不存在", nil)
		}
	} else {
		c.json(Fail, "参数错误", nil)
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	fmt.Println("code:", code)
	if err := c.sendSmsCode(mobile, usage, code); err != nil {
		c.json(Fail, "fail to send mobile code.", nil)
	}
	utils.Put(mobile, code, 3*time.Minute)

	c.json(Success, "", code)
}

func verifyCode(mobile, code string) bool {
	if !utils.IsExist(mobile) {
		return false
	}
	return code == utils.Get(mobile).(string)
}

func (c *UserController) userInfo() {
	c.json("", "", nil)
}

func (c *UserController) saveUserInfo() {
	c.json("", "", nil)
}

func (c *UserController) sendSmsCode(mobile string, usage string, code string) error {
	return nil
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
	return "dhkfjdkjkfj"
}
