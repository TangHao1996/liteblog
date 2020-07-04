package controllers

import (
	"errors"
	"liteblog/models"
)

type LikeController struct {
	BaseController
}

func (this *LikeController) NextPrepare() {
	this.MustLogin()
}

//@router /:type/:key [post]
func (this *LikeController) Like() {
	t := this.Ctx.Input.Param(":type")
	key := this.Ctx.Input.Param(":key")

	//TODO 建立sql点赞log表，记录文章-用户点赞记录，查询是否已赞过
	switch t {
	case "note":
		models.IncrLike(key)
		like := models.QueryNoteLikeByKey(key)
		this.JsonOKData("", JsonData{
			"like": like,
		})
	default:
		this.Abort500(errors.New(t))
	}

}
