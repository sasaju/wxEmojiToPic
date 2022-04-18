package controllers

import (
	"crypto/sha1"
	"github.com/beego/beego/v2/adapter/logs"
	beego "github.com/beego/beego/v2/server/web"
	"sort"
)

type WxController struct {
	beego.Controller
}

func (c *WxController) Get() {
	signature := c.GetString("signature", "未成功")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	echostr := c.GetString("echostr")
	token := "yes"

	list0 := []string{token, timestamp, nonce}
	sort.Strings(list0)
	sha := sha1.New()
	for i := range list0 {
		sha.Write([]byte(list0[i]))
	}
	bs := sha.Sum(nil)
	if string(bs) == signature {
		c.Ctx.WriteString(echostr)
	} else {
		logs.Error("token未通过")
		c.Ctx.WriteString("")
	}
}
