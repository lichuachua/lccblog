package controllers

import (
	"errors"
	"github.com/astaxie/beego/utils"
	"gopkg.in/gomail.v2"
	"liteblog/models"
	"liteblog/syserror"
	"strconv"
	"strings"
)

type UesrController struct {
	BaseController
}

//@router /login [post]
func (this *UesrController) Login() {
	//email
	email := this.GetMustString("email", "邮箱不能为空！")
	//password
	pwd := this.GetMustString("password", "密码不能为空！")
	pwd = this.MD5Util(pwd)
	user, err := models.QueryUserByEmailAndPwd(email, pwd)
	if err != nil {
		this.Abort500(syserror.New("登陆失败!", err))
	}
	this.SetSession(SESSION_USER_KEY, user)

	this.JsonOk("登陆成功", "/")

}

//@router /reg [post]
func (this *UesrController) Reg() {
	//昵称，邮箱，密码，确认密码，都不能为空
	name := this.GetMustString("name", "昵称不能为空")
	email := this.GetMustString("email", "邮箱不能为空")
	password := this.GetMustString("password", "密码不能为空")
	password2 := this.GetMustString("password1", "确认密码不能为空")
	key := this.UUID()
	if strings.Compare(password, password2) != 0 {
		this.Abort500(errors.New("两次输入的密码不一样"))
	}
	if u, err := models.QueryUserByName(name); err == nil && u.ID > 0 {
		this.Abort500(errors.New("用户昵称已经存在"))
	}
	if u, err := models.QueryUserByEmail(email); err == nil && u.ID > 0 {
		this.Abort500(errors.New("用户邮箱已经存在"))
	}
	password = this.MD5Util(password)
	this.SendEmail1(email, key)
	//this.SendEmailBySender()
	if err := models.SaveUser(&models.User{
		UKey:   key,
		Name:   name,
		Email:  email,
		Pwd:    password,
		Avatar: this.GetRandomAvatar(),
		Role:   1,
		Status: -1,
	}); err != nil {
		this.Abort500(syserror.New("用户保存失败", err))
	}

	this.JsonOk("注册成功,请去邮箱激活！！！", "/user")
}

//@router /logout [get]
func (this *UesrController) Logout() {
	this.MustLogin()
	this.DelSession(SESSION_USER_KEY)
	this.Redirect("/", 302)
}

// @router /sendEmail [post]
func (this *UesrController) SendEmail(email, key string) {
	//email := this.GetMustString("email","email不能为空")
	code, err1 := models.QueryCodeByKey(key)
	if err1 != nil && code.ID > 0 {
		this.Abort500(syserror.New("已经发送请勿重复操作", err1))
	}
	verification := this.GetRandom()
	config := `{"username":"1065618302@qq.com","password":"eurbieigzrxobeeb","host":"smtp.qq.com","port":587,"secure":"SSL"}`
	temail := utils.NewEMail(config)
	temail.To = []string{email}
	temail.From = "1065618302@qq.com"
	temail.Subject = "用户激活"
	temail.HTML = "点击激活：" + "http://127.0.0.1:8080/activation?key=" + key + "&verification=" + verification
	err := temail.Send()
	//将code存入数据库
	if err := models.SaveCode(&models.Code{
		UKey:         key,
		Verification: verification,
		Genre:        0,
	}); err != nil {
		this.Abort500(syserror.New("发送失败", err))
	}
	if err != nil {
		this.Abort500(syserror.New("发送失败", err))
	}
}

func (this *UesrController) SendEmail1(email, key string) {
	code, err1 := models.QueryCodeByKey(key)
	if err1 != nil && code.ID > 0 {
		this.Abort500(syserror.New("已经发送请勿重复操作", err1))
	}
	verification := this.GetRandom()
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": "15534446431@163.com",
		"pass": "lichuachua1314",
		"host": "smtp.163.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", "lichuachua blog"+"<"+mailConn["user"]+">")                                                                             //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", email)                                                                                                                    //发送给多个用户
	m.SetHeader("Subject", "激活账户")                                                                                                              //设置邮件主题
	m.SetBody("text/html", "欢迎注册李歘歘博客系统，请点击以下链接激活账号，如非本人操作，请忽略："+"http://47.103.48.175:8080/activation?key="+key+"&verification="+verification) //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	//将code存入数据库
	if err := models.SaveCode(&models.Code{
		UKey:         key,
		Verification: verification,
		Genre:        0,
	}); err != nil {
		this.Abort500(syserror.New("发送失败", err))
	}
	if err != nil {
		this.Abort500(syserror.New("发送失败", err))
	}
}

// @router /activation [get]
func (this *UesrController) Activation(key, verification string) {
	_, err := models.QueryKeyAndCode(key, verification)
	if err != nil {
		this.Abort500(syserror.New("激活失败", err))
	}
	_, err1 := models.UpdatedStatus(key)
	if err1 != nil {
		this.Abort500(syserror.New("激活失败", err))
	}
	if err := models.DeleteCode(key, verification, 0); err != nil {
		this.Abort500(syserror.New("激活失败", err))
	}
	//this.JsonOk("激活成功","/user")
	this.TplName = "activation.html"
}
