package main

import (
	_ "liteblog/models"
	_ "liteblog/routers"
	"strings"

	"github.com/astaxie/beego"
)

func main() {
	initTemplat()
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
