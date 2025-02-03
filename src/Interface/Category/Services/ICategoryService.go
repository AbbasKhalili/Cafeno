package Interfaces

import (
	"Cafeno/Interface/Category/Models"
	"github.com/google/uuid"
)

type ICategoryService interface {
	Create(model *Models.CategoryModel) (uuid.UUID, error)
	//Update(uuid.UUID id, model *Interfaces.CategoryModel) (uuid.UUID, error)
	//Delete() (uuid.UUID, error)
	//Disable() (uuid.UUID, error)
	//Enable() (uuid.UUID, error)
}
