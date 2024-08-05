package repository

import (
	model "finance-planner/models"

	"gorm.io/gorm"
)

type RecordRepository struct {
	db *gorm.DB
}

func NewRecordRepository(db *gorm.DB) *RecordRepository {
	return &RecordRepository{db: db}
}

func (r *RecordRepository) Create(record *model.Record) error {
	return r.db.Create(record).Error
}

func (r *RecordRepository) ListByWalletID(walletID string) ([]model.Record, error) {
	var records []model.Record
	if err := r.db.Where("wallet_id = ?", walletID).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (r *RecordRepository) ListByTime(walletID, startTime, endTime string) ([]model.Record, error) {
	var records []model.Record
	if err := r.db.Where("wallet_id = ? AND created_at BETWEEN ? AND ?", walletID, startTime, endTime).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}
