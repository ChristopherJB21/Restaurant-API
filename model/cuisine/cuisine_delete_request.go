package model

import "github.com/google/uuid"

type CuisineDeleteRequest struct {
	IDCuisine uuid.UUID `validate:"required"`
	DeletedBy string    `validate:"required"`
}
