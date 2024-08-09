package model

type FoodUpdateRequest struct {
	IDFood    uint   `validate:"required"`
	FoodName  string `validate:"required,min=1,max=255"`
	UpdatedBy   string `validate:"required"`
	IDCuisine uint   `validate:"required"`
}
