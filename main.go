package main

import (
	"log"
	"myapp/config"
	"myapp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库连接
	err := config.InitDB()
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 创建Gin路由实例
	r := gin.Default()

	// 设置路由
	routes.SetupRoutes(r)

	// 启动服务器
	r.Run(":8080")
}
