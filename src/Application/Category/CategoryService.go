package Application

import (
	"Cafeno/Interface/Category/Models"
	"Cafeno/Interface/Category/Services"
	"github.com/google/uuid"
)

type CategoryService struct {
}

func NewCategoryService() Interfaces.ICategoryService {
	return &CategoryService{}
}

func (catsrv *CategoryService) Create(model *Models.CategoryModel) (uuid.UUID, error) {
	return uuid.New(), nil
}
