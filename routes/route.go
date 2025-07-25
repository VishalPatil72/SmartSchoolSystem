// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"smartschoolsystem.go/controllers"
	"smartschoolsystem.go/middleware"
)

func SetupRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
	}

	api := r.Group("/api")
	api.Use(middleware.JWTAuthMiddleware())
	{
		api.GET("/dashboard", controllers.Dashboard)
		// Add more protected routes here
	}
	masterapi := r.Group("/master")
	masterapi.Use(middleware.JWTAuthMiddleware())
	{
		masterapi.GET("/categories", controllers.GetAllCategories)
		masterapi.GET("/categories/:id", controllers.GetCategoryById)
		masterapi.POST("/categories", controllers.CreateCategory)
		masterapi.PUT("/categories/:id", controllers.UpdateCategory)
		masterapi.DELETE("/categories/:id", controllers.DeleteCategory)
	}
}
