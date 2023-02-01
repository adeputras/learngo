package entity

import (
	"time"
)

type Album struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `binding:"required" gorm:"type:varchar(100)" json:"name"`
	Year      uint   `binding:"required" json:"year"`
	Songs     []Song `gorm:"foreignKey:AlbumId; references:ID" json:"songs"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
