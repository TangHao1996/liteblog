package controllers

import (
	"liteblog/models"

	"github.com/astaxie/beego"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type BaseController struct {
	beego.Controller
	User    models.User
	IsLogin bool
}

//所有继承该controller的controller都会先调用该prepare再执行自己的handler
func (this *BaseController) Prepare() {
	this.Data["Path"] = this.Ctx.Request.RequestURI //添加模板变量

	this.IsLogin = false
	if u, ok := this.GetSession(SESSION_USER_KEY).(models.User); ok {
		this.User = u
		this.IsLogin = true
		this.Data["User"] = this.User
	}
}

func (this *BaseController) Abort500(err error) {
	this.Data["error"] = err
	this.Abort("500")
}
