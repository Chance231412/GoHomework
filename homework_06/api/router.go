package api

import (
	"github.com/gin-gonic/gin"
	"homework06/api/middleware"
)

func UserRouter(r *gin.Engine) {

	u := r.Group("/user")
	{
		//设置中间件连接数据库
		u.Use(middleware.ContactWithDatabase)

		u.POST("/register",register)
		u.POST("/login",login)
		u.POST("/cpwd",changePassword)
		u.POST("/rpwd",retrievePassword)
	}
}