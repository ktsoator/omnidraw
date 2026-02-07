package handler

import (
	"net/http"

	"master/internal/service"

	"github.com/gin-gonic/gin"
)

func UploadPlayers(c *gin.Context) {
	type PlayersReq struct {
		Players []string `json:"players"`
	}

	var req PlayersReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	if len(req.Players) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "players is empty"})
		return
	}
	service.SetDrawPlayers(req.Players)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func Draw(c *gin.Context) {
	winner, remaining, err := service.DrawPlayer()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "success",
		"winner":    winner,
		"remaining": remaining,
	})
}
