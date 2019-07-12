package models

//import "github.com/jinzhu/gorm"

type User struct {
	Model
	UKey      string `gorm:"unique:not null"`
	Name string `gorm:"unique_index" json:"name"`
	Email string `gorm:"unique_index" json:"email"`
	Pwd string `json:"-"`
	Avatar string `json:"avatar"`
	Role int `json:"role"`//0代表管理员，1代表用户
	Status    int    `json:"status"` //0代表禁用，1代表正常
}

func QueryUserByEmailAndPwd(email,pwd string) (user User,err error) {
	return user,db.Where("email = ? and pwd = ? and status = 0",email,pwd).Take(&user).Error
}

func QueryUserByName(name string) (user User,err error) {
	return user,db.Where("name = ?",name).Take(&user).Error
}

func QueryUserByEmail(email string) (user User,err error) {
	return user,db.Where("email = ?",email).Take(&user).Error
}

func SaveUser(user *User) error {
	return db.Save(user).Error
}

func UpdatedStatus(key string) (user User, err error) {
	return user, db.Model(&user).Where("u_key = ?", key).Update("status", 0).Take(&user).Error
}