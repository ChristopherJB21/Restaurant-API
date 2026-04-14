package model

import (
	"database/sql"
	model "restaurant/model/cuisine"
	"restaurant/model/web"

	"github.com/google/uuid"
)

type Food struct {
	web.BaseModel
	FoodName  string `gorm:"not null"`
	CreatedBy string `gorm:"not null"`
	UpdatedBy string `gorm:"not null"`
	DeletedBy sql.NullString
	IDCuisine uuid.UUID     `gorm:"not null"`
	Cuisine   model.Cuisine `gorm:"foreignKey:IDCuisine"`
}
