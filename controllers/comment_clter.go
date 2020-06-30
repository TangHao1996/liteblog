package controllers

import (
	"liteblog/models"
	"liteblog/syserror.go"
)

type CommentController struct {
	BaseController
}

// @router /new/:notekey [post]
func (this *CommentController) New() {
	this.MustLogin()
	content := this.GetMustString("content", "内容不能为空")
	key := this.UUID()
	notekey := this.Ctx.Input.Param(":notekey")
	c := &models.Comment{
		UserId:  int(this.User.ID),
		User:    this.User,
		Key:     key,
		NoteKey: notekey,
		Content: content,
	}

	if err := models.CreateComment(c); err != nil {
		this.Abort500(syserror.New("创建评论失败", err))
	}
	this.JsonOKData("", JsonData{
		"code": 0,
	})
}
