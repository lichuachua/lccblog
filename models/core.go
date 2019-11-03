package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"time"
)

var (
	db *gorm.DB
)

//数据库初始化逻辑
func init() {
	//logs.Info("NI HAO INIT fangfa!")
	var err error
	//创建
	if err = os.MkdirAll("data", 0777); err != nil {
		panic("failed " + err.Error())
	}
	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/lichuachua?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()
	db.AutoMigrate(&User{}, &Note{}, &Message{}, &PraiseLog{}, &Code{})
	//如果数据库里面没用用户数据，我们新增一条admin记录
	var count int
	if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&User{
			Name: "李歘歘", //默认的管理员的名字，你可以使用这个名字发文章
			//邮箱
			Email: "lichuachua@qq.com", //默认登录的邮箱
			//密码
			Pwd: "c4ca4238a0b923820dcc509a6f75849b", //默认密码 1 ，已经umd5加密
			//头像地址
			Avatar: "/static/images/info-img.png", //默认头像
			//是否认证 例： lyblog 作者
			Role:   0,
			Status: 0,
		})
	}

}

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//

//package models
//
//import (
////"github.com/astaxie/beego/logs"
////"github.com/jinzhu/gorm"
////"os"
////_ "github.com/jinzhu/gorm/dialects/sqlite"
////"github.com/jinzhu/gorm"
//_ "github.com/go-sql-driver/mysql"
//"database/sql"
//"fmt"
//
//)
//
////var(
////	db *gorm.DB
////)
////数据库初始化逻辑
//func init()  {
//	////logs.Info("NI HAO INIT fangfa!")
//	//var err error
//	////创建
//	//if err=os.MkdirAll("data",0777);err!=nil{
//	//	panic("failed "+err.Error())
//	//}
//	//db, err = gorm.Open("sqlite3", "data/data.db")
//	//if err != nil {
//	//	panic("failed to connect database")
//	//}
//	////defer db.Close()
//
//	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/db_blog?charset=utf8")
//	if err != nil {
//		panic(err.Error())
//	}
//	//Db数据库连接池
//	var DB *sql.DB
//	defer db.Close()
//	DB, _ = sql.Open("mysql", "root:123456@tcp(localhost:3306)/db_beego?charset=utf8")
//	if err := DB.Ping(); err != nil {
//		fmt.Println("opon database fail")
//		return
//	}
//	fmt.Println("connnect success")
//
//}
