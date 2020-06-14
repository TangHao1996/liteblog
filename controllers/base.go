package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

//所有继承该controller的controller都会先调用该prepare再执行自己的handler
func (this *BaseController) Prepare() {
	this.Data["Path"] = this.Ctx.Request.RequestURI //添加模板变量
}
