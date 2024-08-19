package model

type FoodDeleteRequest struct {
	IDFood    uint   `validate:"required"`
}
