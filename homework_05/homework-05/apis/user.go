package apis

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"homework-05/apis/middleware"
	"homework-05/dao"
	"homework-05/models"
	"homework-05/utils"
	"time"
)

func register(c *gin.Context) {
	//获取表单信息
	var form models.User
    err := c.ShouldBind(&form)
	if err!=nil {
		utils.RespFail(c,"verification failed")
		return
	}

	//检查用户名是否存在
	if dao.SelectUserByName(form.UserName) {
		utils.RespFail(c,"该用户名已被注册")
		return
	}

	//检查密码是否合法
	if err=checkPwdCorrect(form.Password);err!=nil {
		utils.RespFail(c,fmt.Sprintf("%v",err))
		return
	}

	//将该用户添加进数据库
	dao.AddUser(form)
	//if err!=nil {
	//	utils.RespFail(c,"注册失败")
	//	return
	//}

	//注册成功
	utils.RespSuccess(c,"注册成功")
}

func login(c *gin.Context) {
	//获取表单信息
	username := c.PostForm("username")
	if username=="" {
		utils.RespFail(c,"未输入用户名")
		return
	}
	password := c.PostForm("password")
	if password=="" {
		utils.RespFail(c,"未输入密码")
		return
	}

	//在数据库中查找该用户名称
	if !dao.SelectUserByName(username) {
		utils.RespFail(c,"该用户不存在")
		return
	}

	//匹配表单密码和数据库中的密码
	if !dao.CheckPassword(username,password) {
		utils.RespFail(c,"密码输入错误")
		return
	}

	//登陆成功,设置jwt
	// 创建一个我们自己的声明
	claim := models.MyClaims{
		Username: username, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 过期时间
			Issuer:    "Yxh",                                // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.RespSuccess(c, tokenString)
}

func changePassword(c *gin.Context) {
	var form models.User
	var username string
	err := c.ShouldBind(&form)
	if err!=nil {
		utils.RespFail(c,"verification failed")
	}
	newPassword := form.Password
	username,err = dao.SelectUserByPhone(form.PhoneNumber)
	if err!=nil {
		utils.RespFail(c,"该手机号未绑定用户")
		return
	}
	dao.ChangeUserPassword(username,newPassword)
	utils.RespSuccess(c,"修改成功")
}

func dealWithForgotPwd(c *gin.Context){
	changePassword(c)
	utils.RespSuccess(c,"修改成功")
}

func checkPwdCorrect(pwd string)error {
	if len(pwd)<6 {
		return fmt.Errorf("the password's length is too short")
	}
	return nil
}
