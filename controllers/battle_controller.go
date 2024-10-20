package controllers

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/extrieve/tekken-babes-gin/database"
	"github.com/extrieve/tekken-babes-gin/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBattle(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Count total documents
    count, err := database.CharacterCollection.CountDocuments(ctx, bson.M{})
    if err != nil || count < 2 {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Not enough characters in the database"})
        return
    }

    // Get two random characters
    rand.Seed(time.Now().UnixNano())
    skip := rand.Int63n(count - 1)

    opts := options.FindOptions{
        Limit: &[]int64{2}[0],
        Skip:  &skip,
    }

    cursor, err := database.CharacterCollection.Find(ctx, bson.M{}, &opts)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer cursor.Close(ctx)

    var characters []models.Character
    if err = cursor.All(ctx, &characters); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if len(characters) < 2 {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve two characters"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "characterOne": characters[0],
        "characterTwo": characters[1],
    })
}

func SubmitVote(c *gin.Context) {
    type VoteInput struct {
        WinnerID     string `json:"winnerId" binding:"required"`
        LoserID      string `json:"loserId" binding:"required"`
        CurrentStreak int    `json:"currentStreak"`
    }

    var input VoteInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    winnerObjectID, err := primitive.ObjectIDFromHex(input.WinnerID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid winner ID"})
        return
    }

    // Update total wins for the winner
    filter := bson.M{"_id": winnerObjectID}
    update := bson.M{"$inc": bson.M{"total_wins": 1}}

    result := database.CharacterCollection.FindOneAndUpdate(ctx, filter, update)
    if result.Err() != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Winner not found"})
        return
    }

    // Create a new battle record
    battle := models.Battle{
        CharacterOneID: winnerObjectID,
        CharacterTwoID: winnerObjectID, // You might need to adjust this
        WinnerID:       winnerObjectID,
        BattleTime:     time.Now(),
    }
    _, err = database.BattleCollection.InsertOne(ctx, battle)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record battle"})
        return
    }

    if input.CurrentStreak+1 < 5 {
        c.JSON(http.StatusOK, gin.H{
            "message":   "Vote recorded",
            "newStreak": input.CurrentStreak + 1,
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "message":     "Character wins!",
            "characterId": input.WinnerID,
        })
    }
}
