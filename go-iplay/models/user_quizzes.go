package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type UserQuizzes struct {
	Id        int64
	User      *User      `orm:"rel(fk)" json:"user"`
	ChoiceOpt *ChoiceOpt `orm:"rel(fk)" json:"choice_opt"`
	Quizzes   *Quizzes   `orm:"rel(fk)" json:"quizzes"`
	Result    bool       `json:"result"`
	Money     int64      `json:"money"`  // 下注金额
	Reward    int64      `json:"reward"` // 竞猜奖励
	Created   time.Time  `orm:"auto_now_add;type(datetime)" json:"-"`
}

func (uq *UserQuizzes) TableName() string {
	return UserQuizzesTBName()
}

func GetUserQuizzesList(userID int64) (*[]UserQuizzes, error) {
	quizzes := []UserQuizzes{}
	_, err := orm.NewOrm().QueryTable(UserQuizzesTBName()).Filter("user_id", userID).RelatedSel().All(&quizzes)
	if err != nil {
		return nil, err
	}
	return &quizzes, nil
}
