package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//初始化
func init() {
	// orm.RegisterModel(new(User), new(Game), new(Player), new(PlayType), new(GameExtFootball), new(Quizzes), new(ChoiceOpt))
	orm.RegisterModel(new(User), new(PlayType), new(Player), new(Game), new(GameExtFootball), new(Quizzes), new(ChoiceOpt), new(UserQuizzes))
}

func TableName(name string) string {
	prefix := beego.AppConfig.String("db_dt_prefix")
	return prefix + name
}

func UserTBName() string {
	return TableName("user")
}

func GameTBName() string {
	return TableName("game")
}

func PlayerTBName() string {
	return TableName("player")
}

func PlayTypeTBName() string {
	return TableName("play_type")
}

func GameExtFootballTBName() string {
	return TableName("game_ext_football")
}

func QuizzesTBName() string {
	return TableName("quizzes")
}

func ChoiceOptTBName() string {
	return TableName("choice_opt")
}

func UserQuizzesTBName() string {
	return TableName("user_quizzes")
}
