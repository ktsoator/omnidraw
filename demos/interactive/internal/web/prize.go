package web

import (
	"interactive/internal/domain"
	"interactive/internal/service"
	"interactive/internal/web/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PrizeHandler struct {
	svc *service.PrizeService
}

func NewPrizeHandler(svc *service.PrizeService) *PrizeHandler {
	return &PrizeHandler{svc: svc}
}

func (h *PrizeHandler) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/prize")
	group.POST("/upload", h.UploadPrize)
}

func (h *PrizeHandler) UploadPrize(ctx *gin.Context) {
	type UploadPrizeReq struct {
		Prizes []domain.Prize `json:"prizes"`
	}
	var req UploadPrizeReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	err := h.svc.AddPrize(ctx, req.Prizes)

	if err != nil {
		ctx.JSON(http.StatusOK, resp.Result{
			Code: 5,
			Msg:  "system error",
		})
		return
	}

	ctx.JSON(http.StatusOK, resp.Result{
		Code: 0,
		Msg:  "success",
	})
}
