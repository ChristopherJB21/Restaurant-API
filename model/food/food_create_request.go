package model

type FoodCreateRequest struct {
	FoodName  string `validate:"required,min=1,max=255"`
	IDCuisine uint   `validate:"required"`
}
