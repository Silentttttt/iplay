package models

import (
	"time"
)

type Game struct {
	Id          int       // 赛事ID
	Type        int       // 赛事类型
	HomeTeam    int       // 主队
	VisitTeam   int       // 客队
	Begin       time.Time // 比赛开始时间
	End         time.Time // 比赛结束时间
	Description string    `orm:"size(32)"` // 赛事描述
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
}
