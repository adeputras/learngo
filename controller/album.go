package controller

import (
	"net/http"
	"strconv"

	"github.com/adeputras/learngo/config"
	"github.com/adeputras/learngo/models"
	"github.com/gin-gonic/gin"
)

func GetAlbums(ctx *gin.Context) {
	albums := []models.Album{}
	config.DB.Find(&albums)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "berhasil",
		"res":     &albums,
	})
}

func GetDetailAlbum(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	album := models.Album{}
	album.Id = uint(id)

	config.DB.Preload("Songs").Find(&album, album.Id)
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
		"res":     &album,
	})
}

func CreateAlbum(ctx *gin.Context) {
	album := models.Album{}
	if err := ctx.Bind(&album); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid payload",
		})
		return
	}
	config.DB.Create(&album)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "berhasil",
		"res":     &album,
	})
}

func UpdateAlbum(ctx *gin.Context) {
	album := models.Album{}
	id := ctx.Params.ByName("id")
	if err := config.DB.Where("id = ?", id).First(&album).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "data not found",
		})
		return
	}

	ctx.BindJSON(&album)
	config.DB.Save(&album)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil Menambahkan Data",
		"res":     &album,
	})
}

func DeleteAlbum(ctx *gin.Context) {
	album := models.Album{}
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
