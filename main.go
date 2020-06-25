package main

import (
	"encoding/gob"
	"liteblog/models"
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
	//添加html模板函数，用于每次访问动态生成html文件
	beego.AddFuncMap("equal", func(x, y string) bool {
		x = strings.Trim(x, "/")
		y = strings.Trim(y, "/")
		return strings.Compare(x, y) == 0
	})
	beego.AddFuncMap("add", func(x, y int) int {
		return x + y
	})
}

//考虑换成redis
func initSession() {
	gob.Register(models.User{}) //为了session能保存user结构体
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteblog"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
	//beego.BConfig.WebConfig.Session.SessionCookieLifeTime
}
