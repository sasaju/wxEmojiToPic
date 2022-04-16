package main

import (
	"github.com/beego/beego/v2/adapter/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "wxEmoji/routers"
)

func main() {
	err := logs.SetLogger(logs.AdapterFile, `{"filename":"run.log"}`)
	if err != nil {
		logs.Error("日志配置失败")
	}
	beego.Run()
}
