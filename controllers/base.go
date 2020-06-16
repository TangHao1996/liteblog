package controllers

import (
	"errors"
	"liteblog/models"
	"liteblog/syserror.go"

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
	if u, ok := this.GetSession(SESSION_USER_KEY).(models.User); ok { //
		this.User = u
		this.IsLogin = true
		this.Data["User"] = this.User
	}
	this.Data["islogin"] = this.IsLogin
}

func (this *BaseController) Abort500(err error) {
	this.Data["error"] = err
	this.Abort("500")
}

func (this *BaseController) GetMustString(key, msg string) string {
	str := this.GetString(key)
	if len(str) == 0 {
		this.Abort500(errors.New(msg))
	}
	return str
}

func (this *BaseController) MustLogin() {
	if !this.IsLogin {
		this.Abort500(syserror.ErrorNoUser{})
	}
}

type JsonData map[string]interface{}

func (this *BaseController) JsonOK(msg, action string) {
	this.Data["json"] = map[string]interface{}{
		"code":   0,
		"msg":    msg,
		"action": action,
	}
	//将Data中的“json"发送
	this.ServeJSON()
}

func (this *BaseController) JsonOKData(msg string, data JsonData) {
	data["msg"] = msg
	data["code"] = 0

	this.Data["json"] = data
	//将Data中的“json"发送
	this.ServeJSON()
}
