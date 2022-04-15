package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "wxEmoji/routers"
)

func main() {
	beego.Run()
}
