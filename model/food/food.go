package model

import (
	"database/sql"
	model "restaurant/model/cuisine"

	"gorm.io/gorm"
)

type Food struct {
	*gorm.Model
	FoodName  string `gorm:"not null"`
	CreatedBy string `gorm:"not null"`
	UpdatedBy string `gorm:"not null"`
	DeletedBy sql.NullString
	IDCuisine uint          `gorm:"not null"`
	Cuisine   model.Cuisine `gorm:"foreignKey:IDCuisine"`
}
