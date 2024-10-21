package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/extrieve/tekken-babes-gin/database"
	"github.com/extrieve/tekken-babes-gin/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetCharacter godoc
// @Summary      Get detailed information about a character
// @Description  Get character details by ID
// @Tags         Character
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Character ID"
// @Success      200  {object}  models.Character
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/characters/{id} [get]
func GetCharacter(c *gin.Context) {
    idParam := c.Param("id")
    objectID, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var character models.Character
    err = database.CharacterCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&character)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
        return
    }

    c.JSON(http.StatusOK, character)
}

// GetLeaderboard godoc
// @Summary      Retrieve the leaderboard of characters
// @Description  Get characters ranked by total wins
// @Tags         Leaderboard
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Character
// @Failure      500  {object}  map[string]string
// @Router       /api/leaderboard [get]
func GetLeaderboard(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    opts := options.Find().SetSort(bson.D{{Key: "total_wins", Value: -1}})

    cursor, err := database.CharacterCollection.Find(ctx, bson.M{}, opts)
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

    c.JSON(http.StatusOK, characters)
}

