package models

import (
	"time"
)

type Player struct {
	Id      int       // 战队ID
	Name    string    `orm:"size(32)"` // 战队名称
	Type    int       // 战队类型
	Logo    string    `orm:"size(256)"` // 战队logo
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}
