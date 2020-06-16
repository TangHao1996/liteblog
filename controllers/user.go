package controllers

import (
	"errors"
	"liteblog/models"
	"liteblog/syserror.go"
)

type UserController struct {
	BaseController
}

// @router /login [post]
func (this *UserController) Login() {
	//从post请求中得到
	email := this.GetMustString("email", "邮箱不能为空！")
	password := this.GetMustString("password", "密码不能为空！")

	user, err := models.QueryUserByEmailAndPassword(email, password)
	if err != nil {
		this.Abort500(syserror.New("登录失败！", err))
	}
	//采用gob编码
	this.SetSession(SESSION_USER_KEY, user)
	// this.Data["json"] = map[string]interface{}{
	// 	"code":   0,
	// 	"action": "/",
	// }
	this.JsonOK("欢迎", "/")
}

// @router /logout [get]
func (this *UserController) Logout() {
	this.MustLogin()
	this.DelSession(SESSION_USER_KEY)
	//临时重定向
	this.Redirect("/", 302)
}

// @router /register [post]
func (this *UserController) Register() {
	name := this.GetMustString("name", "用户名不能为空！")
	email := this.GetMustString("email", "邮箱不能为空！")
	password := this.GetMustString("password", "密码不能为空！")
	//若未abort，查重
	if u, err := models.QueryUserByName(name); err == nil && u.ID > 0 {
		this.Abort500(errors.New("用户名已存在"))
	}
	if u, err := models.QueryUserByEmail(email); err == nil && u.ID > 0 {
		this.Abort500(errors.New("邮箱已被注册"))
	}

	//数据库插入
	newUser := models.User{
		Name:     name,
		Email:    email,
		Password: password,
		Avatar:   "static/images/info-img.png",
		Role:     1,
	}
	if err := models.CreateUser(newUser); err != nil {
		this.Abort500(syserror.New("未知错误", err))
	}
	//成功
	this.JsonOK("注册成功", "/user")
}
