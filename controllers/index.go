package controllers

type IndexController struct {
	BaseController
}

//注解路由，只要加上如下注释，则只用include该controller就能自动注册这些函数
//@router /  [get]
func (this *IndexController) Index() {
	this.TplName = "index.html" //views/tpl
}

//@router /message [get]
func (this *IndexController) Message() {
	this.TplName = "message.html" //views/tpl
}

//@router /about [get]
func (this *IndexController) About() {
	this.TplName = "about.html" //views/tpl
}
