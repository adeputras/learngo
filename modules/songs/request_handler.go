package songs

import (
	"net/http"

	"github.com/adeputras/learngo/config"
	"github.com/adeputras/learngo/entity"
	"github.com/gin-gonic/gin"
)

func GetSongs(ctx *gin.Context) {
	songs := []entity.Song{}
	config.DB.Find(&songs)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menampilkan Data",
		"data":    &songs,
	})
}

func GetDetailSongs(ctx *gin.Context) {
	songs := []entity.Song{}
	id := ctx.Params.ByName("id")
	if err := config.DB.Where("id = ?", id).First(&songs).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "song not found",
		})
		return
	}
	config.DB.Where("id = ?", id).Find(&songs)
	data := make([]SongData, len(songs))

	for i, song := range songs {
		data[i] = SongData{
			ID:        song.ID,
			AlbumId:   song.AlbumId,
			Title:     song.Title,
			Author:    song.Author,
			CreatedAt: song.CreatedAt,
			UpdatedAt: song.UpdatedAt,
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menampilkan Data",
		"data":    data,
	})
}

func CreateSong(ctx *gin.Context) {
	song := entity.Song{}
	if err := ctx.Bind(&song); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid payload",
		})
		return
	}
	config.DB.Create(&song)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "berhasil",
		"data":    &song,
	})
}

func UpdateSong(ctx *gin.Context) {
	song := entity.Song{}
	id := ctx.Params.ByName("id")
	if err := config.DB.Where("id = ?", id).First(&song).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "data not found",
		})
		return
	}

	ctx.BindJSON(&song)
	config.DB.Save(&song)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Merubah Data",
		"res":     &song,
	})
}

func DeleteSong(ctx *gin.Context) {
	song := entity.Song{}
	id := ctx.Params.ByName("id")
	if err := config.DB.Where("id = ?", id).First(&song).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid id",
		})
		return
	}

	config.DB.Delete(song, id)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menghapus Lagu",
	})
}
