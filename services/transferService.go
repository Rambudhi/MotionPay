package services

import (
	"MotionPay/models"
	"MotionPay/repositories"
	"errors"
)

type TransferService interface {
	ProcessTransfer(userID string, amount int64, remarks string) (*models.Transfer, error)
}

type transferService struct {
	transferRepo repositories.TransferRepository
	topUpRepo    repositories.TopUpRepository
}

func NewTransferService(transferRepo repositories.TransferRepository, topUpRepo repositories.TopUpRepository) TransferService {
	return &transferService{transferRepo: transferRepo, topUpRepo: topUpRepo}
}

func (s *transferService) ProcessTransfer(userID string, amount int64, remarks string) (*models.Transfer, error) {

	balance, err := s.topUpRepo.GetUserBalance(userID)
	if err != nil {
		return nil, errors.New("failed to retrieve user balance")
	}

	if balance < amount {
		return nil, errors.New("balance is not enough")
	}

	newBalance := balance - amount

	transfer := &models.Transfer{
		Amount:        amount,
		Remarks:       remarks,
		BalanceBefore: balance,
		BalanceAfter:  newBalance,
	}

	// Save the transfer record
	err = s.transferRepo.CreateTransfer(transfer)
	if err != nil {
		return nil, errors.New("failed to create transfer record")
	}

	return transfer, nil
}
