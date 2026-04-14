package model

import (
	"database/sql"
	"restaurant/model/web"
)

type Cuisine struct {
	web.BaseModel
	CuisineName string `gorm:"not null"`
	CreatedBy   string `gorm:"not null"`
	UpdatedBy   string `gorm:"not null"`
	DeletedBy   sql.NullString
}
