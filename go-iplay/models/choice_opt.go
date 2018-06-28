package models

import "github.com/astaxie/beego/orm"

type ChoiceOpt struct {
	Id      int64    `json:"-"`
	Name    string   `orm:"size(512)" json:"name"` // 竞猜说明
	Odds    float32  `json:"odds"`                 //赔率
	Percent float32  `json:"-"`                    //下注比例
	Count   int64    `json:"count"`                //下注人数
	Totoal  int64    `json:"totoal"`               //下注总金额
	Quizzes *Quizzes `orm:"rel(fk)" json:"-"`
}

func (c *ChoiceOpt) TableName() string {
	return ChoiceOptTBName()
}

func GetChoiceOptById(id int64) (*ChoiceOpt, error) {
	o := orm.NewOrm()
	m := ChoiceOpt{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
