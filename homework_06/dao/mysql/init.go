package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"homework06/model"
)

var DB *gorm.DB

func Init()error {
	//连接数据库
	dsn := "LB:Lbqazwsx12345@tcp(127.0.0.1:3306)/homework06"
	db,err := gorm.Open(mysql.Open(dsn))
	if err!=nil {
		fmt.Println(err)
		return err
	}

	//判断数据表是否存在，不存在则创建数据表
	err=db.AutoMigrate(&model.User{})
	if err!=nil {
		return err
	}

	DB = db
	return nil
}