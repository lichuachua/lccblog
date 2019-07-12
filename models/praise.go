package models

import (
	"github.com/jinzhu/gorm"
	"liteblog/syserror"
)

type PraiseLog struct {
	Model
	UserId int
	Key string
	Table string
	Flag bool
}

type TempPraise struct {
	Praise int
}

func UpdatePraise(table, key string, userid int) (pcnt int,err error) {
	d:=db.Begin()
	defer func() {
		if err!=nil {
			d.Rollback()
		}else {
			d.Commit()
		}
	}()
	//需要判断是否点过赞，点过赞就算，不点过就不算
	//查询点赞的表，看是否有记录
	var p PraiseLog
	err = d.Model(&PraiseLog{}).Where("`key` = ? and `table` = ? and user_id = ?",key,table,userid).Take(&p).Error
	if err==gorm.ErrRecordNotFound {
		p=PraiseLog{
			Key:key,
			Table:table,
			UserId:userid,
			Flag:false,
		}
	}else if err!=nil {
		return 0,err
	}
	if p.Flag {
		return 0,syserror.HasPraiseError{}
	}
	p.Flag=true
	//保存
	if err = d.Save(&p).Error;err!=nil {
		return 0,err
	}
	//更新文章或者留言表的点赞数量
	var (
		ppp  TempPraise
		)

	err = d.Table(table).Where("`key` = ?",key).Select("praise").Scan(&ppp).Error
	if err!=nil {
		return 0,err
	}
	pcnt = ppp.Praise+1
	if err = d.Table(table).Where("`key` = ?",key).Update("praise",pcnt).Error;err!=nil{
		return 0,err
	}
	return pcnt,nil
}