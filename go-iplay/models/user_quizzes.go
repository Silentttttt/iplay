package models

import (
	"github.com/astaxie/beego/orm"
)

type UserQuizzes struct {
	Id        int64
	User      *User      `orm:"rel(fk)" json:"user"`
	ChoiceOpt *ChoiceOpt `orm:"rel(fk)" json:"choice_opt"`
	Quizzes   *Quizzes   `orm:"rel(fk)" json:"quizzes"`
	Game      *Game      `orm:"rel(fk)" json:"-"`
	Result    int        `json:"result"`  // 投注结果{0 待开奖 1 猜对 2 猜错}
	Money     int64      `json:"money"`   // 下注金额
	Reward    int64      `json:"reward"`  // 竞猜奖励
	Created   string     `json:"created"` //下注时间
}

type GameUserQuizzes struct {
	GameId int64
	Count  int
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

func GetUserQuizzesListByGame(userID int64, gameID int64) (*[]UserQuizzes, error) {
	quizzes := []UserQuizzes{}
	_, err := orm.NewOrm().QueryTable(UserQuizzesTBName()).Filter("user_id", userID).Filter("game_id", gameID).RelatedSel().All(&quizzes)
	if err != nil {
		return nil, err
	}
	return &quizzes, nil
}

func GetUserQuizzesGroupByGame(userID int64, gameId int64) (*[]GameUserQuizzes, error) {
	var gameUserQuizzes []GameUserQuizzes
	o := orm.NewOrm()
	if gameId > 0 {
		return &[]GameUserQuizzes{GameUserQuizzes{GameId: gameId, Count: 1}}, nil
	}
	_, err := o.Raw("SELECT game_id, count(*) FROM tb_user_quizzes WHERE user_id = ? group by game_id", userID).QueryRows(&gameUserQuizzes)
	if err != nil {
		return nil, err
	}
	return &gameUserQuizzes, err
}
