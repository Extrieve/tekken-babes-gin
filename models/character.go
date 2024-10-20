package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Character struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name      string             `bson:"name" json:"name"`
    ImageURL  string             `bson:"image_url" json:"image_url"`
    Bio       string             `bson:"bio" json:"bio"`
    TotalWins int                `bson:"total_wins" json:"total_wins"`
}
