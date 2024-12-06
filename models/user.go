package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID      string    `json:"user_id" gorm:"primaryKey;type:uuid"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number" gorm:"unique;not null"`
	Address     string    `json:"address"`
	Pin         string    `json:"-"`
	CreatedDate time.Time `json:"created_date"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UserID = uuid.New().String()
	u.CreatedDate = time.Now()
	return
}
