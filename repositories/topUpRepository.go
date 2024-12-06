package repositories

import (
	"MotionPay/models"

	"gorm.io/gorm"
)

type TopUpRepository interface {
	GetTotalAmount(userID string) (int64, error)
	CreateTopUp(topUp *models.TopUp) error
	GetUserBalance(userID string) (int64, error)
}

type topUpRepository struct {
	db *gorm.DB
}

func NewTopUpRepository(db *gorm.DB) TopUpRepository {
	return &topUpRepository{
		db: db,
	}
}

func (r *topUpRepository) GetTotalAmount(userID string) (int64, error) {
	var totalAmount int64
	if err := r.db.Model(&models.TopUp{}).Where("user_id = ?", userID).Select("amount").Scan(&totalAmount).Error; err != nil {
		return 0, err
	}
	return totalAmount, nil
}

func (r *topUpRepository) CreateTopUp(topUp *models.TopUp) error {

	result := r.db.Create(topUp)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *topUpRepository) GetUserBalance(userID string) (int64, error) {
	var balance int64
	err := r.db.Table("top_ups").Select("balance_after AS balance").Where("user_id = ?", userID).Scan(&balance).Error
	if err != nil {
		return 0, err
	}
	return balance, nil
}
