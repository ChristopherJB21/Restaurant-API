package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type Cuisine struct {
	*gorm.Model
	CuisineName string `gorm:"not null"`
	CreatedBy   string `gorm:"not null"`
	UpdatedBy   string `gorm:"not null"`
	DeletedBy   sql.NullString
}
