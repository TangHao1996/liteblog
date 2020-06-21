package controllers

import (
	"bytes"
	"errors"
	"liteblog/models"
	"liteblog/syserror.go"

	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
)

type NoteController struct {
	BaseController
}

func (this *NoteController) NextPrepare() {
	this.MustLogin()
	if this.User.Role != ADMIN {
		this.Abort500(errors.New("权限不足"))
	}
}

// @router /new [get]
func (this *NoteController) Index() {
	this.Data["key"] = this.UUID()
	this.TplName = "note_new.html"
}

// @router /save/:key [post]
func (this *NoteController) Save() {
	key := this.Ctx.Input.Param(":key")
	title := this.GetMustString("title", "标题不能为空！")
	content := this.GetMustString("content", "内容不能为空！")

	var newNote models.Note
	if note, err := models.QueryNoteByKey(key); err != nil {
		if err == gorm.ErrRecordNotFound { //新文章
			newNote = models.Note{
				Key:     key,
				Title:   title,
				Content: content,
				UserId:  int(this.User.ID),
			}
			if err = models.CreateNote(&newNote); err != nil {
				this.Abort500(syserror.New("创建失败", err))
			}
		} else {
			this.Abort500(syserror.New("保存失败", err))
		}
	} else {
		//key已存在
		note.Title = title
		note.Content = content
		newNote = note
	}

	newNote.Summary, _ = getSummary(content)

	if err := models.SaveNote(&newNote); err != nil {
		this.Abort500(syserror.New("保存失败", err))
	}

	this.JsonOKData("保存成功", JsonData{})
}

func getSummary(html string) (string, error) {
	var bufbytes bytes.Buffer
	_, err := bufbytes.Write([]byte(html))
	if err != nil {
		return "", err
	}
	doc, err := goquery.NewDocumentFromReader(&bufbytes)
	if err != nil {
		return "", err
	}
	htmlstr := doc.Find("body").Text()
	if len(htmlstr) > 600 {
		htmlstr = htmlstr[:600]
	}

	return htmlstr, nil
}
