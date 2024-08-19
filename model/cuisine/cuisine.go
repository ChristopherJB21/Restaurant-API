package model

import (
	"gorm.io/gorm"
)

type Cuisine struct {
	*gorm.Model
	CuisineName string `gorm:"not null"`
}
