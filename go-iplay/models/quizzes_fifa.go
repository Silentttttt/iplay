package models

import (
	"time"
)

// QuizzesFIFA 世界杯竞猜 具体的不同竞猜玩法可以增加字段
type QuizzesFIFA struct {
	Id     int
	GameId int // 赛事ID

	HomeWinOdds  float32 // 主胜赔率
	tieOdds      float32 // 平局赔率
	VisitWinOdds float32 // 客胜赔率

	Created time.Time `orm:"auto_now_add;type(datetime)"`
}
