package sysinit

import "star/utils"

func init() {
	//初始化缓存
	utils.InitCache()
	//初始化数据库
	InitDatabase()
}
