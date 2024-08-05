package service

import (
	model "finance-planner/models"
	"finance-planner/repository"
)

type RecordService struct {
	repo *repository.RecordRepository
}

func NewRecordService(repo *repository.RecordRepository) *RecordService {
	return &RecordService{repo: repo}
}

func (s *RecordService) CreateRecord(record *model.Record) error {
	return s.repo.Create(record)
}

func (s *RecordService) ListRecordsByWalletID(walletID string) ([]model.Record, error) {
	return s.repo.ListByWalletID(walletID)
}

func (s *RecordService) ListRecordsByTime(walletID, startTime, endTime string) ([]model.Record, error) {
	return s.repo.ListByTime(walletID, startTime, endTime)
}
