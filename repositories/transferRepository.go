package repositories

import (
	"gorm.io/gorm"

	"MotionPay/models"
)

type TransferRepository interface {
	CreateTransfer(transfer *models.Transfer) error
}

type transferRepository struct {
	db *gorm.DB
}

func NewTransferRepository(db *gorm.DB) TransferRepository {
	return &transferRepository{
		db: db,
	}
}

func (r *transferRepository) CreateTransfer(transfer *models.Transfer) error {
	if err := r.db.Create(transfer).Error; err != nil {
		return err
	}
	return nil
}
