package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()
var s = NewService()

// SetupRouter This method sets up the REST method router.
func SetupRouter() {
	r.GET("/ping", ping)
	r.GET("/artists", listArtists)
	r.GET("/artists/:artist/albums", listArtistAlbums)
	r.GET("/dump", dump)
}

// StartWebEngine Starts the Gin web engine
func StartWebEngine() {
	s.StartService()

	r.Run("localhost:8080")
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong\n")
}

func listArtists(c *gin.Context) {
	artists := s.GetArtists()

	c.JSON(http.StatusOK, artists)
}

func listArtistAlbums(c *gin.Context) {
	albums := s.GetArtistAlbums(c.Param("artist"))

	c.JSON(http.StatusOK, albums)
}

func dump(c *gin.Context) {
	c.JSON(http.StatusOK, s.DumpTracks())
}

func listArtistInformation(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
