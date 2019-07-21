package controllers

import (
	"fmt"
	"liteblog/models"
	"liteblog/syserror"
)

type MessageController struct {
	BaseController
}

//@router /count [get]
func (this *MessageController) Count() {
	count, err := models.QueryMessagesCountByNoteKey("")
	if err != nil {
		this.Abort500(syserror.New("查询失败", err))
	}
	this.JsonOkH("查询成功", H{"count": count})
}

// @router /query [get]
func (this *MessageController) Query() {
	pageno, err := this.GetInt("pageno", 1)
	if err != nil || pageno < 1 {
		pageno = 1
	}
	pagesize, err := this.GetInt("pagesize", 10)
	if err != nil {
		pagesize = 10
	}

	ms, err := models.QueryPageMessagesByNoteKey("new", pageno, pagesize)
	if err != nil {
		this.Abort500(syserror.New("查询失败", err))
	}
	this.JsonOkH("查询成功", H{"data": ms})
}

//@router /?:key [post]
func (this *MessageController) NewMessage() {
	fmt.Print("开始了")
	this.MustLogin()
	key := this.Ctx.Input.Param(":key")
	content := this.GetMustString("content", "请输入内容！")
	k := this.UUID()
	m := &models.Message{
		Key:     k,
		NoteKey: key,
		User:    this.User,
		UserId:  int(this.User.ID),
		Content: content,
	}
	if err := models.SaveMessage(m); err != nil {
		this.Abort500(syserror.New("保存失败", err))
	}
	this.JsonOkH("保存成功", H{"data": m})
}
