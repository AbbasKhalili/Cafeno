package Domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Category struct {
	Id           uuid.UUID
	Name         string
	IsActive     bool
	CreatedDate  time.Time
	ModifiedDate time.Time
	IsDeleted    bool
	DeletedDate  time.Time
}

func CreateCategory(name string) (*Category, error) {
	if name == "" {
		return nil, errors.New("name must not be nil")
	}
	return &Category{
		uuid.New(),
		name,
		true,
		time.Now(),
		time.Now(),
		false,
		time.Now()}, nil
}

func (category *Category) Update(name string) error {
	category.Name = name
	category.ModifiedDate = time.Now()
	return nil
}

func (category *Category) Delete() error {
	category.IsDeleted = true
	category.DeletedDate = time.Now()
	category.IsActive = false
	return nil
}
