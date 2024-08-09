package model

type FoodCreateRequest struct {
	FoodName  string `validate:"required,min=1,max=255"`
	CreatedBy   string `validate:"required"`
	UpdatedBy   string `validate:"required"`
	IDCuisine uint   `validate:"required"`
}
