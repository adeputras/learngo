package songs

import (
	"time"
)

type SongData struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	AlbumId   uint      `json:"albumId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
