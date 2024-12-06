package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transfer struct {
	TransferID    string    `json:"transfer_id"`
	Amount        int64     `json:"amount"`
	Remarks       string    `json:"remarks"`
	BalanceBefore int64     `json:"balance_before"`
	BalanceAfter  int64     `json:"balance_after"`
	CreatedDate   time.Time `json:"created_date"`
}

func (u *Transfer) BeforeCreate(tx *gorm.DB) (err error) {
	u.TransferID = uuid.New().String()
	u.CreatedDate = time.Now()
	return
}
