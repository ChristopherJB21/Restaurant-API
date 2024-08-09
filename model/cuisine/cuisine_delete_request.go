package model

type CuisineDeleteRequest struct {
	IDCuisine uint   `validate:"required"`
	DeletedBy string `validate:"required"`
}
