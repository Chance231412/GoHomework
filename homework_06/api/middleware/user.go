package middleware

import (
	"github.com/gin-gonic/gin"
	"homework06/dao/mysql"
	"homework06/utils"
)

// ContactWithDatabase 连接数据库
func ContactWithDatabase(c *gin.Context) {
	err := mysql.Init()
	if err!=nil {
		utils.RespFail(c,err)
		c.Abort()
		return
	}
}