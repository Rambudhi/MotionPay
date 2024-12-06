package main

import (
	"MotionPay/config"
	"MotionPay/migrations"
	"MotionPay/routes"
	"log"
)

func main() {
	config.LoadEnv()

	gorm := config.InitializeDB()
	migrations.Migrate(gorm)

	router := routes.SetupRoutes(gorm)

	log.Println("Starting server on :8080")
	log.Fatal(router.Run(":8080"))
}
