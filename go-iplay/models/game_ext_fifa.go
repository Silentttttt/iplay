package models

import (
	"time"
)

// GameExtFIFA 世界杯赛事
type GameExtFIFA struct {
	Id     int
	GameId int // 赛事ID

	RegularResult int // 常规赛事结果 {0 主胜 1 平 2 客胜}
	FinalResult   int // 最终比赛结果 {0 主胜 1 平 2 客胜}

	HomeScore  int // 主队得分
	VisitScore int // 客队得分

	HomeScoreFirstHalf   int // 主队上半场得分
	VisitScoreFirstHalf  int // 客队上半场得分
	HomeScoreSecondHalf  int // 主队下半场得分
	VisitScoreSecondHalf int // 客队下半场得分

	IsOvertime         bool // 是否加时
	HomeScoreOvertime  int  // 加时主队得分
	VisitScoreOvertime int  // 加时客队得分

	IsPenalty         bool // 是否点球大战
	HomeScorePenalty  int  // 点球大战主队得分
	VisitScorePenalty int  // 点球大战客队得分

	Begin time.Time // 比赛开始时间
	End   time.Time // 比赛结束时间

	Description string    `orm:"size(32)"` // 赛事描述
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
}
