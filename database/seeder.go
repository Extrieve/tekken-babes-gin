package database

import "github.com/extrieve/tekken-babes-gin/models"

// SeedData inserts initial data into the database
func SeedData() {
    var count int64
    DB.Model(&models.Character{}).Count(&count)
    if count > 0 {
        return // Data already exists
    }

    characters := []models.Character{
        {
            Name:     "Anna Williams",
            ImageURL: "https://example.com/images/anna.jpg",
            Bio:      "Anna Williams is an Irish female assassin and the younger sister of Nina Williams.",
        },
        {
            Name:     "Nina Williams",
            ImageURL: "https://example.com/images/nina.jpg",
            Bio:      "Nina Williams is an Irish professional assassin who specializes in assassination and other martial arts.",
        },
        // Add more characters...
    }

    for _, character := range characters {
        DB.Create(&character)
    }
}
