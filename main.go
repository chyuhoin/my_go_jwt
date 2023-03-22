package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"proj1/pkg/config"
)

func main() {
	r := gin.Default()

	config.RouterConfig(r)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}