package entity

import (
	"time"
)

type Song struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	AlbumId   uint   `binding:"required" json:"album_id"`
	Title     string `binding:"required" gorm:"type:varchar(255)" json:"title"`
	Author    string `binding:"required" gorm:"type:varchar(100)" json:"author"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
