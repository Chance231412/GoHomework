package middleware

import (
	"github.com/gin-gonic/gin"
	"homework06/dao"
	"homework06/utils"
)

// ContactWithDatabase 连接数据库
func ContactWithDatabase(c *gin.Context) {
	err := dao.Init()
	if err!=nil {
		utils.RespFail(c,err)
		c.Abort()
		return
	}
}