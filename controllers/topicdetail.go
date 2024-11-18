package controllers

import (
	"myapp/config"
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTopics(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var result struct {
		Student      models.Student
		Topic        models.Topic
		Teachers     []models.Teacher
		Tasks        []models.Task
		References   []models.Reference
		Arrangements []models.Arrangement
		Commit       models.Commit
	}

	// 查询学生信息
	if err := config.DB.Where("student_number = ?", studentID).First(&result.Student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	result.Topic.ID = result.Student.TopicID

	// 查询课题信息
	if err := config.DB.Where("topic_id = ?", result.Student.TopicID).First(&result.Topic).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Topic not found"})
		return
	}

	// 查询指导教师信息
	if err := config.DB.Joins("JOIN teacher_topic_relation ON teacher_topic_relation.teacher_id = leader_teachers.teacher_id").
		Where("teacher_topic_relation.topic_id = ?", result.Topic.ID).Find(&result.Teachers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch teachers"})
		return
	}

	// 查询任务信息
	if err := config.DB.Where("topic_id = ?", result.Topic.ID).Find(&result.Tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	// 查询参考文献信息
	if err := config.DB.Where("topic_id = ?", result.Topic.ID).Find(&result.References).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch references"})
		return
	}

	// 查询任务安排计划信息
	if err := config.DB.Where("topic_id = ?", result.Topic.ID).Find(&result.Arrangements).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch arrangements"})
		return
	}

	// 查询审核评价信息
	if err := config.DB.Where("topic_id = ?", result.Topic.ID).First(&result.Commit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch commit"})
		return
	}

	c.JSON(http.StatusOK, result)
}
