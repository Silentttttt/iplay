package models

import (
	"time"
)

type Player struct {
	Id       int64     `json:"-"`               // 战队ID
	Name     string    `orm:"size(32)"`         // 战队名称
	PlayType *PlayType `orm:"rel(fk)" json:"-"` // 战队类型
	Logo     string    `orm:"size(256)"`        // 战队logo
	Created  time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
}

func (p *Player) TableName() string {
	return PlayerTBName()
}
