package models

type ChoiceOpt struct {
	Id      int64
	Name    string   `orm:"size(512)"` // 竞猜说明
	Odds    float32  //赔率
	Percent float32  //下注比例
	Count   int64    //下注人数
	Totoal  int64    //下注总金额
	Quizzes *Quizzes `orm:"rel(fk)"`
}

func (c *ChoiceOpt) TableName() string {
	return ChoiceOptTBName()
}
