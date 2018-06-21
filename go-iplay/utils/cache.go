package utils

import (
	"time"

	"github.com/astaxie/beego/cache"
)

var cc cache.Cache

func InitCache() {
	cc, _ = cache.NewCache("memory", `{"interval":60}`)

}

func Put(key string, value interface{}, timeout time.Duration) error {
	return cc.Put(key, value, timeout)
}

func Get(key string) interface{} {
	return cc.Get(key)
}

func Delete(key string) error {
	return cc.Delete(key)
}

func IsExist(key string) bool {
	return cc.IsExist(key)
}
