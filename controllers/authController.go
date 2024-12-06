package controllers

import (
	"MotionPay/services"
	"MotionPay/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) Register(c *gin.Context) {
	var request struct {
		FirstName   string `json:"first_name" binding:"required"`
		LastName    string `json:"last_name" binding:"required"`
		PhoneNumber string `json:"phone_number" binding:"required"`
		Address     string `json:"address" binding:"required"`
		PIN         string `json:"pin" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	user, err := ac.authService.Register(request.FirstName, request.LastName, request.PhoneNumber, request.Address, request.PIN)
	if err != nil {
		utils.JSONResponse(c, http.StatusConflict, err.Error(), nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "User registered successfully", user)
}

func (ac *AuthController) Login(c *gin.Context) {
	var request struct {
		PhoneNumber string `json:"phone_number" binding:"required"`
		PIN         string `json:"pin" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	// Call the service to login
	accessToken, refreshToken, user, err := ac.authService.Login(request.PhoneNumber, request.PIN)
	if err != nil {
		if err.Error() == "User not found" || err.Error() == "Phone Number and PIN doesn’t match." {
			utils.JSONResponse(c, http.StatusUnauthorized, err.Error(), nil)
		} else {
			utils.JSONResponse(c, http.StatusInternalServerError, "Failed to login", nil)
		}
		return
	}

	session := c.MustGet("session").(*sessions.Session)
	session.Values["user_id"] = user.UserID

	// Simpan session
	if err := session.Save(c.Request, c.Writer); err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "Failed to save session", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Login successful", gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
