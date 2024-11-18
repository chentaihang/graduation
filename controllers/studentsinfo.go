package controllers

// 查询学生选题信息（不包括详细信息）

import (
	"fmt"
	"myapp/config"
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetStudentID(studentID int) (*models.Student, error) {
	var student models.Student
	err := config.DB.Where("student_number = ?", studentID).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func getopic(topicID int, student *models.Student) error {
	// 使用结构体接收查询结果

	var result models.TopicResult

	err := config.DB.Raw(`
        SELECT 
            topic_id,
            name,
            type,
            description,
            DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') AS create_date,
            DATE_FORMAT(updated_at, '%Y-%m-%d %H:%i:%s') AS update_date
        FROM topics 
        WHERE topic_id = ?`, topicID).Scan(&result).Error

	if err != nil {
		return err
	}

	// 手动映射到 Topic 结构体
	models.Topicrevise(student, &result)

	return nil
}

func GetStudentTopicDetails(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的学生ID"})
		return
	}

	student, err := GetStudentID(studentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "学生未找到"})
		return
	}
	err = getopic(student.TopicID, student)
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}
