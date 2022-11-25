package api

import (
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {

	u := r.Group("/user")
	{
		u.POST("/register",register)
		u.POST("/login",login)
		u.POST("/cpwd",changePassword)
		u.POST("/rpwd",retrievePassword)
	}
}