package repository

import (
	model "finance-planner/models"

	"gorm.io/gorm"
)

type TransferRepository struct {
	db *gorm.DB
}

func NewTransferRepository(db *gorm.DB) *TransferRepository {
	return &TransferRepository{db: db}
}

func (r *TransferRepository) Create(transfer *model.Transfer) error {
	return r.db.Create(transfer).Error
}
