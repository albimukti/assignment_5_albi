package service

import (
	model "finance-planner/models"
	"finance-planner/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(category *model.Category) error {
	return s.repo.Create(category)
}

func (s *CategoryService) GetCategoryByID(id uint) (*model.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) ListCategories() ([]model.Category, error) {
	return s.repo.List()
}
