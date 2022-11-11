package apis

import (
	"github.com/gin-gonic/gin"
	"homework-05/dao"
)

func InitRouter() {
	r:=gin.Default()

	r.Use(dao.ContactWithDatabase)

	r.POST("/register",register)
	r.POST("/login",login)
	r.POST("/changePWD",changePassword)
	r.POST("/forgotPWD",dealWithForgotPwd)

	r.Run(":8848")
}
