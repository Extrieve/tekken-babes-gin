package database

import (
	"context"
	"log"
	"time"

	"github.com/extrieve/tekken-babes-gin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SeedData() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    count, err := CharacterCollection.CountDocuments(ctx, bson.M{})
    if err != nil {
        log.Fatal("Failed to count documents:", err)
    }

    if count > 0 {
        log.Println("Data already exists in the database.")
        return // Data already exists
    }

    characters := []interface{}{
        models.Character{
            ID:        primitive.NewObjectID(),
            Name:      "Anna Williams",
            ImageURL:  "https://example.com/images/anna.jpg",
            Bio:       "Anna Williams is an Irish female assassin and the younger sister of Nina Williams.",
            TotalWins: 0,
        },
        models.Character{
            ID:        primitive.NewObjectID(),
            Name:      "Nina Williams",
            ImageURL:  "https://example.com/images/nina.jpg",
            Bio:       "Nina Williams is an Irish professional assassin who specializes in assassination and other martial arts.",
            TotalWins: 0,
        },
        // Add more characters...
    }

    _, err = CharacterCollection.InsertMany(ctx, characters)
    if err != nil {
        log.Fatal("Failed to seed data:", err)
    }

    log.Println("Database seeded with initial data.")
}
