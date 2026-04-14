package model

import "github.com/google/uuid"

type CuisineUpdateRequest struct {
	IDCuisine   uuid.UUID `validate:"required"`
	CuisineName string    `validate:"required,min=1,max=255"`
	UpdatedBy   string    `validate:"required"`
}
