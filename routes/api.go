package routes

import (
	"MotionPay/controllers"
	"MotionPay/middlewares"
	"MotionPay/repositories"
	"MotionPay/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(gorm *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.SessionMiddleware())

	authRepo := repositories.NewAuthRepository(gorm)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
	}

	topUpRepo := repositories.NewTopUpRepository(gorm)
	topUpService := services.NewTopUpService(topUpRepo)
	topUpController := controllers.NewTopUpController(topUpService)

	paymentRepo := repositories.NewPaymentRepository(gorm)
	paymentService := services.NewPaymentService(paymentRepo, topUpRepo)
	paymentController := controllers.NewPaymentController(paymentService)

	transferRepo := repositories.NewTransferRepository(gorm)
	transferService := services.NewTransferService(transferRepo, topUpRepo)
	transferController := controllers.NewTransferController(transferService)

	privateRoutes := router.Group("/api/features")
	privateRoutes.Use(middlewares.AuthMiddleware())
	{
		privateRoutes.POST("/top-up", topUpController.TopUp)
		privateRoutes.POST("/pay", paymentController.Pay)
		privateRoutes.GET("/transfer", transferController.Transfer)
	}

	return router
}
