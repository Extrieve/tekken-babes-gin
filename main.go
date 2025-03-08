package main

import (
	"log"

	"github.com/extrieve/tekken-babes-gin/database"
	_ "github.com/extrieve/tekken-babes-gin/docs"
	"github.com/extrieve/tekken-babes-gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    database.Connect()
    database.SeedData()

    router := gin.Default()
    routes.RegisterRoutes(router)

    if err := router.Run(":8080"); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}