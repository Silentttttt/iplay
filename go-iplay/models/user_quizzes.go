package models

import (
	"time"
)

type UserQuizzes struct {
	Id        int64
	User      *User      `orm:"rel(fk)"`
	ChoiceOpt *ChoiceOpt `orm:"rel(fk)"`
	Result    bool
	Money     float64   // 下注金额
	Reward    float64   // 竞猜奖励
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
}

func (uq *UserQuizzes) TableName() string {
	return UserQuizzesTBName()
}
