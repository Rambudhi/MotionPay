package controllers

import (
	"MotionPay/services"
	"MotionPay/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type TopUpController struct {
	service services.TopUpService
}

func NewTopUpController(service services.TopUpService) *TopUpController {
	return &TopUpController{
		service: service,
	}
}

func (c *TopUpController) TopUp(ctx *gin.Context) {
	var request struct {
		Amount int64 `json:"amount" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.JSONResponse(ctx, http.StatusBadRequest, "Data yang dikirim tidak lengkap atau tidak sesuai", nil)
		return
	}

	session := ctx.MustGet("session").(*sessions.Session)
	userID := session.Values["user_id"]

	if userID == nil {
		utils.JSONResponse(ctx, http.StatusUnauthorized, "User tidak terautentikasi", nil)
		return
	}

	userIDStr, ok := userID.(string)
	if !ok || userIDStr == "" {
		utils.JSONResponse(ctx, http.StatusUnauthorized, "User ID di session tidak valid", nil)
		return
	}

	topUp, err := c.service.ProcessTopUp(request.Amount, userIDStr)
	if err != nil {
		utils.JSONResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.JSONResponse(ctx, http.StatusOK, "Top-up berhasil", topUp)
}
