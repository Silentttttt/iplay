package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id          int64
	Username    string `orm:"size(32)"`
	Realname    string `orm:"size(32)"`
	IdCard      string `orm:"size(32)"` // 提现需要通过身份认证
	Pwd         string `json:"-"`
	Status      int    `orm:"default(0)"` // {0:注册未实名 1:已实名}
	Mobile      string `orm:"size(16)"`   // 提现需要通过手机认证或者邮箱认证
	Passphrase  string `json:"-"`         // 用户钱包的passphrase
	HashAddress string `orm:"size(256)"`  // 用户钱包地址
	Email       string `orm:"size(256)"`  // 提现需要通过手机认证或者邮箱认证
	Avatar      string `orm:"size(256)"`
}

func (u *User) TableName() string {
	return UserTBName()
}

// GetByUserNameAndPassword get user by username and password
func GetByUserNameAndPassword(mobile, userpwd string) (*User, error) {
	m := User{}
	err := orm.NewOrm().QueryTable(UserTBName()).Filter("username", mobile).Filter("pwd", userpwd).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetByUsername get user by username
func GetByUsername(username string) (*User, error) {
	m := User{}
	err := orm.NewOrm().QueryTable(UserTBName()).Filter("username", username).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetByIDCard get user by id_card
func GetByIDCard(IDCard string) (*User, error) {
	m := User{}
	err := orm.NewOrm().QueryTable(UserTBName()).Filter("id_card", IDCard).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
