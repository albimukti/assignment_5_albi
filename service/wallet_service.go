package service

import (
	model "finance-planner/models"
	"finance-planner/repository"
)

type WalletService struct {
	repo *repository.WalletRepository
}

func NewWalletService(repo *repository.WalletRepository) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) CreateWallet(wallet *model.Wallet) error {
	return s.repo.Create(wallet)
}

func (s *WalletService) GetWalletByID(id uint) (*model.Wallet, error) {
	return s.repo.GetByID(id)
}

func (s *WalletService) ListWalletsByUserID(userID string) ([]model.Wallet, error) {
	return s.repo.ListByUserID(userID)
}
