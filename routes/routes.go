package routes

import (
	"github.com/extrieve/tekken-babes-gin/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
    api := router.Group("/api")
    {
        api.GET("/battle", controllers.GetBattle)
        api.POST("/battle/vote", controllers.SubmitVote)
        api.GET("/characters/:id", controllers.GetCharacter)
        api.GET("/characters", controllers.GetCharacters)
        api.GET("/leaderboard", controllers.GetLeaderboard)
    }
}
