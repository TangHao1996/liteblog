package controllers

import (
	"liteblog/syserror.go"

	"github.com/astaxie/beego/logs"
)

type ErrorController struct {
	BaseController
}

//其他地方调用Abort("404")会进入这里
// ajax: {code:, msg:, reason:error}
func (this *ErrorController) Error404() {
	this.Data["content"] = "page not found"
	this.TplName = "error/404.html"
	//这段先不管
	if this.IsAjax() {
		this.jsonError(syserror.Error404{})
	}
}

func (this *ErrorController) Error500() {
	this.TplName = "error/500.html"
	//这个变量是在哪里加的？
	err, ok := this.Data["error"].(error)
	if !ok {
		err = syserror.New("unknown error", nil)
	}

	syserr, ok := err.(syserror.Error)
	if !ok {
		syserr = syserror.New(err.Error(), nil)
	}

	if syserr.ReasonError() != nil {
		logs.Info(syserr.Error(), syserr.ReasonError())
	}

	if this.IsAjax() {
		this.jsonError(syserr)
	} else {
		this.Data["content"] = syserr.Error()
	}
}

func (this *ErrorController) jsonError(err syserror.Error) {
	this.Ctx.Output.Status = 200
	this.Data["json"] = map[string]interface{}{
		"code": err.Code(),
		"msg":  err.Error(),
	}
	this.ServeJSON()
}
