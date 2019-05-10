package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

// SetupRouter This method sets up the REST method router.
func SetupRouter() {
	r.GET("/ping", ping)
}

// StartWebEngine Starts the Gin web engine
func StartWebEngine() {
	r.Run("localhost:8080")
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong\n")
}
