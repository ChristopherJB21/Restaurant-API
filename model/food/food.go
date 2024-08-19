package model

import (
	model "restaurant/model/cuisine"

	"gorm.io/gorm"
)

type Food struct {
	*gorm.Model
	FoodName  string `gorm:"not null"`
	IDCuisine uint          `gorm:"not null"`
	Cuisine   model.Cuisine `gorm:"foreignKey:IDCuisine"`
}
