package controllers

import (
	"crypto/sha1"
	"github.com/beego/beego/v2/adapter/logs"
	beego "github.com/beego/beego/v2/server/web"
	"reflect"
	"sort"
	"wxEmoji/models"
)

type WxController struct {
	beego.Controller
}

func (c *WxController) Get() {
	params := models.WxParam{}
	if err := c.ParseForm(&params); err != nil {
		c.Ctx.WriteString("发生错误")
		logs.Error("token验证失败")
	}
	data := reflect.TypeOf(params)
	filed := data.Field(0)
	signature := c.GetString("signature", "未成功")
	timestamp := filed.Tag.Get("timestamp")
	nonce := filed.Tag.Get("nonce")
	echostr := filed.Tag.Get("echostr")
	token := "yes"

	list0 := []string{token, timestamp, nonce}
	sort.Strings(list0)
	sha := sha1.New()
	for i := range list0 {
		sha.Write([]byte(list0[i]))
	}
	bs := sha.Sum(nil)
	logs.Info("singnature=" + signature)
	if string(bs) == signature {
		c.Ctx.WriteString(echostr)
	} else {
		logs.Error("token未通过")
		c.Ctx.WriteString("")
	}
}
