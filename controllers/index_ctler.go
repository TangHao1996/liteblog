package controllers

import (
	"liteblog/models"
	"liteblog/syserror.go"
)

type IndexController struct {
	BaseController
}

//注解路由，只要加上如下注释，则只用include该controller就能自动注册这些函数
//@router /  [get]
func (this *IndexController) Index() {
	limit := 5
	page, err := this.GetInt("page", 1)
	if err != nil || page <= 0 {
		page = 1
	}

	searchWord := this.GetString("title")

	notes, err := models.QueryNotesByPage(searchWord, page, limit)
	if err != nil {
		this.Abort500(syserror.New("获取博客失败", err))
	}

	count, err := models.QueryNoteCount(searchWord)
	if err != nil {
		this.Abort500(syserror.New("获取博客失败", err))
	}

	totalPage := count / limit
	if count%limit != 0 {
		totalPage++
	}

	this.Data["totalPage"] = totalPage
	this.Data["page"] = page
	this.Data["notes"] = notes
	this.TplName = "index.html" //views/tpl
}

//@router /message [get]
func (this *IndexController) Message() {
	this.TplName = "message.html"
}

//@router /about [get]
func (this *IndexController) About() {
	this.TplName = "about.html"
}

//@router /user [get]
func (this *IndexController) User() {
	if !this.IsLogin {
		this.Redirect("/login", 302)
	}
	this.TplName = "user.html"
}

//@router /login [get]
func (this *IndexController) GoLogin() {
	this.TplName = "login.html"
}

//@router /authentication [post]
func (this *IndexController) Login() {
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

// @router /goRegister [get]
func (this *IndexController) GoRegister() {
	this.TplName = "register.html"
}

// @router /play [get]
func (this *IndexController) GoPlay() {
	this.TplName = "play.html"
}

// @router /details/:key [get]
func (this *IndexController) Details() {
	this.TplName = "details.html"
	key := this.Ctx.Input.Param(":key")
	note, err := models.QueryNoteByKey(key)
	if err != nil {
		this.Abort500(syserror.New("未找到内容", err))
	}
	note.Visit++
	this.Data["note"] = note
	this.Data["title"] = note.Title
	this.Data["updatedAt"] = note.UpdatedAt
	this.Data["content"] = note.Content
	this.Data["author"] = note.User.Name
	this.Data["visit"] = note.Visit
	this.Data["like"] = note.Like

	models.SaveNote(&note)
}

//@router /comment/:key [post]
func (this *IndexController) GoComment() {

	//this.TplName = "comment.html"
}
