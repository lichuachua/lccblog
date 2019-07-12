package controllers

import (
	"errors"
	"liteblog/models"
	"github.com/jinzhu/gorm"
	"liteblog/syserror"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

type NoteController struct {
	BaseController
}

func (this *NoteController) NextPrepare()  {
	this.MustLogin()
	if this.User.Role!=0 {
		this.Abort500(errors.New("权限不足"))
	}
}



//note
// @router /new [get]
func (this *NoteController) Index()  {
	this.Data["key"] = this.UUID()
	this.TplName="note_new.html"
}


//note
// @router /edit/:key [get]
func (this *NoteController) EditPage()  {

	key:= this.Ctx.Input.Param(":key")
	note,err:=models.QueryNoteByKeyAndUserId(key,int(this.User.ID))
	if err!=nil{
		this.Abort500(syserror.New("文章不存在！",err))
	}
	this.Data["note"] = note
	this.Data["key"]=key

	this.TplName="note_new.html"
}


// @router /save/:key [post]
func (this *NoteController) Save() {
	key :=this.Ctx.Input.Param(":key")
	title :=this.GetMustString("title","请输入标题！")
	content:=this.GetMustString("content","请输入文章内容！")

	note,err:=models.QueryNoteByKey(key)
	var n models.Note
	if err!=nil{
		if err == gorm.ErrRecordNotFound{
			n=models.Note{
				Key:key,
				Title:title,
				Content:content,
				UserID:int(this.User.ID),
				User:this.User,
			}
		}else {
			this.Abort500(syserror.New("保存失败2",err))
		}
	}else {
	note.Title = title
	note.Content = content
	n=note
	}


	//摘要
	n.Summary,_=getSummary(content)
	if err=models.SaveNote(&n);err!=nil {
		this.Abort500(syserror.New("保存失败3", err))
	}
	this.JsonOk("保存成功",fmt.Sprintf("/details/%s",key))
}

//note
// @router /del/:key [post]
func (this *NoteController) Del()  {
	key:=this.Ctx.Input.Param(":key")
	if err:=models.DeleteNoteByUserIdAndKey(key,int(this.User.ID));err!=nil{
		this.Abort500(syserror.New("删除失败",err))
	}
	this.JsonOk("删除成功","/")
}

func getSummary(html string) (string,error) {
	var bufbytes bytes.Buffer
	if _,err :=bufbytes.Write([]byte(html));err!=nil{
		return "",err
	}
	doc,err:=goquery.NewDocumentFromReader(&bufbytes)
	if err!=nil{
		return "",err
	}
	htmlstr:=doc.Find("body").Text()
	if len(htmlstr)>600 {
		htmlstr=htmlstr[:600]
	}
	return htmlstr,nil
}