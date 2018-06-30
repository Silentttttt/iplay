package models

import "time"

type Transaction struct {
	Id      int64
	UserID  int64
	Hash    string
	Status  int
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time
}
