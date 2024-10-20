package main

import (
	"log"

	"github.com/extrieve/tekken-babes-gin/database"
	"github.com/extrieve/tekken-babes-gin/models"
	"github.com/extrieve/tekken-babes-gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    database.Connect()

    // Run migrations
    err := database.DB.AutoMigrate(&models.Character{}, &models.Battle{})
    if err != nil {
        log.Fatal("Failed to migrate database: ", err)
    }

    router := gin.Default()

    // Register routes
    routes.RegisterRoutes(router)

    // Start the server
    router.Run(":8080")
}
