package routes

import (
	"myapp/controllers"

	"github.com/gin-gonic/gin"
)

// 设置路由
func SetupRoutes(r *gin.Engine) {
	// 用户注册路由
	r.POST("/register", controllers.Register)
	// 用户登录路由
	r.POST("/login", controllers.Login)
	// 获取学生选题信息
	r.GET("/user/:id", controllers.GetStudentTopicDetails)
	// 获取详细选题信息
	r.GET("/topic/:id", controllers.GetTopics)
}
