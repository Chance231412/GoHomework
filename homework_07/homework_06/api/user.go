package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"homework06/api/middleware"
	"homework06/dao"
	"homework06/model"
	"homework06/utils"
	"time"
)

//register 注册用户
func register(c *gin.Context) {
	//获取表单信息
	var form model.User
	err := c.ShouldBind(&form)
	if err!=nil {
		utils.RespFail(c,err)
		return
	}

	var u model.User
	//验证表单用户名是否已经被注册
	ok := dao.SelectUserByName(form.Name,&u)
	if ok {
		utils.RespFail(c,fmt.Errorf("该用户名已存在"))
		return
	}

	//验证用户密码是否符合要求
	if err=checkPwdCorrect(form.Password);err!=nil {
		utils.RespFail(c,err)
		return
	}

	//将用户信息添加到数据库
	err = dao.AddUser(&form)
	if err!=nil {
		utils.RespFail(c,err)
		return
	}

	utils.RespSuccess(c,"注册成功")
}

//login 用户登录
func login(c *gin.Context) {
	//获取表单信息
	name := c.PostForm("name")
	if name=="" {
		utils.RespFail(c,fmt.Errorf("表单信息不全"))
	}
	password := c.PostForm("password")
	if password=="" {
		utils.RespFail(c,fmt.Errorf("表单信息不全"))
	}

	var u model.User
	//验证表单用户名是否已经被注册
	ok := dao.SelectUserByName(name,&u)
	if !ok {
		utils.RespFail(c,fmt.Errorf("该用户不存在"))
		return
	}

	//验证用户的密码输入是否正确
	if u.Password!=password {
		utils.RespFail(c,fmt.Errorf("密码输入错误"))
		return
	}

	//登陆成功
	claim := model.MyClaims{
		Name: name, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 过期时间
			Issuer:    "Lb",                                // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(middleware.Secret)
	if err!=nil {
		utils.RespFail(c,err)
		return
	}
	utils.RespSuccess(c, tokenString)
}

//changePassword 修改密码
func changePassword(c *gin.Context) {
	//获取表单中的用户名，原密码和新密码
	name := c.PostForm("name")
	if name=="" {
		utils.RespFail(c,fmt.Errorf("表单信息不全"))
	}
	password := c.PostForm("password")
	if password=="" {
		utils.RespFail(c,fmt.Errorf("表单信息不全"))
	}
	newPwd := c.PostForm("newPassword")
	if newPwd=="" {
		utils.RespFail(c,fmt.Errorf("表单信息不全"))
	}

	//检验用户是否存在
	var u model.User
	ok := dao.SelectUserByName(name,&u)
	if !ok {
		utils.RespFail(c,fmt.Errorf("该用户不存在"))
		return
	}

	//检查密码是否正确
	if  password!= u.Password {
		utils.RespFail(c,fmt.Errorf("密码不正确"))
		return
	}

	//更新用户密码
	err := dao.ChangePwd(name,newPwd)
	if err!=nil {
		utils.RespFail(c,err)
		return
	}

	//修改密码成功
	utils.RespSuccess(c,"修改密码成功")
}

// retrievePassword 根据绑定的手机号找回密码
func retrievePassword(c *gin.Context) {
	//获取表单信息
	name := c.PostForm("name")
	if name=="" {
		utils.RespFail(c,fmt.Errorf("表单信息不全"))
	}
	phone := c.PostForm("phone")
	if phone=="" {
		utils.RespFail(c,fmt.Errorf("表单信息不全"))
	}

	//检查用户是否存在，并拿到用户信息
	var u model.User
	ok := dao.SelectUserByName(name,&u)
	if !ok {
		utils.RespFail(c,fmt.Errorf("该用户不存在"))
		return
	}

	//检查用户的电话是否匹配
	if phone!=u.Phone {
		utils.RespFail(c,fmt.Errorf("手机号码输入错误"))
		return
	}

	//返回用户密码
	utils.RespSuccess(c,u.Password)
}

//checkPwdCorrect 检查密码格式，这里就检查了长度
func checkPwdCorrect(pwd string)error {
	if len(pwd)<6 {
		return fmt.Errorf("the password's length is too short")
	}
	return nil
}