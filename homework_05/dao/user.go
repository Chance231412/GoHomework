package dao

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"homework-05/models"
	"io"
	"os"
)

var (
	DB *os.File
	namePwd = make(map[string]string)
	phoneName = make (map[string]string)
	users = make(map[string]*models.User)
)

const (
	filePath  = "dao/users.txt"
)

func SelectUserByName(username string)bool{
	if namePwd[username]=="" {
		return false
	}
	return true
}

func SelectUserByPhone(phone string)(name string,err error) {
	name = phoneName[phone]
	if name=="" {
		return name,fmt.Errorf("该手机未绑定")
	}
	return name,nil
}

func ChangeUserPassword(name string,newPWD string) {
	namePwd[name]=newPWD
	users[name].Password=newPWD
}

func CheckPassword(username,pwd string)bool {
	if namePwd[username]==pwd {
		return true
	}
	return false
}

func AddUser(u models.User) {
	users[u.UserName]=&u
	fmt.Println(u)
}

func ContactWithDatabase(c *gin.Context) {
	err:=contactWithDatabase()
	if err!=nil {
		fmt.Println(err)
		c.Abort()
	}
	err=reloadData()
	if err!=nil {
		fmt.Println(err)
		return
	}
	err=DB.Close()
	if err!=nil {
		fmt.Println(err)
		return
	}
	c.Next()
	writeData()
}

func contactWithDatabase()error {
	var err error
	DB,err = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, os.ModeAppend|os.ModePerm)
	if err!=nil {
		return err
	}
	return nil
}

func reloadData()error {
	var c models.User
	reader := bufio.NewReader(DB)

	for {
		data,err:=reader.ReadString('\n')
		if err!=nil {
			if err==io.EOF {
				return nil
			}
			fmt.Println(err)
			return err
		}
		err=json.Unmarshal([]byte(data),&c)
		if err!=nil {
			fmt.Println(err)
			return err
		}
		namePwd[c.UserName]=c.Password
		phoneName[c.PhoneNumber]=c.UserName
		users[c.UserName]=&c
	}
}

func writeData() {
	var err error
	DB,err = os.OpenFile(filePath,os.O_WRONLY|os.O_TRUNC, 0666)
	if err!=nil {
		fmt.Println(err)
		return
	}
	for _,u:=range users {
		userinfo,err := json.Marshal(u)
		if err!=nil {
			fmt.Println(err)
			return
		}
		userinfo = append(userinfo,'\n')
		_,err=DB.Write(userinfo)
		if err!=nil {
			fmt.Println(err)
			return
		}
	}
}