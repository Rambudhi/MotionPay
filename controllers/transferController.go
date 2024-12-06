package controllers

import (
	"MotionPay/services"
	"MotionPay/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type TransferController struct {
	service services.TransferService
}

func NewTransferController(service services.TransferService) *TransferController {
	return &TransferController{
		service: service,
	}
}

func (c *TransferController) Transfer(ctx *gin.Context) {
	var request struct {
		TargetUser string `json:"“target_user”:" binding:"required"`
		Amount     int64  `json:"amount" binding:"required"`
		Remarks    string `json:"remarks" binding:"required"`
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

	payment, err := c.service.ProcessTransfer(userIDStr, request.Amount, request.Remarks)
	if err != nil {
		utils.JSONResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.JSONResponse(ctx, http.StatusOK, "Payment successful", payment)
}
