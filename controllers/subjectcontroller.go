package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"smartschoolsystem.go/models"
	"smartschoolsystem.go/services"
)

func GetAllSubjects(c *gin.Context) {
	subjects, err := services.GetAllSubjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subjects"})
		return
	}
	c.JSON(http.StatusOK, subjects)
}
func GetSubjectById(c *gin.Context) {
	subjectId := c.Param("subjectId")
	idInt, err := strconv.Atoi(subjectId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject ID"})
		return
	}
	subject, err := services.GetSubjectById(uint(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subject"})
		return
	}
	c.JSON(http.StatusOK, subject)
}
func CreateSubject(c *gin.Context) {
	var subject models.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	subjectId, err := services.CreateSubject(subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subject"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"subjectId": subjectId})
}
func UpdateSubject(c *gin.Context) {
	var subject models.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := services.UpdateSubject(subject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update subject"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subject updated successfully"})
}

func DeleteSubject(c *gin.Context) {
	subjectId := c.Param("subjectId")
	idInt, err := strconv.Atoi(subjectId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject ID"})
		return
	}
	if err := services.DeleteSubject(uint(idInt)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete subject"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subject deleted successfully"})
}
