package middlewares

import (
	"MotionPay/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.JSONResponse(c, http.StatusUnauthorized, "Unauthenticated", nil)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.JSONResponse(c, http.StatusUnauthorized, "Invalid token format", nil)
			c.Abort()
			return
		}

		tokenString := parts[1]

		_, err := utils.VerifyToken(tokenString)
		if err != nil {
			utils.JSONResponse(c, http.StatusUnauthorized, "Invalid or expired token", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
