package DTOs

import (
	"github.com/google/uuid"
)

type CategoryDto struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
