package model

import "github.com/google/uuid"

type FoodUpdateRequest struct {
	IDFood    uuid.UUID `validate:"required"`
	FoodName  string    `validate:"required,min=1,max=255"`
	UpdatedBy string    `validate:"required"`
	IDCuisine uuid.UUID `validate:"required"`
}
