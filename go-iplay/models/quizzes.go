package models

import (
	"time"
)

type Quizzes struct {
	Id          int
	GameId      int       // 赛事ID
	Instruction string    `orm:"size(512)"` // 竞猜说明
	Begin       time.Time // 竞猜开始时间
	End         time.Time // 竞猜结束时间
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	choices     ChoiceOpt
}

type ChoiceOpt struct {
	name    string  `orm:"size(512)"` // 竞猜说明
	odds    float32 //赔率
	percent float32 //下注比例
	count   int64   //下注人数
	totoal  int64   //下注总金额
}
