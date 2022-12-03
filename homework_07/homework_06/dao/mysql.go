package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"homework06/model"
)

var DB *gorm.DB

//ConnMysql 链接mysql数据库
func ConnMysql()error {
	//连接数据库
	dsn := "lb:wjLb231412@tcp(127.0.0.1:3306)/mydatabase"
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

//MysqlSelectUserByName 通过用户名查找用户，若存在返回true，否则返回false
func MysqlSelectUserByName(name string,user *model.User)bool {
	err := ConnMysql()
	if err!=nil {
		fmt.Println(err)
		//不怎么会处理错误
		return false
	}
	err = DB.Where("name = ?",name).First(user).Error
	if err!=nil {
		user=nil
		return false
	}
	return true
}

//AddUser 将用户添加进数据库
func AddUser(u *model.User)error {
	err := ConnMysql()
	if err!=nil {
		fmt.Println(err)
		//不怎么会处理错误
		return err
	}
	err = DB.Create(u).Error
	return err
}

//ChangePwd 修改用户密码,先在mysql中更新数据，然后再更新redis的数据
func ChangePwd(name string,password string)error {
	err := ConnMysql()
	if err!=nil {
		fmt.Println(err)
		//不怎么会处理错误
		return err
	}
	var u model.User
	//根据主键查找
	//可以判断用户是否存在
	err = DB.First(&u,"name = ?",name).Error
	if err!=nil {
		return err
	}
	//修改数据
	u.Password=password
	err = DB.Save(&u).Error
	if err!=nil {
		return err
	}
	//更新redis数据
	err=UpdateData(name,"password",password)
	if err!=nil {
		fmt.Println(err)
		return err
	}
	return nil
}
