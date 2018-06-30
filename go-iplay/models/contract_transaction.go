package models

import "time"

type ContractTransaction struct {
	Id       uint64
	From     string
	To       string
	Value    string
	Function string
	Args     string
	Hash     string
	Status   uint8
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time
}

func (c *ContractTransaction) TableName() string {
	return ContractTransactionTBName()
}
