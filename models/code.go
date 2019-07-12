package models

type Code struct {
	Model
	UKey         string `gorm:"unique:not null"`
	Verification string
	Genre int
}
func SaveCode(code *Code) error {
	return db.Save(code).Error
}

func DeleteCode(key, verification string, genre int) error {
	return db.Delete(&Code{}, "u_key = ? And verification = ? And genre = ? ", key, verification, genre).Error
}

func QueryKeyAndCode(key, verification string) (code Code, err error) {
	return code, db.Where("u_key = ? And verification = ?", key, verification).Take(&code).Error
}

func QueryCodeByKey(key string) (code Code, err error) {
	return code, db.Where("u_key = ?", key).Take(&code).Error
}
