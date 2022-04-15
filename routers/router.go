package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"wxEmoji/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/wx", &controllers.WxController{})
}
