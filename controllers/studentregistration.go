package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"smartschoolsystem.go/models"
	"smartschoolsystem.go/services"
)

// GetAllStudents handles GET requests to retrieve all students
func GetAllStudents(c *gin.Context) {
	students, err := services.GetAllStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve students"})
		return
	}
	c.JSON(http.StatusOK, students)
}

// GetStudentById handles GET requests to retrieve a student by ID
func GetStudentById(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("studentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	student, err := services.GetStudentById(studentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve student"})
		return
	}

	c.JSON(http.StatusOK, student)
}
func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	studentId, err := services.CreateStudent(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"studentId": studentId})
}
func UpdateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.UpdateStudent(student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
}
func DeleteStudent(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("studentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	if err := services.DeleteStudent(studentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}
