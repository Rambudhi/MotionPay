package services

import (
	"MotionPay/models"
	"MotionPay/repositories"
	"MotionPay/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	authRepo repositories.AuthRepository
}

type AuthService interface {
	Register(firstName, lastName, phoneNumber, address, pin string) (*models.User, error)
	Login(phoneNumber, pin string) (string, string, *models.User, error)
}

func NewAuthService(authRepo repositories.AuthRepository) AuthService {
	return &AuthServiceImpl{authRepo: authRepo}
}

func (s *AuthServiceImpl) Register(firstName, lastName, phoneNumber, address, pin string) (*models.User, error) {
	user, err := s.authRepo.FindUserByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("Phone Number already registered")
	}

	hashedPin, err := bcrypt.GenerateFromPassword([]byte(pin), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
		Address:     address,
		Pin:         string(hashedPin),
	}

	err = s.authRepo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *AuthServiceImpl) Login(phoneNumber, pin string) (string, string, *models.User, error) {
	user, err := s.authRepo.FindUserByPhoneNumber(phoneNumber)
	if err != nil {
		return "", "", nil, err
	}
	if user == nil {
		return "", "", nil, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Pin), []byte(pin))
	if err != nil {
		return "", "", nil, errors.New("Phone Number and PIN doesn’t match.")
	}

	accessToken, refreshToken, err := utils.GenerateTokens(user.UserID)
	if err != nil {
		return "", "", nil, err
	}

	return accessToken, refreshToken, user, nil
}
