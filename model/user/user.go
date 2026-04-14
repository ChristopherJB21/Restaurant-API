package user

import (
	"restaurant/model/web"
)

type User struct {
	web.BaseModel
	UserName string `gorm:"uniqueIndex(255)"`
	Password string `gorm:"not null"`
}
