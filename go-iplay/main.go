package main

import (
	_ "iplay/go-iplay/routers"

	_ "iplay/go-iplay/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
