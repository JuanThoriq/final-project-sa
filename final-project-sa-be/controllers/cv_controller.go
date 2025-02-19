package controllers

import (
	"net/http"
	"strconv"

	"final-project-sa-be/database"
	"final-project-sa-be/models"

	"github.com/gin-gonic/gin"
)

func CreateCV(c *gin.Context) {
	var input struct {
		Title    string `json:"title" binding:"required"`
		Content  string `json:"content"`
		Template string `json:"template"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Ambil userID dari context (diset oleh middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	cv := models.CV{
		UserID:   userID.(uint),
		Title:    input.Title,
		Content:  input.Content,
		Template: input.Template,
	}
	result := database.DB.Create(&cv)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, cv)
}

func GetCVs(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var cvs []models.CV
	result := database.DB.Where("user_id = ?", userID).Find(&cvs)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, cvs)
}

func GetCVByID(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CV ID"})
		return
	}
	var cv models.CV
	result := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&cv)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CV not found"})
		return
	}
	c.JSON(http.StatusOK, cv)
}

func UpdateCV(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CV ID"})
		return
	}
	var cv models.CV
	result := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&cv)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CV not found"})
		return
	}
	var input struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		Template string `json:"template"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cv.Title = input.Title
	cv.Content = input.Content
	cv.Template = input.Template
	database.DB.Save(&cv)
	c.JSON(http.StatusOK, cv)
}

func DeleteCV(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CV ID"})
		return
	}
	result := database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.CV{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "CV deleted successfully"})
}

// Contoh endpoint: POST /cv/:id/skills
// Body JSON: { "skill_ids": [1, 2, 3] }
func AddSkillsToCV(c *gin.Context) {
	// Ambil userID dari context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Ambil ID CV dari parameter URL
	idParam := c.Param("id")
	cvID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CV ID"})
		return
	}

	// Cari CV milik user
	var cv models.CV
	if err := database.DB.Where("id = ? AND user_id = ?", cvID, userID).First(&cv).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CV not found"})
		return
	}

	// Ambil list skill IDs dari body request
	var input struct {
		SkillIDs []uint `json:"skill_ids" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil skill berdasarkan input
	var skills []models.Skill
	if err := database.DB.Where("id IN ?", input.SkillIDs).Find(&skills).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching skills"})
		return
	}

	// Tambahkan skill ke CV menggunakan association
	if err := database.DB.Model(&cv).Association("Skills").Append(skills); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding skills to CV"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Skills added successfully", "cv": cv})
}