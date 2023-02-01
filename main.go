package main

import (
	"github.com/adeputras/learngo/config"
	"github.com/adeputras/learngo/modules/albums"
	"github.com/adeputras/learngo/modules/songs"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDb()
	router := gin.Default()

	api := router.Group("/v1")

	album := api.Group("/albums")
	album.GET("/", albums.GetAlbums)
	album.GET("/:id", albums.GetDetailAlbum)
	album.POST("/", albums.CreateAlbum)
	album.PUT("/:id", albums.UpdateAlbum)
	album.DELETE("/:id", albums.DeleteAlbum)

	song := api.Group("/songs")
	song.GET("/", songs.GetSongs)
	song.GET("/:id", songs.GetDetailSongs)
	song.POST("/", songs.CreateSong)
	song.PUT("/:id", songs.UpdateSong)
	song.DELETE("/:id", songs.DeleteSong)

	router.Run() // listen an
}
