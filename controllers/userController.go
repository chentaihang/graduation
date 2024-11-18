package controllers

import (
	"myapp/config"
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 注册处理函数
func Register(c *gin.Context) {
	id := c.Query("id")
	password := c.Query("password")
	type1 := c.Query("type")
	var teacher models.Teacher
	var student models.Student
	var err interface{}
	if type1 == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type不能为空"})
		return
	} else if type1 == "student" {
		student.StudentNumber, _ = strconv.Atoi(id)
		student.Password = password
		// 绑定JSON数据
		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else if type1 == "teacher" {
		teacher.TeacherID, _ = strconv.Atoi(id)
		teacher.Password = password
		// 绑定JSON数据
		if err := c.ShouldBindJSON(&teacher); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	// 创建用户
	if type1 == "student" {
		err = config.DB.Create(&student).Error
	} else if type1 == "teacher" {
		err = config.DB.Create(&teacher).Error
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "注册成功"})
}

// 登录处理函数
func Login(c *gin.Context) {
	id := c.Query("id")
	password := c.Query("password")
	type1 := c.Query("type")
	if type1 == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type不能为空"})
		return
	}
	if type1 == "student" {
		var student models.Student
		err := config.DB.Where("student_number = ? AND password = ?", id, password).First(&student).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "学生未找到"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
	} else if type1 == "teacher" {
		var teacher models.Teacher
		err := config.DB.Where("teacher_id = ? AND password = ?", id, password).First(&teacher).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "教师未找到"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的type"})
	}
}
