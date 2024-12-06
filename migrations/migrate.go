package migrations

import (
	"MotionPay/models"

	"gorm.io/gorm"
)

func Migrate(gorm *gorm.DB) {
	gorm.AutoMigrate(
		&models.User{},
		&models.TopUp{},
		&models.Payment{},
	)
}
