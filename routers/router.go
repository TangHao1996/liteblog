package routers

import (
	"liteblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	// beego.Router("/test", &controllers.MainController{}, "get:TestGetFunc")
	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(&controllers.IndexController{})
	beego.Include(&controllers.UserController{})
	beego.AddNamespace(
		beego.NewNamespace("note",
			beego.NSInclude(&controllers.NoteController{}),
		),
		beego.NewNamespace("comment",
			beego.NSInclude(&controllers.CommentController{}),
		),
		beego.NewNamespace("like",
			beego.NSInclude(&controllers.LikeController{}),
		),
	)

}
