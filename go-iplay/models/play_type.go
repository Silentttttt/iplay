package models

import (
	"time"
)

/*
 足球 - 世界杯
 足球 - 欧洲杯
*/
type PlayType struct {
	Id      int64
	Name    string    `orm:"size(32)"` // 类型名称
	Parent  int64     // 父类型
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}
