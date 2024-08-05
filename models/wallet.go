package model

import "time"

type Wallet struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
