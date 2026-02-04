package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ktsoator/omnidraw/demos/master/internal/service"
)

func UploadPlayers(ctx *gin.Context) {
	type PlayersReq struct {
		Players []string `json:"players"`
	}

	req := PlayersReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	if len(req.Players) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "players is empty"})
		return
	}
	service.SetDrawPlayers(req.Players)
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func Draw(ctx *gin.Context) {
	winner, msg := service.DrawPlayer()
	if winner == "" {
		ctx.JSON(http.StatusOK, gin.H{"message": msg})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "success",
		"winner":    winner,
		"remaining": msg,
	})
}
