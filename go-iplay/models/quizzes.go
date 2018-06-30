package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Quizzes struct {
	Id          int64        `json:"id"`
	Game        *Game        `orm:"rel(fk)" json:"-"`             // 赛事ID
	Instruction string       `orm:"size(512)" json:"instruction"` // 竞猜说明
	Begin       string       `json:"begin"`                       // 竞猜开始时间
	End         string       `json:"end"`                         // 竞猜结束时间
	Created     time.Time    `orm:"auto_now_add;type(datetime)" json:"-"`
	ChoiceOpt   []*ChoiceOpt `orm:"reverse(many)" json:"choice_opt"`
}

func (q *Quizzes) TableName() string {
	return QuizzesTBName()
}

func GetQuizzesById(id int64) (*Quizzes, error) {
	o := orm.NewOrm()
	m := Quizzes{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func GetQuizzesListFromNow(gameID int64) (*[]Quizzes, error) {

	quizzes := []Quizzes{}
	choiceOpts := []ChoiceOpt{}
	_, err := orm.NewOrm().QueryTable(QuizzesTBName()).Filter("game_id", gameID).Filter("end__gt", time.Now()).RelatedSel().All(&quizzes)
	if err != nil {
		return nil, err
	}

	for k := range quizzes {
		_, err = orm.NewOrm().QueryTable(ChoiceOptTBName()).Filter("quizzes_id", quizzes[k].Id).All(&choiceOpts)
		if err != nil {
			return nil, err
		}
		quizzes[k].SetChoiceOpt(choiceOpts)
	}

	return &quizzes, nil
}

func GetAllQuizzes() (*[]Quizzes, error) {
	quizzes := []Quizzes{}
	choiceOpts := []ChoiceOpt{}
	_, err := orm.NewOrm().QueryTable(QuizzesTBName()).RelatedSel().All(&quizzes)
	if err != nil {
		return nil, err
	}

	for k := range quizzes {
		_, err = orm.NewOrm().QueryTable(ChoiceOptTBName()).Filter("quizzes_id", quizzes[k].Id).All(&choiceOpts)
		if err != nil {
			return nil, err
		}
		quizzes[k].SetChoiceOpt(choiceOpts)
	}

	return &quizzes, nil
}

func (q *Quizzes) SetChoiceOpt(choiceOpts []ChoiceOpt) {
	var choices []*ChoiceOpt
	for k := range choiceOpts {
		choices = append(choices, &choiceOpts[k])
	}
	q.ChoiceOpt = choices
}
