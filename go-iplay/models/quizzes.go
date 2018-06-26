package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Quizzes struct {
	Id          int64        `json:"-""`
	Game        *Game        `orm:"rel(fk)" json:"game"`          // 赛事ID
	Instruction string       `orm:"size(512)" json:"instruction"` // 竞猜说明
	Begin       time.Time    `json:"begin"`                       // 竞猜开始时间
	End         time.Time    `json:"end"`                         // 竞猜结束时间
	Created     time.Time    `orm:"auto_now_add;type(datetime)" json:"-"`
	ChoiceOpt   []*ChoiceOpt `orm:"reverse(many)" json:"choice_opt"`
}

func (q *Quizzes) TableName() string {
	return QuizzesTBName()
}

func GetQuizzesListFromNow(gameID int64) (*Quizzes, error) {
	quizzes := Quizzes{}
	choiceOpts := []ChoiceOpt{}
	err := orm.NewOrm().QueryTable(QuizzesTBName()).Filter("game_id", gameID).Filter("begin__lt", time.Now()).Filter("end__gt", time.Now()).RelatedSel().One(&quizzes)
	if err != nil {
		return nil, err
	}
	_, err = orm.NewOrm().QueryTable(ChoiceOptTBName()).Filter("quizzes_id", quizzes.Id).All(&choiceOpts)
	if err != nil {
		return nil, err
	}
	quizzes.SetChoiceOpt(choiceOpts)
	return &quizzes, nil
}

func (q *Quizzes) SetChoiceOpt(choiceOpts []ChoiceOpt) {
	var choices []*ChoiceOpt
	for _, v := range choiceOpts {
		choices = append(choices, &v)
	}
	q.ChoiceOpt = choices
}
