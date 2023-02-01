package models

import (
	"time"
)

type Album struct {
	// gorm.Model
	Id        uint   `json:"id"`
	Name      string `gorm:"type:varchar(100)" json:"name"`
	Year      uint   `json:"year"`
	Songs     []Song `json:"songs"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
