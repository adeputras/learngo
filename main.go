package main

import (
	"github.com/adeputras/learngo/config"
	"github.com/adeputras/learngo/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDb()
	router := gin.Default()

	api := router.Group("/v1")

	album := api.Group("/albums")
	album.GET("/", controller.GetAlbums)
	album.GET("/:id", controller.GetDetailAlbum)
	album.POST("/", controller.CreateAlbum)
	album.PUT("/:id", controller.UpdateAlbum)
	album.DELETE("/:id", controller.DeleteAlbum)

	song := api.Group("/songs")
	song.GET("/", controller.GetSongs)
	song.GET("/:id", controller.GetDetailSongs)
	song.POST("/", controller.CreateSong)
	song.PUT("/:id", controller.UpdateSong)
	song.DELETE("/:id", controller.DeleteSong)

	router.Run() // listen an
}
