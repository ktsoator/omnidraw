package main

import (
	"interactive/internal/repository"
	"interactive/internal/repository/dao"
	"interactive/internal/service"
	"interactive/internal/web"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	db := dao.InitDB()
	initPrize(db, r)
	r.Run(":8080")
}

func initPrize(db *gorm.DB, router *gin.Engine) {
	prizeDAO := dao.NewPrizeDAO(db)
	prizeRepo := repository.NewPrizeRepository(prizeDAO)
	prizeService := service.NewPrizeService(prizeRepo)
	prizeHandler := web.NewPrizeHandler(prizeService)
	prizeHandler.RegisterRoutes(router)
}
