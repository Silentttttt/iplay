package models

import (
	"time"
)

type Game struct {
	Id          int64     // 赛事ID
	PlayType    *PlayType `orm:"rel(fk)"` // 赛事类型
	HomeTeam    *Player   `orm:"rel(fk)"` // 主队
	VisitTeam   *Player   `orm:"rel(fk)"` // 客队
	Begin       time.Time // 比赛开始时间
	End         time.Time // 比赛结束时间
	Description string    `orm:"size(256)"` // 赛事描述
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
}
