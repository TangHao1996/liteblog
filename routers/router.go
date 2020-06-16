package routers

import (
	"liteblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	// beego.Router("/test", &controllers.MainController{}, "get:TestGetFunc")
	beego.Include(&controllers.IndexController{})
	beego.Include(&controllers.UserController{})
	beego.ErrorController(&controllers.ErrorController{})
}
