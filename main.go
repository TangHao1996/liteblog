package main

import (
	_ "liteblog/models"
	_ "liteblog/routers"
	"strings"

	"github.com/astaxie/beego"
)

func main() {
	initTemplat()
	initSession()
	beego.Run()
}

func initTemplat() {
	//添加html模板函数
	beego.AddFuncMap("equal", func(x, y string) bool {
		x = strings.Trim(x, "/")
		y = strings.Trim(y, "/")
		return strings.Compare(x, y) == 0
	})
}

//考虑换成redis
func initSession() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteblog"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}
