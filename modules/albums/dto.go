package albums

import (
	"time"

	"github.com/adeputras/learngo/entity"
)

// type AlbumDataInput struct {
// 	Name      string    `json:"name"`
// 	Year      uint      `json:"year"`
// 	CreatedAt time.Time `json:"createdAt"`
// 	UpdatedAt time.Time `json:"updatedAt"`
// }
type AlbumData struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Songs     []SongData `json:"songs"`
	Year      uint       `json:"year"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}
type Albumitem struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Year      uint       `json:"year"`
	Songs     []SongData `json:"songs"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type SongData struct {
	ID      uint   `json:"id"`
	AlbumId uint   `json:"albumId"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	// CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt time.Time `json:"updatedAt"`
}

func NewSongList(songs []entity.Song) []SongData {
	songData := make([]SongData, len(songs))
	for i, data := range songs {
		songData[i] = SongData{
			ID:      data.ID,
			AlbumId: data.AlbumId,
			Title:   data.Title,
			Author:  data.Author,
			// CreatedAt: order.CreatedAt,
			// UpdatedAt: order.UpdatedAt,
		}
	}
	return songData
}
