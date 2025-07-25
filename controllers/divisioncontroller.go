package controllers

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"smartschoolsystem.go/models"
	"smartschoolsystem.go/services"
)

func GetAllDivisions(c *gin.Context) {
	divisions, err := services.GetAllDivisions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch divisions"})
		return
	}
	c.JSON(http.StatusOK, divisions)
}
func GetDivisionById(c *gin.Context) {
	divisionId := c.Param("divisionId")
	idInt, err := strconv.Atoi(divisionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid division ID"})
		return
	}
	division, err := services.GetDivisionById(uint(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch division"})
		return
	}
	c.JSON(http.StatusOK, division)
}
func CreateDivision(c *gin.Context) {
	var division models.Division
	if err := c.ShouldBindJSON(&division); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	divisionId, err := services.CreateDivision(division)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create division"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"divisionId": divisionId})
}
func UpdateDivision(c *gin.Context) {
	var division models.Division
	if err := c.ShouldBindJSON(&division); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := services.UpdateDivision(division); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update division"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Division updated successfully"})
}

func DeleteDivision(c *gin.Context) {
	divisionId := c.Param("divisionId")
	idInt, err := strconv.Atoi(divisionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid division ID"})
		return
	}
	if err := services.DeleteDivision(uint(idInt)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete division"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Division deleted successfully"})
}
