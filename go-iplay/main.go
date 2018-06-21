package main

import (
	_ "star/routers"

	_ "star/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
