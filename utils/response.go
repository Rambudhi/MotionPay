package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Message string      `json:"message,omitempty"`
}

func JSONResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	var response Response
	if statusCode == 200 || statusCode == 201 {
		response = Response{
			Status: "SUCCESS",
			Result: data,
		}
	} else {
		response = Response{
			Message: message,
		}
	}

	// Kirim response dalam format JSON
	c.JSON(statusCode, response)
}
