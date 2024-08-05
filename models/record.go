package model

import "time"

type Record struct {
	ID              uint      `gorm:"primaryKey"`
	WalletID        uint      `json:"wallet_id"`
	Amount          float64   `json:"amount"`
	TransactionType string    `json:"transaction_type"` // "income" or "expense"
	CategoryID      uint      `json:"category_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
