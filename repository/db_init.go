package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/community?charset=utf8mb4&parseTime=True&loc=Local"
	// 打开dsn表示的数据库
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return err
}
