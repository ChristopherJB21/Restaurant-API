package model

import "github.com/google/uuid"

type FoodDeleteRequest struct {
	IDFood    uuid.UUID `validate:"required"`
	DeletedBy string    `validate:"required"`
}
