package repositories

import (
	"MotionPay/models"
	"errors"

	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *models.User) error
	FindUserByPhoneNumber(phoneNumber string) (*models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

// NewAuthRepository creates a new instance of AuthRepository
func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

// CreateUser saves a new user in the database
func (r *authRepository) CreateUser(user *models.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindUserByPhoneNumber retrieves a user from the database by phone number
func (r *authRepository) FindUserByPhoneNumber(phoneNumber string) (*models.User, error) {
	var user models.User
	result := r.db.Where("phone_number = ?", phoneNumber).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
