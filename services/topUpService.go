package services

import (
	"MotionPay/models"
	"MotionPay/repositories"
	"errors"
)

type TopUpService interface {
	ProcessTopUp(amount int64, userID string) (*models.TopUp, error)
}

type topUpService struct {
	repo repositories.TopUpRepository
}

func NewTopUpService(repo repositories.TopUpRepository) TopUpService {
	return &topUpService{
		repo: repo,
	}
}

func (s *topUpService) ProcessTopUp(amount int64, userID string) (*models.TopUp, error) {
	if amount <= 0 {
		return nil, errors.New("Jumlah Amount Harus lebih besar dari 0")
	}

	totalTopUpAmount, err := s.repo.GetTotalAmount(userID)
	if err != nil {
		return nil, err
	}

	// Hitung saldo setelah top-up
	balanceAfter := totalTopUpAmount + amount

	topUp := &models.TopUp{
		UserID:        userID,
		Amount:        amount,
		BalanceBefore: balanceAfter - amount,
		BalanceAfter:  balanceAfter,
	}

	err = s.repo.CreateTopUp(topUp)

	if err != nil {
		return nil, err
	}

	return topUp, nil
}
