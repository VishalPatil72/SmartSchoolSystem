package main

import (
	"log"

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
	routes.SetupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed:", err)
	}
}
