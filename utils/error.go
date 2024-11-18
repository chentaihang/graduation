package utils

import "github.com/gin-gonic/gin"

// 格式化错误响应
func HandleError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}
