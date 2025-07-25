// controllers/user_controller.go
package controllers

import (
	"net/http"

	"smartschoolsystem.go/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&creds); err != nil {
		var raw map[string]interface{}
		c.BindJSON(&raw) // force parse and show what was received
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
			"raw":   raw,
			"note":  "Expected: username/password lowercase keys",
		})
		return
	}

	token, err := services.AuthenticateUser(creds.Username, creds.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Dashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the SmartSchool dashboard"})
}
