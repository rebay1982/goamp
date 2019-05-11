package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type artistInfo struct {
	Artist string
	Albums []albumInfo
}

type albumInfo struct {
	Album  string
	Tracks []trackInfo
}

type trackInfo struct {
	Track       string
	TrackNumber int
	Length      int
}

var r = gin.Default()

// SetupRouter This method sets up the REST method router.
func SetupRouter() {
	r.GET("/ping", ping)
	r.GET("/artists", listArtists)
	r.GET("/artists/:artist", listArtistInformation)
}

// StartWebEngine Starts the Gin web engine
func StartWebEngine() {
	r.Run("localhost:8080")
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong\n")
}

func listArtists(c *gin.Context) {
	artists := getArtistList()
	var artistNames []string

	// We're only interested in the artist names for now
	for _, ai := range artists {
		artistNames = append(artistNames, ai.Artist)
	}

	c.JSON(http.StatusOK, artistNames)
}

func listArtistInformation(c *gin.Context) {
	artistRequested := c.Param("artist")
	artists := getArtistList()

	for _, ai := range artists {
		if artistRequested == ai.Artist {
			c.JSON(http.StatusOK, ai)
			return
		}
	}

	c.String(http.StatusNotFound, "Artist [%s] is not available.", artistRequested)
}

func getArtistList() []artistInfo {
	a := []artistInfo{
		artistInfo{"Judas Priest", []albumInfo{
			albumInfo{"Painkiller", []trackInfo{
				trackInfo{"Painkiller", 1, 600},
				trackInfo{"Hell Patrol", 2, 660}}}}},
		artistInfo{"Morbid Angel", []albumInfo{
			albumInfo{"Covenant", []trackInfo{
				trackInfo{"Rapture", 1, 600},
				trackInfo{"God of Emptiness", 10, 660}}},
			albumInfo{"Domination", []trackInfo{
				trackInfo{"Where the Slime Lives", 2, 660},
				trackInfo{"Dawn of the Angry", 6, 600}}}}}}

	return a
}
