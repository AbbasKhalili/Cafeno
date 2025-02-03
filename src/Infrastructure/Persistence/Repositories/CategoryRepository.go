package Repositories

import Domain "Cafeno/Domain/Categories"

type CategoryRepository struct {
}

func NewCategoryRepository() Domain.ICategoryRepository {
	return &CategoryRepository{}
}

func (repo *CategoryRepository) Create() error {
	return nil
}
