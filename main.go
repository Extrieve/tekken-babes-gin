package main

import (
	"log"

	"github.com/extrieve/tekken-babes-gin/database"
	_ "github.com/extrieve/tekken-babes-gin/docs" // Import generated docs
	"github.com/extrieve/tekken-babes-gin/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Tekken Hotness Battle API
// @version         1.0
// @description     This is the API documentation for the Tekken Hotness Battle application.
// @termsOfService  http://swagger.io/terms/
// @contact.name    Your Name
// @contact.email   youremail@example.com
// @license.name    MIT License
// @license.url     https://opensource.org/licenses/MIT
// @host            localhost:8080
// @BasePath        /api
func main() {
    database.Connect()
    database.SeedData()

    router := gin.Default()

    // Register routes
    routes.RegisterRoutes(router)

    // Swagger endpoint
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Start the server
    err := router.Run(":8080")
    if err != nil {
        log.Fatal("Failed to start server:", err)
    }
}