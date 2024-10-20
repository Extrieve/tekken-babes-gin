package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/extrieve/tekken-babes-gin/database"
	"github.com/extrieve/tekken-babes-gin/models"
	"github.com/gin-gonic/gin"
)

func GetBattle(c *gin.Context) {
    var characters []models.Character
    result := database.DB.Find(&characters)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    if len(characters) < 2 {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Not enough characters in the database"})
        return
    }

    rand.Seed(time.Now().UnixNano())
    indices := rand.Perm(len(characters))[:2]
    characterOne := characters[indices[0]]
    characterTwo := characters[indices[1]]

    c.JSON(http.StatusOK, gin.H{
        "characterOne": characterOne,
        "characterTwo": characterTwo,
    })
}

func SubmitVote(c *gin.Context) {
    type VoteInput struct {
        WinnerID     uint `json:"winnerId" binding:"required"`
        LoserID      uint `json:"loserId" binding:"required"`
        CurrentStreak uint `json:"currentStreak"`
    }

    var input VoteInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update total wins for the winner
    var winner models.Character
    if err := database.DB.First(&winner, input.WinnerID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Winner not found"})
        return
    }
    winner.TotalWins++
    database.DB.Save(&winner)

    // Create a new battle record
    battle := models.Battle{
        CharacterOneID: input.WinnerID,
        CharacterTwoID: input.LoserID,
        WinnerID:       input.WinnerID,
        BattleTime:     time.Now(),
    }
    database.DB.Create(&battle)

    if input.CurrentStreak+1 < 5 {
        c.JSON(http.StatusOK, gin.H{
            "message":   "Vote recorded",
            "newStreak": input.CurrentStreak + 1,
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "message":     "Character wins!",
            "characterId": winner.ID,
        })
    }
}
