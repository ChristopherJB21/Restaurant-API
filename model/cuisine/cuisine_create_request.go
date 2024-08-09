package model

type CuisineCreateRequest struct {
	CuisineName string `validate:"required,min=1,max=255"`
	CreatedBy   string `validate:"required"`
	UpdatedBy   string `validate:"required"`
}
