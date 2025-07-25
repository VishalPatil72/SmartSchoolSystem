package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"smartschoolsystem.go/config"
	"smartschoolsystem.go/database"
	"smartschoolsystem.go/routes"
)

func main() {
	config.LoadEnv() // load from .env
	db := database.InitMySQL()
	defer db.Close()

	r := gin.Default()
	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.SetupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed:", err)
	}
}
