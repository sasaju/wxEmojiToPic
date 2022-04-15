package models

type WxParam struct {
	//Id    int         `form:"-"`
	//Name  interface{} `form:"username"`
	//Age   int         `form:"age"`
	//Email string
	signature string
	timestamp string
	nonce     string
	echostr   string
}
