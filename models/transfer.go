package model

import "time"

type Transfer struct {
	ID           uint      `gorm:"primaryKey"`
	FromWalletID uint      `json:"from_wallet_id"`
	ToWalletID   uint      `json:"to_wallet_id"`
	Amount       float64   `json:"amount"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
