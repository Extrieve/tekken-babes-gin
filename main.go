package main

import (
	"log"

	"github.com/extrieve/tekken-babes-gin/database"
	"github.com/extrieve/tekken-babes-gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    database.Connect()

    // Seed data if necessary
    database.SeedData()

    router := gin.Default()

    // Register routes
    routes.RegisterRoutes(router)

    // Start the server
    err := router.Run(":8080")
    if err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
