package albums

import (
	"net/http"

	"github.com/adeputras/learngo/config"
	"github.com/adeputras/learngo/entity"
	"github.com/gin-gonic/gin"
)

func GetAlbums(ctx *gin.Context) {
	albums := []entity.Album{}
	config.DB.Preload("Songs").Find(&albums)
	data := make([]Albumitem, len(albums))
	for i, album := range albums {
		data[i] = Albumitem{
			ID:        album.ID,
			Name:      album.Name,
			Year:      album.Year,
			Songs:     NewSongList(album.Songs),
			CreatedAt: album.CreatedAt,
			UpdatedAt: album.UpdatedAt,
		}
	}

	// var data []Albumitem
	// for _, album := range albums {
	// 	data = append(data, Albumitem{
	// 		ID:        album.ID,
	// 		Name:      album.Name,
	// 		Year:      album.Year,
	// 		CreatedAt: album.CreatedAt,
	// 		UpdatedAt: album.UpdatedAt,
	// 	})
	// }

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    data,
		"message": "berhasil",
	})
}

func GetDetailAlbum(ctx *gin.Context) {
	album := entity.Album{}
	id := ctx.Params.ByName("id")
	if err := config.DB.Where("id = ?", id).First(&album).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "album not found",
		})
		return
	}
	config.DB.Preload("Songs").First(&album, id)
	albumData := AlbumData{
		ID:        album.ID,
		Name:      album.Name,
		Songs:     NewSongList(album.Songs),
		Year:      album.Year,
		CreatedAt: album.CreatedAt,
		UpdatedAt: album.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menampilkan Data",
		"data":    albumData,
	})
}

func CreateAlbum(ctx *gin.Context) {
	album := entity.Album{}
	if err := ctx.Bind(&album); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid payload",
		})
		return
	}
	config.DB.Create(&album)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "berhasil",
		"data":    &album,
	})
}

func UpdateAlbum(ctx *gin.Context) {
	album := entity.Album{}
	id := ctx.Params.ByName("id")
	if err := config.DB.Where("id = ?", id).First(&album).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "album not found",
		})
		return
	}

	ctx.BindJSON(&album)
	config.DB.Save(&album)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Merubah Data",
		"data":    &album,
	})
}

func DeleteAlbum(ctx *gin.Context) {
	album := entity.Album{}
	id := ctx.Params.ByName("id")
	if err := config.DB.Where("id = ?", id).First(&album).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid id",
		})
		return
	}

	config.DB.Delete(album, id)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menghapus Albums",
	})
}
