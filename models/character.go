package models

import (
	"time"

	"gorm.io/gorm"
)

type Character struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Name      string         `json:"name"`
    ImageURL  string         `json:"image_url"`
    Bio       string         `json:"bio"`
    TotalWins uint           `json:"total_wins"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
