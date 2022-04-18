package controllers

import (
	"crypto/sha1"
	"github.com/beego/beego/v2/adapter/logs"
	beego "github.com/beego/beego/v2/server/web"
	"sort"
	"strings"
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
	sha.Write([]byte(strings.Join(list0, "")))
	bs := sha.Sum(nil)
	if string(bs) == signature {
		c.Ctx.WriteString(echostr)
	} else {
		logs.Error("token未通过")
		logs.Error("signature=" + signature + "e=" + echostr)
		c.Ctx.WriteString("")
	}
}
