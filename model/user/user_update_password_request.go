package user

import "github.com/google/uuid"

type UserUpdatePasswordRequest struct {
	IDUser      uuid.UUID `validate:"required"`
	NewPassword string    `validate:"required,min=1,max=255"`
}
