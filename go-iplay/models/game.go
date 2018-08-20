package models

import (
	"time"

	"iplay/go-iplay/utils"

	"github.com/astaxie/beego/orm"
)

type Game struct {
	Id          int64     `json:"id"`                       // 赛事ID
	PlayType    *PlayType `orm:"rel(fk)" json:"type"`       // 赛事类型
	HomeTeam    *Player   `orm:"rel(fk)" json:"home_team"`  // 主队
	VisitTeam   *Player   `orm:"rel(fk)" json:"visit_team"` // 客队
	HomeScore   int       `json:"home_score"`
	VisitScore  int       `json:"visit_score"`
	Status      int       `json:"status"`                      // {0 竞猜中 1 进行中 2 已结束}
	Begin       string    `json:"begin"`                       // 比赛开始时间
	End         string    `json:"end"`                         // 比赛结束时间
	Description string    `orm:"size(256)" json:"description"` // 赛事描述
	Created     time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
}

func (g *Game) TableName() string {
	return GameTBName()
}

func GetGameById(id int64) (*Game, error) {
	m := Game{}
	err := orm.NewOrm().QueryTable(GameTBName()).Filter("id", id).RelatedSel().One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func GetGameListFromNow(pageNo int) (*utils.Page, error) {
	games := []Game{}
	begin := DefaultPageSize * (pageNo - 1)
	number, err := orm.NewOrm().QueryTable(GameTBName()).OrderBy("-begin").Limit(DefaultPageSize, begin).RelatedSel().All(&games)
	if err != nil {
		return nil, err
	}

	page := utils.PageUtil(int(number), pageNo, DefaultPageSize, games)
	return &page, nil
}
