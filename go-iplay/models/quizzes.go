package models

import (
	"time"
)

type Quizzes struct {
	Id          int64
	Game        *Game        `orm:"rel(fk)"`   // 赛事ID
	Instruction string       `orm:"size(512)"` // 竞猜说明
	Begin       time.Time    // 竞猜开始时间
	End         time.Time    // 竞猜结束时间
	Created     time.Time    `orm:"auto_now_add;type(datetime)"`
	ChoiceOpt   []*ChoiceOpt `orm:"reverse(many)"`
}

func (q *Quizzes) TableName() string {
	return QuizzesTBName()
}
