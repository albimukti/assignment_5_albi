package repository

import (
	model "finance-planner/models"

	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) Create(wallet *model.Wallet) error {
	return r.db.Create(wallet).Error
}

func (r *WalletRepository) GetByID(id uint) (*model.Wallet, error) {
	var wallet model.Wallet
	if err := r.db.First(&wallet, id).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *WalletRepository) ListByUserID(userID string) ([]model.Wallet, error) {
	var wallets []model.Wallet
	if err := r.db.Where("user_id = ?", userID).Find(&wallets).Error; err != nil {
		return nil, err
	}
	return wallets, nil
}
