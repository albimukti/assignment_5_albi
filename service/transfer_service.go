package service

import (
	model "finance-planner/models"
	"finance-planner/repository"
)

type TransferService struct {
	repo *repository.TransferRepository
}

func NewTransferService(repo *repository.TransferRepository) *TransferService {
	return &TransferService{repo: repo}
}

func (s *TransferService) CreateTransfer(transfer *model.Transfer) error {
	return s.repo.Create(transfer)
}
