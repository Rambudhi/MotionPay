package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TopUp struct {
	TopUpID       string    `json:"top_up_id" gorm:"primaryKey;type:uuid"`
	UserID        string    `json:"user_id"`
	Amount        int64     `json:"amount_top_up"`
	BalanceBefore int64     `json:"balance_before"`
	BalanceAfter  int64     `json:"balance_after"`
	CreatedDate   time.Time `json:"created_date"`
}

func (u *TopUp) BeforeCreate(tx *gorm.DB) (err error) {
	u.TopUpID = uuid.New().String()
	u.CreatedDate = time.Now()
	return
}
