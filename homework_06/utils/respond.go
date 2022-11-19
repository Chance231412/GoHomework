package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespFail(c *gin.Context,err error) {
	message := fmt.Sprintf("[WARNING] an error occurred:%v", err)
	c.JSON(http.StatusInternalServerError,gin.H{
		"status":500,
		"message":message,
	})
}

func RespSuccess(c *gin.Context,message string) {
	c.JSON(http.StatusOK,gin.H{
		"status":200,
		"message":message,
	})
}
