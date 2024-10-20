package controllers

import (
	"net/http"
	"strconv"

	"github.com/extrieve/tekken-babes-gin/database"
	"github.com/extrieve/tekken-babes-gin/models"
	"github.com/gin-gonic/gin"
)

func GetCharacter(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
        return
    }

    var character models.Character
    if err := database.DB.First(&character, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
        return
    }

    c.JSON(http.StatusOK, character)
}

func GetLeaderboard(c *gin.Context) {
    var characters []models.Character
    result := database.DB.Order("total_wins desc").Find(&characters)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, characters)
}
