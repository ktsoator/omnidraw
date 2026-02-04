package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ktsoator/omnidraw/demos/master/internal/handler"
)

func SetRoute() *gin.Engine {
	r := gin.Default()
	g := r.Group("/master")
	g.POST("/players", handler.UploadPlayers)
	g.GET("/draw", handler.Draw)
	return r
}
