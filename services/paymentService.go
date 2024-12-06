package services

import (
	"MotionPay/models"
	"MotionPay/repositories"
	"errors"
)

type PaymentService interface {
	ProcessPayment(userID string, amount int64, remarks string) (*models.Payment, error)
}

type paymentService struct {
	paymentRepo repositories.PaymentRepository
	topUpRepo   repositories.TopUpRepository
}

func NewPaymentService(paymentRepo repositories.PaymentRepository, topUpRepo repositories.TopUpRepository) PaymentService {
	return &paymentService{paymentRepo: paymentRepo, topUpRepo: topUpRepo}
}

func (s *paymentService) ProcessPayment(userID string, amount int64, remarks string) (*models.Payment, error) {
	balance, err := s.topUpRepo.GetUserBalance(userID)
	if err != nil {
		return nil, errors.New("failed to retrieve user balance")
	}

	if balance < amount {
		return nil, errors.New("balance is not enough")
	}

	newBalance := balance - amount

	payment := &models.Payment{
		Amount:        amount,
		Remarks:       remarks,
		BalanceBefore: balance,
		BalanceAfter:  newBalance,
	}

	err = s.paymentRepo.CreatePayment(payment)
	if err != nil {
		return nil, errors.New("failed to create payment record")
	}

	return payment, nil
}
