package model

type CuisineUpdateRequest struct {
	IDCuisine   uint   `validate:"required"`
	CuisineName string `validate:"required,min=1,max=255"`
}
