package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Battle struct {
    ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    CharacterOneID primitive.ObjectID `bson:"character_one_id" json:"character_one_id"`
    CharacterTwoID primitive.ObjectID `bson:"character_two_id" json:"character_two_id"`
    WinnerID       primitive.ObjectID `bson:"winner_id" json:"winner_id"`
    BattleTime     time.Time          `bson:"battle_time" json:"battle_time"`
}
