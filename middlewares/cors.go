package middlewares

import "github.com/gin-contrib/cors"

// var corsConfig cors.Config

func GetCorsConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	corsConfig.AllowHeaders = []string{"*"}

	return corsConfig
}
