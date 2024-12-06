package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	PaymentID     string    `json:"payment_id" gorm:"primaryKey;type:uuid"`
	Amount        int64     `gorm:"not null"`
	Remarks       string    `gorm:"not null"`
	BalanceBefore int64     `gorm:"not null"`
	BalanceAfter  int64     `gorm:"not null"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
}

func (u *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	u.PaymentID = uuid.New().String()
	u.CreatedDate = time.Now()
	return
}
