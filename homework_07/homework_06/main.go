package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"homework06/api"
	"homework06/dao"
)

func main() {
	r := gin.Default()
	//链接redis
	err := dao.ConnRedis()
	if err!=nil {
		fmt.Println(err)
		return
	}
	//路由
	api.UserRouter(r)
	//启动服务器
	err = r.Run("0.0.0.0:80")
	if err!=nil {
		fmt.Println(err)
		return
	}
}
