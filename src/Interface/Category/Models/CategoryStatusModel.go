package Models

import "github.com/google/uuid"

type CategoryStatusModel struct {
	Id       uuid.UUID
	IsActive bool
}
