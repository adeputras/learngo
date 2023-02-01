package models

import (
	"time"
)

type Song struct {
	// gorm.Model
	Id        uint   `json:"id"`
	AlbumId   uint   `binding:"required" json:"album_id"`
	Title     string `binding:"required" json:"title"`
	Author    string `binding:"required" json:"author"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
