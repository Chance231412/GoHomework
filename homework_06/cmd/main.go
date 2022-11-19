package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"homework06/api"
)

func main() {
	r := gin.Default()

	api.UserRouter(r)

	err := r.Run()

	if err!=nil {
		fmt.Println(err)
		return
	}
}
