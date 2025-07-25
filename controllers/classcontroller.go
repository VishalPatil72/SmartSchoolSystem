package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"smartschoolsystem.go/models"
	"smartschoolsystem.go/services"
)

func GetAllClasses(c *gin.Context) {
	classes, err := services.GetAllClasses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch classes"})
		return
	}
	c.JSON(http.StatusOK, classes)
}
func GetClassById(c *gin.Context) {
	classId := c.Param("classId")
	idInt, err := strconv.Atoi(classId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}
	class, err := services.GetClassById(uint(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch class"})
		return
	}
	c.JSON(http.StatusOK, class)
}
func CreateClass(c *gin.Context) {
	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	classId, err := services.CreateClass(class)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create class"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"classId": classId})
}
func UpdateClass(c *gin.Context) {
	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := services.UpdateClass(class); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update class"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Class updated successfully"})
}
func DeleteClass(c *gin.Context) {
	classId := c.Param("classId")
	idInt, err := strconv.Atoi(classId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}
	if err := services.DeleteClass(uint(idInt)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete class"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}
