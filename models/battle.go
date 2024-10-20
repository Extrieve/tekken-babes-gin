package models

import (
	"time"

	"gorm.io/gorm"
)

type Battle struct {
    ID             uint           `gorm:"primaryKey" json:"id"`
    CharacterOneID uint           `json:"character_one_id"`
    CharacterTwoID uint           `json:"character_two_id"`
    WinnerID       uint           `json:"winner_id"`
    BattleTime     time.Time      `json:"battle_time"`
    CreatedAt      time.Time      `json:"created_at"`
    UpdatedAt      time.Time      `json:"updated_at"`
    DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
