package user

import "github.com/google/uuid"

type UserUpdateRequest struct {
	IDUser   uuid.UUID `validate:"required"`
	UserName string    `validate:"required,min=1,max=255"`
}
