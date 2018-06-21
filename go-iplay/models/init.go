package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//初始化
func init() {
	orm.RegisterModel(new(User), new(Article))
}

func TableName(name string) string {
	prefix := beego.AppConfig.String("db_dt_prefix")
	return prefix + name
}

func UserTBName() string {
	return TableName("user")
}

func ArticleTBName() string {
	return TableName("article")
}
