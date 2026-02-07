package router

import (
	"github.com/gin-gonic/gin"

	"master/internal/handler"
)

func SetRoute() *gin.Engine {
	server := gin.Default()

	g := server.Group("/master")
	g.POST("/players", handler.UploadPlayers)
	g.GET("/draw", handler.Draw)
	return server
}
