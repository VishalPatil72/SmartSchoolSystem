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
		// category
		masterapi.GET("/categories", controllers.GetAllCategories)
		masterapi.GET("/categories/:id", controllers.GetCategoryById)
		masterapi.POST("/categories", controllers.CreateCategory)
		masterapi.PUT("/categories/:id", controllers.UpdateCategory)
		masterapi.DELETE("/categories/:id", controllers.DeleteCategory)
		// class
		masterapi.GET("/classes", controllers.GetAllClasses)
		masterapi.GET("/classes/:classId", controllers.GetClassById)
		masterapi.POST("/classes", controllers.CreateClass)
		masterapi.PUT("/classes/:classId", controllers.UpdateClass)
		masterapi.DELETE("/classes/:classId", controllers.DeleteClass)
		// division
		masterapi.GET("/divisions", controllers.GetAllDivisions)
		masterapi.GET("/divisions/:divisionId", controllers.GetDivisionById)
		masterapi.POST("/divisions", controllers.CreateDivision)
		masterapi.PUT("/divisions/:divisionId", controllers.UpdateDivision)
		masterapi.DELETE("/divisions/:divisionId", controllers.DeleteDivision)
		//	subject
		masterapi.GET("/subjects", controllers.GetAllSubjects)
		masterapi.GET("/subjects/:subjectId", controllers.GetSubjectById)
		masterapi.POST("/subjects", controllers.CreateSubject)
		masterapi.PUT("/subjects/:subjectId", controllers.UpdateSubject)
		masterapi.DELETE("/subjects/:subjectId", controllers.DeleteSubject)

	}
}
