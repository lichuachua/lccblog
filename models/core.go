package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	//"bytes"
	//"github.com/PuerkitoBio/goquery"
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
	db, err = gorm.Open("sqlite3", "data/data.db")
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()
	db.AutoMigrate(&User{}, &Note{}, &Message{}, &PraiseLog{}, &Code{})
	//如果数据库里面没用用户数据，我们新增一条admin记录
	var count int
	if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&User{
			Name: "李歘歘",
			//邮箱
			Email: "lcc1314..",
			//密码
			Pwd: "be07dee52c22ca2fd494285440880b1b",
			//头像地址
			Avatar: "/static/images/info-img.png",
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
