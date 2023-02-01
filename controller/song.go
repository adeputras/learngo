package controller

import (
	"net/http"
	"strconv"

	"github.com/adeputras/learngo/config"
	"github.com/adeputras/learngo/models"
	"github.com/gin-gonic/gin"
)

func GetSongs(ctx *gin.Context) {
	songs := []models.Song{}
	config.DB.Find(&songs)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "berhasil",
		"res":     &songs,
	})
}

func GetDetailSongs(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	song := models.Song{}
	song.Id = uint(id)

	config.DB.Find(&song, song.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Gagal Menampilkan Data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menampilkan Data",
		"res":     &song,
	})
}

func CreateSong(ctx *gin.Context) {
	song := models.Song{}
	if err := ctx.Bind(&song); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid payload",
		})
		return
	}
	config.DB.Create(&song)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "berhasil",
		"res":     &song,
	})
}

func UpdateSong(ctx *gin.Context) {
	song := models.Song{}
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
		"message": "Berhasil Menambahkan Data",
		"res":     &song,
	})
}

func DeleteSong(ctx *gin.Context) {
	song := models.Song{}
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
