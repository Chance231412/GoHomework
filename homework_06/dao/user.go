package dao

import (
	"homework06/model"
)

//SelectUserByName 通过用户名查找用户，若存在返回true，否则返回false
func SelectUserByName(name string,user *model.User)bool {
	err := DB.Where("name = ?",name).First(user).Error
	if err!=nil {
		user=nil
		return false
	}
	return true
}

//AddUser 将用户添加进数据库
func AddUser(u *model.User)error {
	err := DB.Create(u).Error
	return err
}

//ChangePwd 修改用户密码
func ChangePwd(name string,password string)error {
	var u model.User
	//根据主键查找
	//可以判断用户是否存在
	err := DB.First(&u,"name = ?",name).Error
	if err!=nil {
		return err
	}
	u.Password=password
	err = DB.Save(&u).Error
	if err!=nil {
		return err
	}
	return nil
}
