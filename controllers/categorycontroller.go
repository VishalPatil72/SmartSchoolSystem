package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"smartschoolsystem.go/services"
)

func GetAllCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}
func GetSubCategories(c *gin.Context) {
	subCategories, err := services.GetSubCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subcategories"})
		return
	}
	c.JSON(http.StatusOK, subCategories)
}
func GetAllReligions(c *gin.Context) {
	religions, err := services.GetAllReligions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch religions"})
		return
	}
	c.JSON(http.StatusOK, religions)
}
