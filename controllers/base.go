package controllers

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"liteblog/models"
	"liteblog/syserror"
	"math/rand"
	"strconv"
	"time"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type BaseController struct {
	beego.Controller
	User    models.User
	IsLogin bool
}

type NestPreparer interface {
	NestPrepare()
}

func (this *BaseController) Prepare() {
	this.Data["Path"] = this.Ctx.Request.RequestURI
	u, ok := this.GetSession(SESSION_USER_KEY).(models.User)
	this.IsLogin = false
	if ok {
		this.User = u
		this.IsLogin = true
		this.Data["User"] = this.User
	}
	this.Data["islogin"] = this.IsLogin
	if a, ok := this.AppController.(NestPreparer); ok {
		a.NestPrepare()
	}
}

func (this *BaseController) Abort500(err error) {
	this.Data["error"] = err
	this.Abort("500")
}

func (this *BaseController) GetMustString(key, msg string) string {
	email := this.GetString(key, "")

	if len(email) == 0 {
		this.Abort500(errors.New(msg))

	}
	return email
}

func (this *BaseController) MustLogin() {
	if !this.IsLogin {
		this.Abort500(syserror.NoUserError{})
	}
}

type H map[string]interface{}

func (this *BaseController) JsonOk(msg, action string) {
	this.Data["json"] = H{
		"code":   0,
		"msg":    msg,
		"action": action,
	}
	this.ServeJSON()
}
func (this *BaseController) JsonOkH(msg string, data H) {
	if data == nil {
		data = H{}
	}
	data["code"] = 0
	data["msg"] = msg
	this.Data["json"] = data
	this.ServeJSON()
}

func (this *BaseController) UUID() string {
	u, err := uuid.NewV4()
	if err != nil {
		this.Abort500(syserror.New("系统错误", err))
	}
	return u.String()
}

func (this *BaseController) GetRandom() string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, 6)
	for i := 0; i < 6; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func (this *BaseController) MD5Util(password string) string {
	data := []byte(password)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)
	return md5Str
}

func (this *BaseController) GetRandomAvatar() string {
	rand.Seed(time.Now().Unix())
	number := strconv.Itoa(rand.Intn(10))
	//fmt.Println(number)
	avatar := "/static/images/" + number + ".jpg"
	return avatar
}
